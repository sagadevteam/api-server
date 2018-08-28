package controllers

import (
	"api-server/common"
	"api-server/models"
	"api-server/requests"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const eth = 10000000000000000

func GetUserByEmail(c *gin.Context) {
	email := c.Params.ByName("email")
	tx, err := models.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		log.Fatal(err)
	}

	columns := []string{"*"}
	user, err := models.FindUserByEmail(email, columns, tx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"user": email, "status": "no value", "msg": err.Error()})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, user)
	}
}

func GetInsertUser(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter email"})
		return
	}

	tx, err := models.DB.Begin()
	defer tx.Rollback()
	if err != nil {
	}

	user := models.User{
		Email:      email,
		Password:   email + "_test",
		EthAddress: "0x0",
		EthValue:   "0x0",
		SagaPoint:  0,
		IsAdmin:    0,
	}

	if err := user.Save(tx); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, user)
	}
}

// BuyPoints - user buy point from admin user
func BuyPoints(c *gin.Context) {
	// get user id
	session := sessions.Default(c)
	userID := session.Get("user")
	if userID == nil {
		fmt.Println("Page not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	columns := []string{"*"}
	user, err := models.FindUserByID(userID.(int), columns, nil)
	if err != nil {
		fmt.Println("Page not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	// check if admin
	userAdmin := user.IsAdmin
	if userAdmin != 0 {
		fmt.Println("Admin can not buy points")
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin can not buy points"})
		return
	}

	// check user amount
	userAmount, ok := common.ParseBig256(user.EthValue)
	if !ok {
		fmt.Println("database eth_value type error")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "database eth_value type error", "error": "database eth_value type error"})
		return
	}

	// bind post data
	var buyPointsInput requests.BuyPointsRequest
	if err := c.BindJSON(&buyPointsInput); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data format", "error": err.Error()})
		return
	}

	// Query eth rate
	ethRate, err := models.FindEthrateBySymbol(buyPointsInput.Symbol)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query symbol error", "error": err.Error()})
		return
	}
	if !ethRate.Price.Valid {
		fmt.Printf("symbol %s is nil\n", buyPointsInput.Symbol)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query symbol error", "error": fmt.Sprintf("symbol %s is nil", buyPointsInput.Symbol)})
		return
	}

	// count eth amount to minus
	rate := big.NewFloat(ethRate.Price.Float64)                            // rate as big int
	amount := big.NewFloat(float64(buyPointsInput.Amount))                 // amount as big int
	ethAmount := new(big.Float).Quo(amount, rate)                          // eth = amount / rate
	weiAmount := new(big.Float).Mul(ethAmount, big.NewFloat(float64(eth))) // wei = eth * 10**18
	toSubAmount := common.BigFloatToBigInt(weiAmount)
	finalAmount := new(big.Int).Sub(userAmount, toSubAmount)
	if finalAmount.Cmp(big.NewInt(int64(0))) == -1 {
		fmt.Println("eth value not enough")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "eth value not enough", "error": "eth value not enough"})
		return
	}

	// save user
	tx, err := models.DB.Begin()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "tx begin failed", "error": err.Error()})
		return
	}
	defer tx.Rollback()
	// update admin saga point
	if err := models.SelectAndUpdateAdminWithMinusSagaPoint(buyPointsInput.Amount, tx); err != nil {
		fmt.Println(err.Error())
		if err.Error() == models.ErrorMsgSagaPointNotEnough {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Admin minus saga point error", "error": err.Error()})
		}
		return
	}
	// update user eth and saga point
	user.SagaPoint += buyPointsInput.Amount
	user.EthValue = common.BigIntToHex(finalAmount)
	if err := user.Update(tx); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "User save error", "error": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "tx commit failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Buy point successfully"})
	return
}
