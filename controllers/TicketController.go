package controllers

import (
	"api-server/models"
	"api-server/requests"
	"api-server/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetTicketsWithInventoryID - Get tickets with inventory id
func GetTicketsWithInventoryID(c *gin.Context) {
	// get inventory id
	var inventoryID int
	var err error
	inventoryIDIn := c.Query("inventory_id")
	if inventoryID, err = strconv.Atoi(inventoryIDIn); len(inventoryIDIn) > 0 && err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data type not int", "error": err.Error()})
		return
	}

	// query database
	var tickets []models.Ticket
	tickets, err = models.SelectTicketsWithInventoryID(inventoryID)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Data select empty", "error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data select failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tickets": tickets})
	return
}

// GetTickets - Get tickets with user id in session
func GetTickets(c *gin.Context) {
	// get user id
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		fmt.Println("Page not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	userID := user.(models.User).UserID

	// query database
	var tickets []responses.UserTicketsResponse
	var err error
	tickets, err = models.SelectTicketsWithUserID(userID)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Data select empty", "error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data select failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tickets": tickets})
	return
}

// BuyTickets - buy ticket with user id in session
func BuyTickets(c *gin.Context) {
	// get user id
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		fmt.Println("Page not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	userID := user.(models.User).UserID

	// bind post data
	var buyTicketsInput requests.BuyTicketsRequest
	if err := c.BindJSON(&buyTicketsInput); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data format", "error": err.Error()})
		return
	}

	// Validate buy tickets form struct
	if _, err := govalidator.ValidateStruct(buyTicketsInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// query database
	tx, err := models.DB.Begin()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "tx begin failed", "error": err.Error()})
		return
	}
	defer tx.Rollback()
	err = models.SelectTicketsAndUpdate(userID, buyTicketsInput.TicketID, tx)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "buy ticket error", "error": err.Error()})
		return
	}
	err = models.InsertPayments(userID, buyTicketsInput.TicketID, tx)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "add payment error", "error": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "tx commit failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "tickets bought successfully"})
	return
}
