package models

import (
	"database/sql"
)

// Payment - struct for database
type Payment struct {
	PaymentID   int `db:"payment_id" json:"payment_id"`
	TicketID    int `db:"ticket_id" json:"ticket_id"`
	UserID      int `db:"user_id" json:"user_id"`
	CreatedTime int `db:"created_time" json:"created_time"`
	DeletedTime int `db:"deleted_time" json:"deleted_time"`
}

// InsertPayments - insert payments with UserId
func InsertPayments(userID int, tickets []int, tx *sql.Tx) error {
	// query statement
	sqlStr := `INSERT INTO payments (
					user_id,
					ticket_id,
					created_time
				)
				VALUES `

	// prepare statements
	vals := []interface{}{}
	for _, ticket := range tickets {
		sqlStr += "(?,?,unix_timestamp()),"
		vals = append(vals, userID, ticket)
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	if tx != nil {
		_, err := tx.Exec(sqlStr, vals...)
		return err
	}

	_, err := DB.Exec(sqlStr, vals...)
	return err
}
