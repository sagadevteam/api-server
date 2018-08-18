package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"api-server/models"
	"api-server/requests"

	"github.com/gin-gonic/gin"
)

// GetInventory - Get inventory with inventory_id
func GetInventory(c *gin.Context) {
	// get inventory id
	inventoryIDQuery := c.Query("inventory_id")
	inventoryID, err := strconv.Atoi(inventoryIDQuery)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data type not int", "error": err.Error()})
		return
	}

	// query database
	inventory, err := models.FindInventoryByID(inventoryID)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Data select empty", "error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data select failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventory)

	return
}

// GetInventories - Get inventories with page and page size
func GetInventories(c *gin.Context) {
	// get page and page size
	var page, pageSize int
	var err error
	pageIn := c.Query("page")
	pageSizeIn := c.Query("page_size")
	if page, err = strconv.Atoi(pageIn); len(pageIn) > 0 && err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data type not int", "error": err.Error()})
		return
	}
	if pageSize, err = strconv.Atoi(pageSizeIn); len(pageSizeIn) > 0 && err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data type not int", "error": err.Error()})
		return
	}

	// query database
	var inventories []models.Inventory
	inventories, err = models.SelectInventoriesWithPage(page, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Data select empty", "error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data select failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"inventories": inventories})

	return
}

// AddInventory - Post function for adding inventory
func AddInventory(c *gin.Context) {
	// bind post data
	var inventoryInput requests.NewInventoryRequest
	if err := c.BindJSON(&inventoryInput); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// insert into database
	tx, err := models.DB.Begin()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "tx begin failed", "error": err.Error()})
		return
	}
	inventory := inventoryInput.ToTableStruct()
	if err := inventory.Save(tx); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data insert failed", "error": err.Error()})
		return
	}
	if err := models.InsertManyTickets(inventory.InventoryID, inventory.StartTime, inventory.EndTime, tx); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data insert failed", "error": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "tx commit failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Inventory inserted"})

	return
}
