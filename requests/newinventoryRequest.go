package requests

import "api-server/models"

// NewInventoryRequest is schema in new inventory request
type NewInventoryRequest struct {
	Price       int    `json:"price" valid:"required"`
	StartTime   int    `json:"start_time" valid:"required"`
	EndTime     int    `json:"end_time" valid:"required"`
	Metadata    []int  `json:"metadata" `
	RoomNo      string `json:"room_no" valid:"required"`
	Title       string `json:"title" valid:"required"`
	Description string `json:"description" valid:"required"`
}

// ToTableStruct - transfer request form to table struct
func (in *NewInventoryRequest) ToTableStruct() (out models.Inventory) {
	out.StartTime = in.StartTime
	out.EndTime = in.EndTime
	out.Price = in.Price
	out.RoomNo = in.RoomNo
	out.Title = in.Title
	out.Description = in.Description
	for _, item := range in.Metadata {
		out.Metadata |= item
	}
	return
}
