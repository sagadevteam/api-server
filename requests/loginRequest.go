package requests

// LoginRequest is login schema in post form
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
