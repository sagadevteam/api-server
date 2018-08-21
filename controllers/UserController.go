package controllers

import (
	"api-server/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
