package tools

import (
	crypto "crypto/rand"
	"math/big"
)

func NewCryptoRand() int {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return int(safeNum.Int64())
}
