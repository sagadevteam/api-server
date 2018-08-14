package database

import (
	"log"

	config "api-server/config"

	_ "github.com/go-sql-driver/mysql" // driver for mysql
	"github.com/jmoiron/sqlx"
)

var (
	// Session - Session for mysql query
	Session *sqlx.DB
)

func init() {
	var err error
	dbStr := config.DB.User + ":" + config.DB.Pwd + "@tcp(" + config.DB.Host + ")/" + config.DB.Table
	Session, err = sqlx.Connect(`mysql`, dbStr)
	if err != nil {
		log.Fatalln(err)
	}
}
