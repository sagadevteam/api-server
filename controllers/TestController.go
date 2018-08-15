package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "api-server/database"
)

// Insert test user to db
func GetInsertUser(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter email"})
		return
	}
	db := database.Session
	tx := db.MustBegin()
	defer tx.Rollback()
	_, err := tx.Exec(`INSERT INTO users (email, password, eth_addr, eth_value, saga_point, is_admin) VALUES (?, ?, ?, ?, ?, ?)`, email, email+"_test", "0x0", "0x0", "0x0", 1)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"msg": "User " + email + " inserted"})
	}
}
