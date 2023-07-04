package main

import "MyBlockchain/state"

func main() {
	for i := 1; i <= 6; i++ {
		state.Leveldb(i)
		state.TestLeveldb(i)
	}
}
