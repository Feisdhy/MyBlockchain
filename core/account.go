package core

import (
	"math/big"
)

type Account struct {
	Nonce   uint64
	Balance *big.Int
}
