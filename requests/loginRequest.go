package requests

// LoginRequest is login schema in post form
type LoginRequest struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"matches([a-zA-Z0-9]+),runelength(6|15),required"`
}
