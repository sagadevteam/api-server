package models

import (
	"log"
	"time"

	config "api-server/config"

	_ "github.com/go-sql-driver/mysql" // driver for mysql
	"github.com/jmoiron/sqlx"
)

var (
	// Session - Session for mysql query
	db *sqlx.DB
)

func init() {
	var err error
	dbStr := config.DB.User + ":" + config.DB.Pwd + "@tcp(" + config.DB.Host + ")/" + config.DB.Table
	db, err = sqlx.Connect(`mysql`, dbStr)
	db.SetMaxIdleConns(config.DB.MaxIdleConn)
	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetMaxOpenConns(config.DB.MaxConn)
	if err != nil {
		log.Fatalln(err)
	}
}
