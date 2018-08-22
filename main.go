package main

import (
	"encoding/gob"
	"net/http"
	"time"

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

func setupRouter(r *gin.Engine) {

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

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

	// Exchange rate
	r.GET("/rate/:symbol", controllers.GetRate)

	// Exchange ETH to SAGA points. format {symbol: 'sagatwd', wei_quantity: '0x3345678'}
	r.POST("/buypoints", controllers.BuyPoints)
	//r.POST("/buypoints", middlewares.AuthRequired(), controllers.BuyPoints)

	// Pay by SAGA points. format: [ticket_id1, ticket_id2...]
	r.POST("/buyticket", middlewares.AuthRequired(), controllers.BuyTickets)

	return
}

func main() {
	// Register models
	gob.Register(models.User{})

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Set session
	store := sessions.NewCookieStore([]byte("sagasessionkey"))
	r.Use(sessions.Sessions("sagasession", store))

	// Setup cors
	// Use this when in production
	corsConfig := cors.Config{
		AllowOrigins:     config.API.CORSDomains,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Set mode
	if config.API.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(cors.New(corsConfig))

	// Setup router
	setupRouter(r)

	// Listen and Server in config.API.Domain:config.API.Port
	// Maybe change to use efficient way to concat string
	r.Run(config.API.Domain + ":" + config.API.Port)
}
