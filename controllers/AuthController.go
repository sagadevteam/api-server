package controllers

import (
	"api-server/config"
	apiCrypto "api-server/crypto"
	"api-server/models"
	"api-server/requests"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/crypto"
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

// replace zero prefix
// please ensure your string is lowercase
func replaceZeroPrefix(hexString string) string {
	if strings.Index(hexString, "0x") < 0 {
		return hexString
	}
	return strings.Replace(hexString, "0x", "", 1)
}

// Signup the user
func Signup(c *gin.Context) {
	var signupForm requests.SignupRequest

	err := c.BindJSON(&signupForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "err": err.Error()})
		return
	}

	// Validate signup form struct
	_, err = govalidator.ValidateStruct(signupForm)
	if err != nil {
		errMap := govalidator.ErrorsByField(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "err": errMap})
		return
	}

	// Maybe add same tag in govalidator
	if signupForm.Password != signupForm.PasswordAgain {
		errMap := make(map[string]string)
		errMap["password_again"] = "Password again must be equal to password"
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "err": errMap})
		return
	}

	// Maybe add email verification code
	// hash password
	bytePassword := []byte(signupForm.Password)
	hashedPassword := hashPassword(bytePassword)

	// Save user
	tx, err := models.DB.Begin()
	defer tx.Rollback()

	// Generate private key
	privKey, _ := crypto.GenerateKey()

	// Encrypt private key
	var encPrivKey string
	encPrivKey, err = apiCrypto.Encrypt([]byte(config.API.CipherKey), strings.ToLower(hex.EncodeToString(crypto.FromECDSA(privKey))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Something wrong happened", "err": err.Error()})
		return
	}

	user := models.User{}
	user.Email = signupForm.Email
	user.Password = hashedPassword
	user.EthAddress = replaceZeroPrefix(strings.ToLower(crypto.PubkeyToAddress(privKey.PublicKey).Hex()))
	user.EthPrivateKey = encPrivKey
	user.EthValue = "0"
	user.SagaPoint = 0
	user.IsAdmin = 0
	if err = user.Save(tx); err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "err": err.Error()})
		return
	}

	// Validate login form struct
	_, err = govalidator.ValidateStruct(loginForm)
	if err != nil {
		errMap := govalidator.ErrorsByField(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "err": errMap})
		return
	}

	columns := []string{"*"}
	user := models.User{}
	user, err = models.FindUserByEmail(loginForm.Email, columns, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Something wrong happened", "err": err.Error()})
		return
	}

	bytePassword := []byte(loginForm.Password)
	if comparePassword(user.Password, bytePassword) != true {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Something wrong happened"})
		return
	}

	// Set session
	session.Set("user", user.UserID)
	if err = session.Save(); err != nil {
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
		c.JSON(http.StatusOK, gin.H{"msg": "Log out successfully"})
	} else {
		// Foridden in GuestRequired
		c.JSON(http.StatusNotFound, gin.H{"err": "Page not found"})
	}
}

// Authenticated the user
func Authenticated(c *gin.Context) {
	session := sessions.Default(c)
	columns := []string{"*"}
	if userID := session.Get("user"); userID != nil {
		if user, err := models.FindUserByID(userID.(int), columns, nil); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Something wrong happened", "err": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "Authenticated", "user": user})
		}
	} else {
		// Foridden in AuthRequired
		c.JSON(http.StatusNotFound, gin.H{"err": "Page not found"})
	}
}
