package state

import (
	"github.com/ethereum/go-ethereum/core/rawdb"
)

const (
	minCache    = 2048
	minHandles  = 2048
	levelDBPath = "levelDB"
)

type levelDBConfig struct {
	File      string
	Cache     int
	Handles   int
	Namespace string
	Readonly  bool
}

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
