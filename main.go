package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	config "api-server/config"
	database "api-server/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

// DB - tmp global var
var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {

	db := database.Session
	people := []database.User{}
	db.Select(&people, `SELECT * FROM users ORDER BY email ASC`)
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)

	jason = database.User{}
	if err := db.Get(&jason, "SELECT * FROM users WHERE email=?", "hello@john.net"); err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Printf("%#v\n", jason)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + config.API.Port)
}
