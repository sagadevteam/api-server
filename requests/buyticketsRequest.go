package requests

// BuyTicketsRequest is buy ticket schema in post form
type BuyTicketsRequest struct {
	TicketID []int `json:"ticket_id" valid:"required"`
}
