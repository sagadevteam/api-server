package models

import (
	"database/sql"
	"fmt"
	"strings"
)

// User is user schema in mysql
type User struct {
	UserID        int    `db:"user_id" json:"user_id"`
	Email         string `db:"email" json:"email"`
	Password      string `db:"password" json:"-"`
	EthAddress    string `db:"eth_addr" json:"eth_addr"`
	EthPrivateKey string `db:"eth_priv"`
	EthValue      string `db:"eth_value" json:"eth_value"`
	SagaPoint     int    `db:"saga_point" json:"saga_point"`
	IsAdmin       int    `db:"is_admin" json:"is_admin"`
}

// FindUserByEmail find user by email
func FindUserByEmail(email string, columns []string, tx *sql.Tx) (user User, err error) {
	sql := fmt.Sprintf(`SELECT %s FROM users WHERE email=? limit 1`, strings.Join(columns[:], ","))
	if tx != nil {
		err = tx.QueryRow(sql, email).
			Scan(&user.UserID, &user.Email, &user.Password, &user.EthAddress, &user.EthValue, &user.SagaPoint, &user.IsAdmin)
		return
	}
	err = DB.Get(&user, sql, email)
	return
}

// Save user
func (user *User) Save(tx *sql.Tx) (err error) {
	insertQuery := `INSERT INTO users (
						email,
						password,
						eth_addr,
						eth_priv,
						eth_value,
						saga_point,
						is_admin)
					VALUES (?, ?, ?, ?, ?, ?, ?)`
	if tx != nil {
		result, err := tx.Exec(insertQuery, user.Email, user.Password, user.EthAddress, user.EthPrivateKey, user.EthValue, user.SagaPoint, user.IsAdmin)

		userID64, _ := result.LastInsertId()
		user.UserID = int(userID64)

		return err
	}
	_, err = DB.Exec(insertQuery, user.Email, user.Password, user.EthAddress, user.EthValue, user.SagaPoint, user.IsAdmin)

	return err
}
