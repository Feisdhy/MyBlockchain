package core

import (
	"math/big"
)

type Header struct {
	ParentHash Hash
	Miner      Address
	Number     *big.Int
	Time       uint64
	Nonce      *big.Int
}

type Body struct {
	Transaction []*transaction
}

type Block struct {
	header *Header
	body   *Body
}

func (block *Block) Mine() {
	target := big.NewInt(1)
	target.Lsh(target, 256-Difficulty)

}
