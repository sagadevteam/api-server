package models

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullInt64 - type for sql.NullInt64
type NullInt64 sql.NullInt64

// Inventory - struct for database
type Inventory struct {
	InventoryID int       `db:"inventory_id" json:"inventory_id"`
	BuyerID     NullInt64 `db:"buyer_id" json:"buyer_id"`
	Price       int       `db:"price" json:"price"`
	Metadata    int       `db:"metadata" json:"metadata"`
	StartTime   int       `db:"start_time" json:"start_time"`
	EndTime     int       `db:"end_time" json:"end_time"`
	CreatedTime int       `db:"created_time" json:"created_time"`
}

// Insert - insert new inventory into table
func (inventory *Inventory) Insert() error {
	_, err := Session.Exec(`INSERT INTO inventories ( price, metadata, start_time, end_time, created_time) VALUES ( ?, ?, ?, ?, unix_timestamp())`, inventory.Price, inventory.Metadata, inventory.StartTime, inventory.EndTime)
	return err
}

// SelectWithID - insert new inventory into table
func (inventory *Inventory) SelectWithID() (out Inventory, err error) {
	err = Session.Get(&out, `SELECT * FROM inventories WHERE inventory_id=?`, inventory.InventoryID)
	return
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