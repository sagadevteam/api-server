package models

import (
	"api-server/common"
	"api-server/responses"
	"database/sql"
	"errors"
)

// Ticket - struct for database
type Ticket struct {
	TicketID    int              `db:"ticket_id" json:"ticket_id"`
	InventoryID int              `db:"inventory_id" json:"inventory_id"`
	UserID      common.NullInt64 `db:"user_id" json:"user_id"`
	Time        int              `db:"end_time" json:"end_time"`
	OnChain     bool             `db:"on_chain" json:"on_chain"`
	ToUser      bool             `db:"to_user" json:"to_user"`
}

// Save - insert one ticket into table
func (ticket *Ticket) Save(tx *sql.Tx) error {

	_, err := tx.Exec(
		`INSERT INTO tickets (
			inventory_id, 
			time) 
		VALUES 
	  		( ?, ?)`,
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

// SelectTicketsAndUpdate - select tickets and update point
func SelectTicketsAndUpdate(user *User, tickets []int, tx *sql.Tx) (err error) {
	// query statement
	selectTicketBoughtQuery := "SELECT t.user_id, t.inventory_id, i.price FROM tickets AS t LEFT JOIN inventories AS i ON t.inventory_id=i.inventory_id WHERE ticket_id=?"
	updateTicketSetUserQuery := "UPDATE tickets SET user_id=? WHERE ticket_id=?"
	updateUserSetPointQuery := "UPDATE users SET saga_point=? WHERE user_id=?"
	// find tickets
	var inventory, totalPrice int
	for _, item := range tickets {
		var userID common.NullInt64
		var inventoryTmp, price int
		err = tx.QueryRow(selectTicketBoughtQuery, item).Scan(&userID, &inventoryTmp, &price)
		if err != nil {
			return
		}
		if userID.Valid {
			err = errors.New("ticket already bought")
			return
		}
		if inventory == 0 {
			inventory = inventoryTmp
		} else {
			if inventory != inventoryTmp {
				err = errors.New("inventory not same")
				return
			}
		}
		totalPrice += price
	}

	// check user sagapoint
	if user.SagaPoint < totalPrice {
		err = errors.New("saga point not enough")
		return
	}
	user.SagaPoint -= totalPrice

	// update all
	for _, item := range tickets {
		_, err = tx.Exec(updateTicketSetUserQuery, user.UserID, item)
		if err != nil {
			return
		}
	}
	_, err = tx.Exec(updateUserSetPointQuery, user.SagaPoint, user.UserID)
	return
}

// SelectTicketWithUserID - find tickets by user id
func SelectTicketWithUserID(userID, ticketID int, tx *sql.Tx) (owned bool, err error) {
	var ownerID common.NullInt64
	sql := `SELECT user_id FROM tickets WHERE ticket_id=?`
	if tx != nil {
		err = tx.QueryRow(sql, ticketID).Scan(&ownerID)
	} else {
		err = DB.QueryRow(`sql`, ticketID).Scan(&ownerID)
	}
	if err != nil {
		return
	}
	if ownerID.Valid && ownerID.Int64 == int64(userID) {
		owned = true
	}
	return
}

// UpdateTicketUserID - find tickets by user id
func UpdateTicketUserID(userID, ticketID int, tx *sql.Tx) (err error) {
	sql := `UPDATE tickets SET user_id=? WHERE ticket_id=?`
	if tx != nil {
		_, err = tx.Exec(sql, userID, ticketID)
	} else {
		_, err = DB.Exec(sql, userID, ticketID)
	}
	return
}
