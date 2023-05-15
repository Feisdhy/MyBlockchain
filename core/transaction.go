package core

import "math/big"

type transaction struct {
	from  Address
	to    Address
	value *big.Int
}
