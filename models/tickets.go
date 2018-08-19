package models

import (
	"api-server/common"
	"api-server/responses"
	"database/sql"
)

// Tickets - struct for database
type Ticket struct {
	TicketID    int              `db:"ticket_id" json:"ticket_id"`
	InventoryID int              `db:"inventory_id" json:"inventory_id"`
	UserID      common.NullInt64 `db:"user_id" json:"user_id"`
	Time        int              `db:"end_time" json:"end_time"`
	OnChain     bool             `db:"on_chain" json:"on_chain"`
}

// Save - insert one ticket into table
func (ticket *Ticket) Save(tx *sql.Tx) error {

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
func SelectTicketsWithInventoryID(inventoryID int) (tickets []Ticket, err error) {
	rows, errQuery := DB.Query(`SELECT ticket_id, user_id, inventory_id, time FROM tickets WHERE inventory_id=?`, inventoryID)
	if errQuery != nil {
		err = errQuery
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ticket Ticket
		err = rows.Scan(&ticket.TicketID, &ticket.UserID, &ticket.InventoryID, &ticket.Time)
		if err != nil {
			return
		}
		tickets = append(tickets, ticket)
	}
	err = rows.Err()
	return
}

// SelectTicketsWithUserID - find tickets by user id
func SelectTicketsWithUserID(userID int) (tickets []responses.UserTicketsResponse, err error) {
	rows, errQuery := DB.Query(`SELECT t.ticket_id, t.user_id, t.inventory_id, t.time, i.price, i.title, i.description, i.room_no, i.metadata  FROM tickets AS t INNER JOIN inventories AS i ON T.inventory_id=I.inventory_id WHERE user_id=?`, userID)
	if errQuery != nil {
		err = errQuery
		return
	}
	defer rows.Close()
	for rows.Next() {
		var response responses.UserTicketsResponse
		err = rows.Scan(&response.TicketID, &response.UserID, &response.InventoryID, &response.Time, &response.Price, &response.Title, &response.Description, &response.RoomNo, &response.Metadata)
		if err != nil {
			return
		}
		tickets = append(tickets, response)
	}
	err = rows.Err()
	return
}
