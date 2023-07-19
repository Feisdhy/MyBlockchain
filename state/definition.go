package state

import (
	"github.com/ethereum/go-ethereum/core/rawdb"
	"math/big"
)

const (
	minCache   = 2048
	minHandles = 2048

	nativeDBPath  = "/home/fuzh/copy4/abi/native_leveldb"
	accountDBPath = "/home/fuzh/copy4/state/trie_leveldb_in_1W/accounts"

	rootHash     = "ROOT_HASH"
	StateDBPath1 = "/home/fuzh/copy4/state/trie_leveldb_in_1W"
	StateDBPath2 = "/home/fuzh/copy4/state/trie_leveldb_in_10W"
	StateDBPath3 = "/home/fuzh/copy4/state/trie_leveldb_in_100W"
	StateDBPath4 = "/home/fuzh/copy4/state/trie_leveldb_in_2834886"
	StateDBPath5 = "/home/fuzh/copy4/state/trie_leveldb_in_1000W"
	StateDBPath6 = "/home/fuzh/copy4/state/trie_leveldb_in_10000W"

	filepath = "D:/Project/MyBlockchain/state/file"
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

	//account struct {
	//	Address  common.Address
	//	Nonce    uint64
	//	Balance  *big.Int
	//	Root     common.Hash // merkle root of the storage trie
	//	CodeHash []byte
	//}

	output struct {
		Time1 int64
		Time2 int64
	}
)

var (
	base     = big.NewInt(10)
	exponent = big.NewInt(21)
	Balance  = new(big.Int).Exp(base, exponent, nil)

	base1     = big.NewInt(10)
	exponent1 = big.NewInt(20)
	Balance1  = new(big.Int).Exp(base, exponent, nil)
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
