package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
	"time"

	config "api-server/config"

	_ "github.com/go-sql-driver/mysql" // driver for mysql
	"github.com/jmoiron/sqlx"
)

var (
	// DB - DB for mysql query
	DB *sqlx.DB
)

// NullInt64 - type for sql.NullInt64
type NullInt64 sql.NullInt64

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

// Scan for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
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
