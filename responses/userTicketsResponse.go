package responses

import "api-server/common"

// UserTicketsResponse - response struct for usertickets api
type UserTicketsResponse struct {
	TicketID    int              `db:"ticket_id" json:"ticket_id"`
	InventoryID int              `db:"inventory_id" json:"inventory_id"`
	UserID      common.NullInt64 `db:"user_id" json:"user_id"`
	Time        int              `db:"end_time" json:"end_time"`
	Price       int              `db:"price" json:"price"`
	Metadata    int              `db:"metadata" json:"metadata"`
	RoomNo      string           `db:"room_no" json:"room_no"`
	Title       string           `db:"title" json:"title"`
	Description string           `db:"description" json:"description"`
}
