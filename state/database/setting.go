package database

import (
	"runtime"
)

const (
	StateDBPath1 = "/home/fuzh/leveldb/state/trie_leveldb_in_1W"
	StateDBPath2 = "/home/fuzh/leveldb/state/trie_leveldb_in_10W"
	StateDBPath3 = "/home/fuzh/leveldb/state/trie_leveldb_in_100W"
	StateDBPath4 = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886"
	StateDBPath5 = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W"
	StateDBPath6 = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W"
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
	path := StateDBPath1
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

func defaultRawConfigWithSwitch(i int) *RawConfig {
	var path string
	switch i {
	case 1:
		{
			path = StateDBPath1
		}
	case 2:
		{
			path = StateDBPath2
		}
	case 3:
		{
			path = StateDBPath3
		}
	case 4:
		{
			path = StateDBPath4
		}
	case 5:
		{
			path = StateDBPath5
		}
	case 6:
		{
			path = StateDBPath6
		}
	}

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
