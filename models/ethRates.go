package models

import "api-server/common"

// Ethrate - struct for database
type Ethrate struct {
	ID     int                `db:"id" json:"id"`
	Symbol string             `db:"symbol" json:"symbol"`
	Price  common.NullFloat64 `db:"price" json:"price"`
	Time   int                `db:"time" json:"time"`
}

// FindEthrateBySymbol - find ethrate with symbol
func FindEthrateBySymbol(symbol string) (ethrate Ethrate, err error) {
	err = DB.Get(&ethrate, `SELECT * FROM eth_rates WHERE symbol=?`, symbol)
	return
}
