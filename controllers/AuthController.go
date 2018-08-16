package controllers

import (
	"api-server/models"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/contrib/sessions"
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

// Signup the user
func Signup(c *gin.Context) {
	var signupForm models.SignupForm

	err := c.BindJSON(&signupForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// Validate email format
	err = checkmail.ValidateFormat(signupForm.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// Validate email host
	// err = checkmail.ValidateHost(signupForm.Email)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
	// 	return
	// }

	// Validate password
	passwordLen := len(signupForm.Password)
	if passwordLen < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": "Password must longer than 6"})
		return
	}

	// Maybe add email verification code
	// hash password
	bytePassword := []byte(signupForm.Password)
	hashedPassword := hashPassword(bytePassword)
	signupForm.Password = hashedPassword

	// Save signup form
	err = signupForm.Create()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "User " + signupForm.Email + " inserted"})
	}
}

// Login the user
func Login(c *gin.Context) {
	var loginForm models.LoginForm

	session := sessions.Default(c)
	err := c.BindJSON(&loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// Validate email format
	err = checkmail.ValidateFormat(loginForm.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
		return
	}

	// Validate email host
	// err = checkmail.ValidateHost(loginForm.Email)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": err.Error()})
	// 	return
	// }

	// Validate password
	passwordLen := len(loginForm.Password)
	if passwordLen < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "error": "Password must longer than 6"})
		return
	}

	user := models.User{}
	err = user.FindByEmail(loginForm.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Something wrong happened", "error": err.Error()})
		return
	}

	bytePassword := []byte(loginForm.Password)
	if comparePassword(user.Password, bytePassword) != true {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Something wrong happened"})
		return
	}

	// Set session
	session.Set("user", user)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "Login successfully"})
	}
}
