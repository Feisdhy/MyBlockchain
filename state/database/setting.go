package database

import "runtime"

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
			Path:      "levelDB",
			Cache:     2048,
			Handles:   5120,
			Ancient:   "levelDB/ancient",
			Namespace: "state/levelDB",
			ReadOnly:  false,
		}
	} else {
		return &RawConfig{
			Path:      "levelDB",
			Cache:     2048,
			Handles:   5120,
			Ancient:   "levelDB/ancient",
			Namespace: "state/levelDB",
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
			Journal:   "levelDB/stateData",
			Preimages: false,
		}
	} else {
		return &StateDBConfig{
			Cache:     614,
			Journal:   "levelDB/stateData",
			Preimages: false,
		}
	}
}
