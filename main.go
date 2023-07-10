package main

import "MyBlockchain/state"

func main() {
	//state.MPTForHundredMillionOne()
	//state.MPTForHundredMillionTwo()

	for i := 1; i <= 6; i++ {
		//state.Leveldb(i)
		state.TestLeveldb(i)
	}
}
