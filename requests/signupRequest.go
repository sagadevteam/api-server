package requests

// SignupRequest is signup schema in post form
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
