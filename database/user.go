package database

// User is user schema in mysql
type User struct {
	ID         int    `db:"id"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	EthAddress string `db:"eth_address"`
	IsAdmin    int    `db:"is_admin"`
}
