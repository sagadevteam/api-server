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