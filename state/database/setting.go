package database

import (
	"runtime"
)

const (
	StateDBPath1 = "levelDB/trie leveldb in 1W"
	StateDBPath2 = "levelDB/trie leveldb in 10W"
	StateDBPath3 = "levelDB/trie leveldb in 100W"
	StateDBPath4 = "levelDB/trie leveldb in 2834886"
	StateDBPath5 = "levelDB/trie leveldb in 1000W"
	StateDBPath6 = "levelDB/trie leveldb in 10000W"
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
	path := StateDBPath4
	if runtime.GOOS == "darwin" { // MacOS
		return &RawConfig{
			Path:      path,
			Cache:     2048,
			Handles:   5120,
			Ancient:   path + "/ancient",
			Namespace: "state/" + path,
			ReadOnly:  false,
		}
	} else {
		return &RawConfig{
			Path:      path,
			Cache:     2048,
			Handles:   5120,
			Ancient:   path + "/ancient",
			Namespace: "state/" + path,
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
