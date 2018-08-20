package common

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullInt64 - type for sql.NullInt64
type NullInt64 sql.NullInt64

// NullFloat64 - type for sql.NullFloat64
type NullFloat64 sql.NullFloat64

// Scan for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// Scan for NullInt64
func (ni *NullFloat64) Scan(value interface{}) error {
	var i sql.NullFloat64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullFloat64{i.Float64, false}
	} else {
		*ni = NullFloat64{i.Float64, true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni NullFloat64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Float64)
}
