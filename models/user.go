package models

// User is user schema in mysql
type User struct {
	UserID     int    `db:"user_id"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	EthAddress string `db:"eth_addr"`
	EthValue   string `db:"eth_value"`
	SagaPoint  string `db:"saga_point"`
	IsAdmin    int    `db:"is_admin"`
}

// FindUserByEmail find user by email
func FindUserByEmail(email string) (User, error) {
	userModel := User{}
	err := DB.Get(&userModel, `SELECT * FROM users WHERE email=?`, email)
	return userModel, err
}

// Save user
func (user *User) Save() error {
	_, err := DB.Exec(`
	INSERT INTO users (
		email,
		password,
		eth_addr,
		eth_value,
		saga_point,
		is_admin) 
	VALUES (?, ?, ?, ?, ?, ?)`,
		user.Email, user.Password, user.EthAddress, user.EthValue, user.SagaPoint, user.IsAdmin)

	return err
}
