package requests

// SignupRequest is signup schema in post form
type SignupRequest struct {
	Email         string `json:"email" valid:"email,required"`
	Password      string `json:"password" valid:"matches([a-zA-Z0-9]+),runelength(6|15),required"`
	PasswordAgain string `json:"password_again" valid:"matches([a-zA-Z0-9]+),runelength(6|15),required"`
}
