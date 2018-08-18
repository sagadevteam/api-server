package controllers

import (
	"api-server/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTickets - Get tickets with inventory id
func GetTickets(c *gin.Context) {
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
