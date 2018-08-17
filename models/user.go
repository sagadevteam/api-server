package models

import (
	"database/sql"
	"fmt"
)

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
func FindUserByEmail(email string, tx *sql.Tx) (user User, err error) {
	sql := `SELECT * FROM users WHERE email=? limit 1`
	if tx != nil {
		err = tx.QueryRow(sql, email).
			Scan(&user.UserID, &user.Email, &user.Password, &user.EthAddress, &user.EthValue, &user.SagaPoint, &user.IsAdmin)
		return
	}
	err = DB.Get(&user, sql, email)
	return
}

// Save user
func (user *User) Save(tx *sql.Tx) error {
	var err error
	insertQuery := `INSERT INTO users (
				email,
				password,
				eth_addr,
				eth_value,
				saga_point,
				is_admin)
			VALUES (?, ?, ?, ?, ?, ?)`
	if tx != nil {
		var stmt *sql.Stmt
		if stmt, err = tx.Prepare(insertQuery); err != nil {
		}

		if _, err = stmt.Exec(user.Email, user.Password, user.EthAddress, user.EthValue, user.SagaPoint, user.IsAdmin); err != nil {
		}

		fmt.Println(user)
		return err
	}
	_, err = DB.Exec(insertQuery, user.Email, user.Password, user.EthAddress, user.EthValue, user.SagaPoint, user.IsAdmin)

	return err
}
