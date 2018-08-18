package models

import (
	"database/sql"
)

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

// Save - insert new inventory into table
func (inventory *Inventory) Save(tx *sql.Tx) error {

	result, err := tx.Exec(
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

	InventoryID64, _ := result.LastInsertId()
	inventory.InventoryID = int(InventoryID64)
	return err
}

// FindInventoryByID - find inventory with id
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

func (inventory *Inventory) ToTicketTableStruct() (ticket Tickets) {
	ticket.InventoryID = inventory.InventoryID
	return
}
