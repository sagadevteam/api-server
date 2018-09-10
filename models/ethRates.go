package models

import (
	"api-server/common"
	"database/sql"
)

// Ethrate - struct for database
type Ethrate struct {
	ID     int                `db:"id" json:"id"`
	Symbol string             `db:"symbol" json:"symbol"`
	Price  common.NullFloat64 `db:"price" json:"price"`
	Time   int                `db:"time" json:"time"`
}

// FindEthrateBySymbol - find ethrate with symbol
func FindEthrateBySymbol(symbol string, tx *sql.Tx) (ethrate Ethrate, err error) {
	sqlStr := `SELECT * FROM eth_rates WHERE symbol=?`
	if tx != nil {
		err = tx.QueryRow(sqlStr, symbol).Scan(&ethrate.ID, &ethrate.Symbol, &ethrate.Price, &ethrate.Time)
	} else {
		err = DB.QueryRow(sqlStr, symbol).Scan(&ethrate.ID, &ethrate.Symbol, &ethrate.Price, &ethrate.Time)
	}
	return
}
