package forms

import (
	"api-server/models"
)

// SignupForm is signup schema in post form
type SignupForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (signupForm *SignupForm) Create() error {
	user := models.User{}
	user.Email = signupForm.Email
	user.Password = signupForm.Password
	user.EthAddress = "0"
	user.EthValue = "0"
	user.SagaPoint = "0"
	user.IsAdmin = 0
	err := user.Create()

	return err
}
