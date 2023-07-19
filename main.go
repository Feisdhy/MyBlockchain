package main

import "MyBlockchain/state"

func main() {
	//state.MPTForHundredMillionOne()
	//state.MPTForHundredMillionTwo()
	//
	//for i := 1; i <= 6; i++ {
	//	state.Leveldb(i)
	//}

	//for i := 3; i <= 3; i++ {
	//	for j := 1; j <= 6; j++ {
	//		//state.TestLeveldbSequential1(j, strconv.Itoa(i))
	//		//state.TestLeveldbRandom1(j, strconv.Itoa(i))
	//		//state.TestLeveldbSequential2(j, strconv.Itoa(i))
	//		state.TestLeveldbRandom2(j, strconv.Itoa(i))
	//	}
	//}

	//state.Txt2CsvForSequentialRead()
	//state.Txt2CsvForSequentialWrite()
	//state.Txt2CsvForRandomRead()
	state.Txt2CsvForRandomWrite()
	//state.Csv2CsvForSequentialRead()
	//state.Csv2CsvForSequentialWrite()
	//state.Csv2CsvForRandomRead()
	state.Csv2CsvForRandomWrite()
}
