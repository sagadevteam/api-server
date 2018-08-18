package models

// Payments - struct for database
type Payments struct {
	PaymentID   int `db:"payment_id" json:"payment_id"`
	TicketID    int `db:"ticket_id" json:"ticket_id"`
	UserID      int `db:"user_id" json:"user_id"`
	CreatedTime int `db:"created_time" json:"created_time"`
	DeletedTime int `db:"deleted_time" json:"deleted_time"`
}
