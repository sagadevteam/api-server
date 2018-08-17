package controllers

import (
	"api-server/models"
	"api-server/requests"
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
	var signupForm requests.SignupRequest

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

	// Save user
	tx, err := models.DB.Begin()
	defer tx.Rollback()

	user := models.User{}
	user.Email = signupForm.Email
	user.Password = hashedPassword
	user.EthAddress = "0"
	user.EthValue = "0"
	user.SagaPoint = "0"
	user.IsAdmin = 0
	err = user.Save(tx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"msg": "User " + signupForm.Email + " inserted"})
	}
}

// Login the user
func Login(c *gin.Context) {
	var loginForm requests.LoginRequest

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
	user, err = models.FindUserByEmail(loginForm.Email, nil)
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

// Logout the user
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Log out successfully"})
	} else {
		// Foridden in GuestRequired
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	}
}

// Authenticated the user
func Authenticated(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Authenticated", "user": user})
	} else {
		// Foridden in AuthRequired
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	}
}
