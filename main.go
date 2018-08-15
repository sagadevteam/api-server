package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	config "api-server/config"
	controllers "api-server/controllers"
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
	r.GET("/user/:email", func(c *gin.Context) {
		db := database.Session
		email := c.Params.ByName("email")
		user := database.User{}
		err := db.Get(&user, `SELECT * FROM users WHERE email=?`, email)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"user": email, "value": user})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"user": email, "status": "no value", "msg": err.Error()})
		}
	})

	// Insert test user to db
	r.GET("/insertUser", controllers.GetInsertUser)

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
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + config.API.Port)
}
