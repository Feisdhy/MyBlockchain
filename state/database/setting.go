package database

import (
	"runtime"
)

const (
	StateDBPath1 = "levelDB/trie leveldb in 1W"
	StateDBPath2 = "levelDB/trie leveldb in 10W"
	StateDBPath3 = "levelDB/trie leveldb in 100W"
	StateDBPath4 = "levelDB/trie leveldb in 1000W"
	StateDBPath5 = "levelDB/trie leveldb in 10000W"
)

type RawConfig struct {
	Path      string
	Cache     int
	Handles   int
	Ancient   string
	Namespace string
	ReadOnly  bool
}

func defaultRawConfig() *RawConfig {
	if runtime.GOOS == "darwin" { // MacOS
		return &RawConfig{
			Path:      StateDBPath1,
			Cache:     2048,
			Handles:   5120,
			Ancient:   StateDBPath1 + "/ancient",
			Namespace: "state/" + StateDBPath1,
			ReadOnly:  false,
		}
	} else {
		return &RawConfig{
			Path:      StateDBPath1,
			Cache:     2048,
			Handles:   5120,
			Ancient:   StateDBPath1 + "/ancient",
			Namespace: "state/" + StateDBPath1,
			ReadOnly:  false,
		}
	}
}

type StateDBConfig struct {
	Cache     int
	Journal   string
	Preimages bool
}

func defaultStateDBConfig() *StateDBConfig {
	if runtime.GOOS == "darwin" { // MacOS
		return &StateDBConfig{
			Cache:     614,
			Journal:   "levelDB/state leveldb",
			Preimages: false,
		}
	} else {
		return &StateDBConfig{
			Cache:     614,
			Journal:   "levelDB/state leveldb",
			Preimages: false,
		}
	}
}
