package models

// SignupForm is signup schema in post form
type SignupForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (signupForm *SignupForm) Create() error {
	db := Session
	_, err := db.Exec(`INSERT INTO users (email, password, eth_addr, eth_value, saga_point, is_admin) VALUES (?, ?, ?, ?, ?, ?)`, signupForm.Email, signupForm.Password, "0", "0", "0", 0)

	return err
}
