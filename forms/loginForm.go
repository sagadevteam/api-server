package forms

// LoginForm is login schema in post form
type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
