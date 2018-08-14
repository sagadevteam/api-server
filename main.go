package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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

type User struct {
	ID         int    `db:"id"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	EthAddress string `db:"eth_address"`
	IsAdmin    int    `db:"is_admin"`
}

func main() {
	db, err := sqlx.Connect(`mysql`, `root:123456@/saga`)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec(`INSERT INTO users (id, email, password, eth_address, is_admin) VALUES (?, ?, ?, ?, ?)`, 1, "jmoiron@jmoiron.net", "Jason", "0x0", 1)
	tx.MustExec(`INSERT INTO users (id, email, password, eth_address, is_admin) VALUES (?, ?, ?, ?, ?)`, 2, "hello@john.net", "John", "0x0", 0)
	tx.NamedExec(`INSERT INTO users (id, email, password, eth_address, is_admin) VALUES (?, ?, ?, ?, ?)`, &User{3, "guest@jane.net", "Jane", "0x0", 0})
	tx.Commit()

	people := []User{}
	db.Select(&people, `SELECT * FROM users ORDER BY email ASC`)
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)

	jason = User{}
	err = db.Get(&jason, "SELECT * FROM users WHERE email=$1", "hello@john.net")
	fmt.Printf("%#v\n", jason)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
