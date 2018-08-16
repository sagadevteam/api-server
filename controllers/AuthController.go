package controllers

import (
	"api-server/models"
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	return string(hash)
}

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		return false
	}
	return true
}

// SignupData is login schema in post form
type SignupData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (signupData *SignupData) Save() error {
	// hash password
	bytePassword := []byte(signupData.Password)
	hashedPassword := hashPassword(bytePassword)
	db := models.Session
	_, err := db.Exec(`INSERT INTO users (email, password, eth_addr, eth_value, saga_point, is_admin) VALUES (?, ?, ?, ?, ?, ?)`, signupData.Email, hashedPassword, "0", "0", "0", 1)

	return err
}

func Signup(c *gin.Context) {
	var signupData SignupData

	err := c.BindJSON(&signupData)

	if err != nil {
		fmt.Println(err)
	}

	// Validate email format
	err = checkmail.ValidateFormat(signupData.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// Validate email host
	// err = checkmail.ValidateHost(signupData.Email)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
	// 	return
	// }

	// Validate password
	passwordLen := len(signupData.Password)
	if passwordLen < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": "Password must longer than 6"})
		return
	}

	// Maybe add email verification code
	// Save signup data
	err = signupData.Save()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "User " + signupData.Email + " inserted"})
	}
}
