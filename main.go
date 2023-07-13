package main

import "MyBlockchain/state"

func main() {
	//state.MPTForHundredMillionOne()
	//state.MPTForHundredMillionTwo()

	//for i := 1; i <= 10; i++ {
	//	for j := 1; j <= 6; j++ {
	//		state.TestLeveldbSequential1(j, strconv.Itoa(i))
	//		state.TestLeveldbRandom1(j, strconv.Itoa(i))
	//		state.TestLeveldbSequential2(j, strconv.Itoa(i))
	//		state.TestLeveldbRandom2(j, strconv.Itoa(i))
	//	}
	//}

	//var (
	//	base     = big.NewInt(10)
	//	exponent = big.NewInt(21)
	//	Balance  = new(big.Int).Exp(base, exponent, nil)
	//)

	//db, _ := leveldb.OpenFile("leveldb", &opt.Options{
	//	OpenFilesCacheCapacity: 2048,
	//	BlockCacheCapacity:     2048 / 2 * opt.MiB,
	//	WriteBuffer:            2048 / 4 * opt.MiB, // Two of these are used internally
	//	ReadOnly:               false,
	//})
	//defer db.Close()

	//batch := new(leveldb.Batch)
	//
	//startTime := time.Now()
	//batch.Put([]byte("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"), []byte(Balance.String()))
	//time1 := time.Since(startTime).Nanoseconds()
	//fmt.Println(time1)
	//
	//startTime = time.Now()
	//db.Write(batch, nil)
	//time2 := time.Since(startTime).Nanoseconds()
	//fmt.Println(time2)

	//startTime := time.Now()
	//db.Get([]byte("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"), nil)
	//time1 := time.Since(startTime).Nanoseconds()
	//fmt.Println(time1)

	state.Txt2CsvForSequentialRead()
	state.Txt2CsvForSequentialWrite()
	state.Txt2CsvForRandomRead()
	state.Txt2CsvForRandomWrite()
	state.Csv2CsvForSequentialRead()
	state.Csv2CsvForSequentialWrite()
	state.Csv2CsvForRandomRead()
	state.Csv2CsvForRandomWrite()
}
