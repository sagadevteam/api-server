package controllers

import (
	"api-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByEmail(c *gin.Context) {
	email := c.Params.ByName("email")
	user, err := models.FindUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"user": email, "status": "no value", "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetInsertUser(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter email"})
		return
	}
	user := models.User{
		Email:      email,
		Password:   email + "_test",
		EthAddress: "0x0",
		EthValue:   "0x0",
		SagaPoint:  "0x0",
		IsAdmin:    0,
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
