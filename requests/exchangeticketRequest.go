package requests

// ExchangeTicketRequest is change ticket to another user schema in post form
type ExchangeTicketRequest struct {
	TicketID int `json:"ticket_id" valid:"required"`
	UserID   int `json:"user_id" valid:"required"`
}
