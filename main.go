package main

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	config "api-server/config"
	controllers "api-server/controllers"
	"api-server/middlewares"
	"api-server/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

// DB - tmp global var
var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Set session
	store := sessions.NewCookieStore([]byte("sagasessionkey"))
	r.Use(sessions.Sessions("sagasession", store))

	// Setup cors
	// Use this when in production
	// config := cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8081"},
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH", "HEAD"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	// r.GET("/user/:email", controllers.GetUserByEmail)

	// Insert test user to db
	r.GET("/insertUser", controllers.GetInsertUser)

	// Signup the user
	r.POST("/signup", middlewares.GuestRequired(), controllers.Signup)

	// Login the user
	r.POST("/login", middlewares.GuestRequired(), controllers.Login)

	// Logout the user
	r.POST("/logout", middlewares.AuthRequired(), controllers.Logout)

	// Get authenticated the user
	r.GET("/authenticated", middlewares.AuthRequired(), controllers.Authenticated)

	// Add new inventory
	r.POST("/inventory", controllers.AddInventory)

	// Get inventory with id
	r.GET("/inventory", controllers.GetInventory)

	// Get inventories
	r.GET("/inventories", controllers.GetInventories)

	// Get tickets
	r.GET("/tickets", controllers.GetTicketsWithInventoryID)

	// Get tickets
	r.GET("/usertickets", middlewares.AuthRequired(), controllers.GetTickets)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		DB[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	// Register models
	gob.Register(models.User{})

	// Setup router
	r := setupRouter()

	// Listen and Server in config.API.Domain:config.API.Port
	// Maybe change to use efficient way to concat string
	r.Run(config.API.Domain + ":" + config.API.Port)
}
