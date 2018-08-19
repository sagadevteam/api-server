package models

import (
	"log"
	"time"

	config "api-server/config"

	_ "github.com/go-sql-driver/mysql" // driver for mysql
	"github.com/jmoiron/sqlx"
)

var (
	// DB - DB for mysql query
	DB *sqlx.DB
)

const defaultSize = 10 // default page size 10

func init() {
	var err error
	dbStr := config.DB.User + ":" + config.DB.Pwd + "@tcp(" + config.DB.Host + ")/" + config.DB.Table
	DB, err = sqlx.Connect(`mysql`, dbStr)
	DB.SetMaxIdleConns(config.DB.MaxIdleConn)
	DB.SetConnMaxLifetime(2 * time.Minute)
	DB.SetMaxOpenConns(config.DB.MaxConn)
	if err != nil {
		log.Fatalln(err)
	}
}

func pageToLimit(page, pageSize int) (limit, limitSize int) {
	if pageSize == 0 {
		limitSize = defaultSize
	} else {
		limitSize = pageSize
	}
	limit = page * limitSize
	return
}
