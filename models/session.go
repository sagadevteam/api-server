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
	Session *sqlx.DB
)

func init() {
	var err error
	dbStr := config.DB.User + ":" + config.DB.Pwd + "@tcp(" + config.DB.Host + ")/" + config.DB.Table
	Session, err = sqlx.Connect(`mysql`, dbStr)
	Session.SetMaxIdleConns(config.DB.MaxIdleConn)
	Session.SetConnMaxLifetime(2 * time.Minute)
	Session.SetMaxOpenConns(config.DB.MaxConn)
	if err != nil {
		log.Fatalln(err)
	}
}
