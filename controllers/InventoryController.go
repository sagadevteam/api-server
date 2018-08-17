package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"api-server/models"

	"github.com/gin-gonic/gin"
)

// inventoryForm is inventory schema in post form
type inventoryForm struct {
	Price     int   `json:"price"`
	StartTime int   `json:"start_time"`
	EndTime   int   `json:"end_time"`
	Metadata  []int `json:"metadata"`
}

// GetInventory - Get inventory with inventory_id
func GetInventory(c *gin.Context) {

	// get inventory id
	inventoryID := c.Query("id")
	var inventory models.Inventory
	var err error
	inventory.InventoryID, err = strconv.Atoi(inventoryID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data type not int", "error": err.Error()})
		return
	}

	// query database
	inventory, err = inventory.SelectWithID()
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Data select empty", "error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data select failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Inventory inserted", "result": inventory})

	return
}

// AddInventory - Post function for adding inventory
func AddInventory(c *gin.Context) {

	// bind post data
	var inventoryInput inventoryForm
	if err := c.BindJSON(&inventoryInput); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// auth seesion and admin

	// insert into database
	inventory := inventoryInput.toTableStruct()
	if err := inventory.Insert(); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data insert failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Inventory inserted"})

	return
}

func (in *inventoryForm) toTableStruct() (out models.Inventory) {
	out.StartTime = in.StartTime
	out.EndTime = in.EndTime
	out.Price = in.Price
	for _, item := range in.Metadata {
		out.Metadata |= item
	}
	return
}
