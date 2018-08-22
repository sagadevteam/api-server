package requests

// BuyPointsRequest is buy point schema in post form
type BuyPointsRequest struct {
	Symbol string `json:"symbol" valid:"required"`
	Amount int    `json:"amount" valid:"required"`
}
