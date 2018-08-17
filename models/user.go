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

// FindByEmail find user by email
func (user *User) FindByEmail(email string) error {
	db := Session
	err := db.Get(user, `SELECT * FROM users WHERE email=?`, email)
	return err
}

// Create user
func (user *User) Create() error {
	db := Session
	_, err := db.Exec(`INSERT INTO users (email, password, eth_addr, eth_value, saga_point, is_admin) VALUES (?, ?, ?, ?, ?, ?)`, user.Email, user.Password, user.EthAddress, user.EthValue, user.SagaPoint, user.IsAdmin)

	return err
}
