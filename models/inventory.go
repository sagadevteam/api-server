package models

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

const defaultSize = 10 // default page size 10

// NullInt64 - type for sql.NullInt64
type NullInt64 sql.NullInt64

// Inventory - struct for database
type Inventory struct {
	InventoryID int    `db:"inventory_id" json:"inventory_id"`
	Price       int    `db:"price" json:"price"`
	Metadata    int    `db:"metadata" json:"metadata"`
	StartTime   int    `db:"start_time" json:"start_time"`
	EndTime     int    `db:"end_time" json:"end_time"`
	CreatedTime int    `db:"created_time" json:"created_time"`
	RoomNo      string `db:"room_no" json:"room_no"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

// Insert - insert new inventory into table
func (inventory *Inventory) Save(tx *sql.Tx) error {

	_, err := tx.Exec(
		`INSERT INTO inventories (
			price, 
			metadata, 
			start_time, 
			end_time, 
			created_time,
			room_no,
			title,
			description) 
		VALUES 
	  		( ?, ?, ?, ?, unix_timestamp(), ?, ?, ?)`,
		inventory.Price,
		inventory.Metadata,
		inventory.StartTime,
		inventory.EndTime,
		inventory.RoomNo,
		inventory.Title,
		inventory.Description)
	return err
}

// SelectWithID - select inventory with id
func FindInventoryByID(inventoryID int) (inventory Inventory, err error) {
	err = DB.Get(&inventory, `SELECT * FROM inventories WHERE inventory_id=?`, inventoryID)
	return
}

// SelectInventoriesWithPage - select inventories with page and page size
func SelectInventoriesWithPage(page, pageSize int) (inventories []Inventory, err error) {

	// change page to limit
	limit, limitSize := pageToLimit(page, pageSize)

	// select inventories with limit
	err = DB.Select(&inventories, `SELECT * FROM inventories ORDER BY created_time DESC LIMIT ?, ?`, limit, limitSize)
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

func pageToLimit(page, pageSize int) (limit, limitSize int) {
	if pageSize == 0 {
		limitSize = defaultSize
	} else {
		limitSize = pageSize
	}
	limit = page * limitSize
	return
}
