package forms

// SignupForm is signup schema in post form
type SignupForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
