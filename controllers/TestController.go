package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "api-server/models"
)

// Get user by email
func GetUserByEmail(c *gin.Context) {
	db := models.Session
	email := c.Params.ByName("email")
	user := models.User{}
	err := db.Get(&user, `SELECT * FROM users WHERE email=?`, email)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"user": email, "value": user})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"user": email, "status": "no value", "msg": err.Error()})
	}
}

// Insert test user to db
func GetInsertUser(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter email"})
		return
	}
	db := models.Session
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
