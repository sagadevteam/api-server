package common

import (
	"fmt"
	"math/big"
)

// ParseBig256 parses s as a 256 bit integer in decimal or hexadecimal syntax.
// Leading zeros are accepted. The empty string parses as zero.
func ParseBig256(s string) (*big.Int, bool) {
	if s == "" {
		return new(big.Int), true
	}
	var bigint *big.Int
	var ok bool
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		bigint, ok = new(big.Int).SetString(s[2:], 16)
	} else {
		bigint, ok = new(big.Int).SetString(s, 10)
	}
	if ok && bigint.BitLen() > 256 {
		bigint, ok = nil, false
	}
	return bigint, ok
}

func BigFloatToBigInt(bigval *big.Float) *big.Int {

	bigval.SetPrec(64)
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func BigIntToHex(n *big.Int) string {
	return fmt.Sprintf("0x%x", n) // or %X or upper case
}
