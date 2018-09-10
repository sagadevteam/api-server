package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

// ErrorMsgSagaPointNotEnough - error message of admin saga point not enough
const ErrorMsgSagaPointNotEnough = "admin saga point not enough"

// User is user schema in mysql
type User struct {
	UserID        int    `db:"user_id" json:"user_id"`
	Email         string `db:"email" json:"email"`
	Password      string `db:"password" json:"-"`
	EthAddress    string `db:"eth_addr" json:"eth_addr"`
	EthPrivateKey string `db:"eth_priv" json:"-"`
	EthValue      string `db:"eth_value" json:"eth_value"`
	SagaPoint     int    `db:"saga_point" json:"saga_point"`
	IsAdmin       int    `db:"is_admin" json:"is_admin"`
}

// FindUserByEmail find user by email
func FindUserByEmail(email string, columns []string, tx *sql.Tx) (user User, err error) {
	sql := fmt.Sprintf(`SELECT %s FROM users WHERE email=? limit 1`, strings.Join(columns[:], ","))
	if tx != nil {
		err = tx.QueryRow(sql, email).
			Scan(&user.UserID, &user.Email, &user.Password, &user.EthAddress, &user.EthPrivateKey, &user.EthValue, &user.SagaPoint, &user.IsAdmin)
		return
	}
	err = DB.Get(&user, sql, email)
	return
}

// FindUserByID find user by user_id
func FindUserByID(userID int, columns []string, tx *sql.Tx) (user User, err error) {
	sql := fmt.Sprintf(`SELECT %s FROM users WHERE user_id=? limit 1`, strings.Join(columns[:], ","))
	if tx != nil {
		err = tx.QueryRow(sql, userID).
			Scan(&user.UserID, &user.Email, &user.Password, &user.EthAddress, &user.EthPrivateKey, &user.EthValue, &user.SagaPoint, &user.IsAdmin)
		return
	}
	err = DB.Get(&user, sql, userID)
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

		if err != nil {
			return err
		}

		userID64, _ := result.LastInsertId()
		user.UserID = int(userID64)

		return err
	}
	_, err = DB.Exec(insertQuery, user.Email, user.Password, user.EthAddress, user.EthValue, user.SagaPoint, user.IsAdmin)

	return
}

// Update user
func (user *User) Update(tx *sql.Tx) (err error) {
	updateQuery := `UPDATE users SET
						email=?,
						password=?,
						eth_addr=?,
						eth_priv=?,
						eth_value=?,
						saga_point=?,
						is_admin=?
					WHERE user_id=?`
	if tx != nil {
		_, err = tx.Exec(updateQuery, user.Email, user.Password, user.EthAddress, user.EthPrivateKey, user.EthValue, user.SagaPoint, user.IsAdmin, user.UserID)
		return
	}
	_, err = DB.Exec(updateQuery, user.Password, user.EthAddress, user.EthPrivateKey, user.EthValue, user.SagaPoint, user.IsAdmin, user.UserID)

	return
}

// SelectAndUpdateAdminWithMinusSagaPoint - check saga point enough and update
func SelectAndUpdateAdminWithMinusSagaPoint(sagaPoint int, tx *sql.Tx) (err error) {
	var adminID, adminSaga int
	selectQuery := `SELECT user_id,saga_point FROM users WHERE is_admin=1`
	if tx != nil {
		err = tx.QueryRow(selectQuery).Scan(&adminID, &adminSaga)
	} else {
		err = DB.QueryRow(selectQuery).Scan(&adminID, &adminSaga)
	}
	if err != nil {
		return
	}
	if adminSaga < sagaPoint {
		err = errors.New(ErrorMsgSagaPointNotEnough)
		return
	}
	adminSaga -= sagaPoint
	updateQuery := `UPDATE users SET saga_point=? WHERE user_id=?`
	if tx != nil {
		_, err = tx.Exec(updateQuery, adminSaga, adminID)
	} else {
		_, err = DB.Exec(updateQuery, adminSaga, adminID)
	}

	return
}
