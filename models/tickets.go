package models

import "database/sql"

const dayTime = 86400

// Tickets - struct for database
type Tickets struct {
	TicketID    int       `db:"ticket_id" json:"ticket_id"`
	InventoryID int       `db:"inventory_id" json:"inventory_id"`
	UserID      NullInt64 `db:"user_id" json:"user_id"`
	Time        int       `db:"end_time" json:"end_time"`
}

// Save - insert one ticket into table
func (ticket *Tickets) Save(tx *sql.Tx) error {

	_, err := tx.Exec(
		`INSERT INTO tickets (
			inventory_id, 
			time) 
		VALUES 
	  		( ?, ?, ?)`,
		ticket.InventoryID,
		ticket.Time)

	return err
}

// InsertManyTickets - insert many ticket into table
func InsertManyTickets(inventoryID, start, end int, tx *sql.Tx) error {
	sqlStr := `INSERT INTO tickets (
					inventory_id,
					time)
				VALUES `
	vals := []interface{}{}
	for time := start; time < end; time += dayTime {
		sqlStr += "(?,?),"
		vals = append(vals, inventoryID, time)
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	if tx != nil {
		_, err := tx.Exec(sqlStr, vals...)
		return err
	}

	_, err := DB.Exec(sqlStr, vals...)
	return err
}

// SelectTicketsWithInventoryID - find tickets by inventory id
func SelectTicketsWithInventoryID(inventoryID int) (tickets []Tickets, err error) {
	rows, errQuery := DB.Query(`SELECT ticket_id, user_id, inventory_id, time FROM tickets WHERE inventory_id=?`, inventoryID)
	if errQuery != nil {
		err = errQuery
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ticket Tickets
		err = rows.Scan(&ticket.TicketID, &ticket.UserID, &ticket.InventoryID, &ticket.Time)
		if err != nil {
			return
		}
		tickets = append(tickets, ticket)
	}
	err = rows.Err()
	return
}
