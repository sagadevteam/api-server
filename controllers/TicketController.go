package controllers

import (
	"api-server/models"
	"api-server/responses"
	"fmt"
	"net/http"
	"strconv"

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
	var tickets []models.Tickets
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
