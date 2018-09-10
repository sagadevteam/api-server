package controllers

import (
	"api-server/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRate - get rate with symbol
func GetRate(c *gin.Context) {
	// get symbol
	symbol := c.Param("symbol")

	ethrate, err := models.FindEthrateBySymbol(symbol, nil)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data ethrate find error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ethrate)
	return
}
