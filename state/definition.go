package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"math/big"
)

const (
	minCache      = 2048
	minHandles    = 2048
	nativeDBPath  = "D:/Project/leveldb/abi/native leveldb"
	accountDBPath = "D:/Project/leveldb/state/account leveldb"

	rootHash     = "ROOT_HASH"
	StateDBPath1 = "D:/Project/leveldb/state/trie leveldb in 1W"
	StateDBPath2 = "D:/Project/leveldb/state/trie leveldb in 10W"
	StateDBPath3 = "D:/Project/leveldb/state/trie leveldb in 100W"
	StateDBPath4 = "D:/Project/leveldb/state/trie leveldb in 2834886"
	StateDBPath5 = "D:/Project/leveldb/state/trie leveldb in 1000W"
	StateDBPath6 = "D:/Project/leveldb/state/trie leveldb in 10000W"
)

type (
	levelDBConfig struct {
		File      string
		Cache     int
		Handles   int
		Namespace string
		Readonly  bool
	}

	//处理后的账户数据的数据格式
	accountFormat struct {
		Hash       string
		IsContract bool
	}

	account struct {
		Address  common.Address
		Nonce    uint64
		Balance  *big.Int
		Root     common.Hash // merkle root of the storage trie
		CodeHash []byte
	}
)

var (
	base     = big.NewInt(10)
	exponent = big.NewInt(21)
	Balance  = new(big.Int).Exp(base, exponent, nil)
)

var DefaultLevelDBConfig = levelDBConfig{
	File:      "levelDB",
	Cache:     2024,
	Handles:   5120,
	Namespace: "userDB",
	Readonly:  false,
}

var DefaultOpenOptions = rawdb.OpenOptions{
	Type:              "leveldb",
	Directory:         "levelDB",
	AncientsDirectory: "",
	Namespace:         "userDB",
	Cache:             2048,
	Handles:           5120,
	ReadOnly:          false,
}
