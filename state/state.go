package state

import (
	"MyBlockchain/state/config"
	"MyBlockchain/state/database"
	"fmt"
	"github.com/DarcyWep/pureData"
	"github.com/DarcyWep/pureData/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"os"
	"reflect"
	"time"
)

func GetAllAccounts() {
	db1, _ := openLeveldb(nativeDBPath, true)
	defer db1.Close()

	db2, _ := openLeveldb(accountDBPath, false)
	defer db2.Close()

	batch := new(leveldb.Batch)
	mapping := make(map[string]string)
	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i := min; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, _ := pureData.GetTransactionsByNumber(db1, i)

		for _, tx := range txs {
			if tx != nil {
				if tx.To != nil {
					_, ok := mapping[tx.To.String()]
					if !ok {
						mapping[tx.To.String()] = ""
						batch.Put([]byte(tx.To.String()), []byte("0"))
					}
				}
				if tx.From != nil {
					_, ok := mapping[tx.From.String()]
					if !ok {
						mapping[tx.From.String()] = ""
						batch.Put([]byte(tx.From.String()), []byte("0"))
					}
				}
			}
		}

		db2.Write(batch, nil)
		batch.Reset()
		fmt.Println("Block", i.String(), ": All accounts is saved in leveldb!")
	}

	// The number of accounts in all blocks is 2715684
	fmt.Println("The number of accounts in all blocks is", len(mapping))
}

func ShowAllAccounts() {
	db, _ := openLeveldb(accountDBPath, true)
	defer db.Close()

	// 创建迭代器
	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	// 遍历键值对

	for count := 1; iter.Next(); {
		key := string(iter.Key())
		value := string(iter.Value())
		// 处理键值对，例如打印到控制台
		log.Printf("%d: Key: %s, Value: %s", count, key, value)
		count += 1
	}

	// 检查迭代器错误
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}
}

func StateDB() {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	//sdb := database.NewStateDB(common.HexToHash("0x9ae8603e271652576a83b33908facc1780e237e553eb602b43c7183116d7bd51"), database.NewStateCache(db), nil)

	sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
	hash, _ := sdb.Commit(false)
	fmt.Println(hash)

	balance := sdb.GetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	fmt.Println(balance)

	sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)
}

func MPTForHundredMillionOne() {
	db, _ := database.OpenDatabaseWithFreezerAndSwitch(&config.DefaultsEthConfig, 6)
	defer db.Close()

	sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	//sdb := database.NewStateDB(common.HexToHash("0xa18bd9c951b0d0fd85e3692716a2e60cde7037044aaa886c3be2e501e7378264"), database.NewStateCache(db), nil)

	accountdb, _ := openLeveldb(StateDBPath6+"/accounts", false)
	defer accountdb.Close()

	file, _ := os.OpenFile(StateDBPath6+"/output1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	// 创建迭代器
	iter := accountdb.NewIterator(nil, nil)
	defer iter.Release()

	log.Println("The number of accounts has achieved", 0, "The time of the commitment is", 0)
	fmt.Fprintln(file, time.Now().Format("2006-01-02 15:04:05.000000000"), "The number of accounts has achieved", 0, "The time of the commitment is", 0)

	// 遍历键值对
	count := 1
	for iter.Next() {
		key := string(iter.Key())
		sdb.SetBalance(common.HexToAddress(key), Balance)

		if count%100000 == 0 {
			startTime := time.Now()
			hash, _ := sdb.Commit(false)
			sdb.Database().TrieDB().Commit(hash, false)
			distance := time.Since(startTime).Nanoseconds()
			log.Println("The number of accounts has achieved", count, "The time of the commitment is", distance)
			fmt.Fprintln(file, time.Now().Format("2006-01-02 15:04:05.000000000"), "The number of accounts has achieved", count, "The time of the commitment is", distance)
		}

		count += 1
	}

	// 检查迭代器错误
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}

	hash, _ := sdb.Commit(false)
	sdb.Database().TrieDB().Commit(hash, false)
	accountdb.Put([]byte(rootHash), []byte(hash.String()), nil)
	log.Println(rootHash, hash.String())
	fmt.Fprintln(file, time.Now().Format("2006-01-02 15:04:05.000000000"), rootHash, hash.String())
}

func MPTForHundredMillionTwo() {
	db, _ := database.OpenDatabaseWithFreezerAndSwitch(&config.DefaultsEthConfig, 6)
	defer db.Close()

	accountdb, _ := openLeveldb(StateDBPath6+"/accounts", false)
	defer accountdb.Close()

	root_hash, _ := accountdb.Get([]byte(rootHash), nil)
	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash(string(root_hash)), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	file, _ := os.OpenFile(StateDBPath6+"/output2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i := min; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, _ := pureData.GetTransactionsByNumber(nativedb, i)

		for _, tx := range txs {
			if tx != nil && len(tx.Transfers) > 0 {
				for _, trs := range tx.Transfers {
					if trs.GetLabel() == 1 {
						value, _ := trs.(*transaction.StorageTransition)
						hash := sdb.GetState(value.Contract, value.Slot)
						if reflect.DeepEqual(hash, common.Hash{}) {
							sdb.SetState(value.Contract, value.Slot, value.PreValue)
							//fmt.Println("Address", value.Contract, "Slot", value.Slot, "Value", value.PreValue)
						}
					}
				}
			}
		}

		startTime := time.Now()
		hash, _ := sdb.Commit(false)
		sdb.Database().TrieDB().Commit(hash, false)
		difference := time.Since(startTime).Nanoseconds()

		log.Println("Block", i.String(), "is completed!", "The time of the commitment is", difference)
		fmt.Fprintln(file, time.Now().Format("2006-01-02 15:04:05.000000000"), "Block", i.String(), "is completed!", "The time of the commitment is", difference)
	}

	hash, _ := sdb.Commit(false)
	sdb.Database().TrieDB().Commit(hash, false)
	accountdb.Put([]byte(rootHash), []byte(hash.String()), nil)
	log.Println(rootHash, hash.String())
	fmt.Fprintln(file, time.Now().Format("2006-01-02 15:04:05.000000000"), rootHash, hash.String())
}

func TestForHundredMillion() {
	db, err := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0x5be16d8197509546ee5ef661a2c8c06ca088d1daa1528696dfa91bf8f7935b68"), database.NewStateCache(db), nil)

	accountdb, _ := openLeveldb(StateDBPath6+"/accounts", true)
	defer accountdb.Close()

	file, _ := os.OpenFile(StateDBPath6+"/test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 创建迭代器
	iter := accountdb.NewIterator(nil, nil)
	defer iter.Release()

	log.Println("The number of checked accounts has achieved", 0)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of checked accounts has achieved", 0)

	// 遍历键值对
	count := 1
	for iter.Next() {
		key := string(iter.Key())
		if key == rootHash {
			continue
		}

		balance := sdb.GetBalance(common.HexToAddress(key))
		if balance.Cmp(Balance) != 0 {
			log.Println("Address", key, "Balance", balance)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Address", key, "Balance", balance)
		}

		if count%100000 == 0 {
			hash, _ := sdb.Commit(false)
			sdb.Database().TrieDB().Commit(hash, false)
			log.Println("The number of checked accounts has achieved", count)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of checked accounts has achieved", count)
		}
		count += 1
	}

	// 检查迭代器错误
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}
}

func Leveldb(i int) {
	var (
		path   string
		number int
		name   string
	)

	switch i {
	case 1:
		{
			path = StateDBPath1
			number = 1000
			name = "trie_leveldb_in_1W"
		}
	case 2:
		{
			path = StateDBPath2
			number = 10000
			name = "trie_leveldb_in_10W"
		}
	case 3:
		{
			path = StateDBPath3
			number = 100000
			name = "trie_leveldb_in_100W"
		}
	case 4:
		{
			path = StateDBPath4
			number = 100000
			name = "trie_leveldb_in_2834886"
		}
	case 5:
		{
			path = StateDBPath5
			number = 100000
			name = "trie_leveldb_in_1000W"
		}
	case 6:
		{
			path = StateDBPath6
			number = 100000
			name = "trie_leveldb_in_10000W"
		}
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The processing of", name, "is started!")

	_, err := os.Stat(path + "/experiment/process.txt")
	if err != nil {
		accountdb1, _ := openLeveldb(path+"/accounts", true)
		iter := accountdb1.NewIterator(nil, nil)

		accountdb2, _ := openLeveldb(path+"/random accounts", false)

		file1, _ := os.OpenFile(path+"/experiment/process.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		log.SetOutput(file1)

		count := 1
		batch := new(leveldb.Batch)
		for iter.Next() {
			key := string(iter.Key())
			value := string(iter.Value())

			if key == rootHash {
				batch.Put([]byte(key), []byte(value))
				continue
			} else {
				newvalue := []byte(key)
				key = common.Bytes2Hex(crypto.Keccak256([]byte(key)))
				batch.Put([]byte("0x"+key), newvalue)

				count += 1
				if count%number == 0 {
					log.Println("The number of changed accounts has achieved", count)
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of checked accounts has achieved", count)
					accountdb2.Write(batch, nil)
					batch.Reset()
				}
			}
		}

		accountdb2.Write(batch, nil)
		batch.Reset()

		if (count-1)%number != 0 {
			log.Println("The number of changed accounts has achieved", count)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of checked accounts has achieved", count)
		}

		file1.Close()
		iter.Release()
		accountdb1.Close()
		accountdb2.Close()
	}

	//_, err = os.Stat(path + "/Process2.txt")
	//if err != nil {
	//	accountdb2, _ := openLeveldb(path+"/accounts", true)
	//	iter2 := accountdb2.NewIterator(nil, nil)
	//
	//	count := 1
	//
	//	file2, _ := os.OpenFile(path+"/Process2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//	log.SetOutput(file2)
	//
	//	db, _ := database.OpenDatabaseWithFreezerAndSwitch(&config.DefaultsEthConfig, i)
	//	defer db.Close()
	//
	//	hash, _ := accountdb2.Get([]byte(rootHash), nil)
	//	sdb := database.NewStateDB(common.HexToHash(string(hash)), database.NewStateCache(db), nil)
	//
	//	for iter2.Next() {
	//		key := string(iter2.Key())
	//		value := string(iter2.Value())
	//
	//		if key == rootHash {
	//			continue
	//		} else {
	//			balance := sdb.GetBalance(common.HexToAddress(value))
	//			if balance.Cmp(Balance) != 0 {
	//				log.Println("Address", value, "Balance", balance)
	//			}
	//			if count%number == 0 {
	//				log.Println("The number of checked accounts has achieved", count)
	//				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of checked accounts has achieved", count)
	//			}
	//		}
	//
	//		count += 1
	//	}
	//
	//	if (count-1)%number != 0 {
	//		log.Println("The number of changed accounts has achieved", count)
	//		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of checked accounts has achieved", count)
	//	}
	//
	//	file2.Close()
	//	iter2.Release()
	//	accountdb2.Close()
	//}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The processing of", name, "is completed!")
	fmt.Println()
}

func TestLeveldb(i int) {
	var (
		path   string
		number int
		name   string
	)

	switch i {
	case 1:
		{
			path = StateDBPath1
			number = 10000
			name = "trie_leveldb_in_1W"
		}
	case 2:
		{
			path = StateDBPath2
			number = 100000
			name = "trie_leveldb_in_10W"
		}
	case 3:
		{
			path = StateDBPath3
			number = 100000
			name = "trie_leveldb_in_100W"
		}
	case 4:
		{
			path = StateDBPath4
			number = 100000
			name = "trie_leveldb_in_2834886"
		}
	case 5:
		{
			path = StateDBPath5
			number = 100000
			name = "trie_leveldb_in_1000W"
		}
	case 6:
		{
			path = StateDBPath6
			number = 100000
			name = "trie_leveldb_in_10000W"
		}
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The processing of", name, "is started!")

	accountdb, _ := openLeveldb(path+"/accounts", true)
	iter := accountdb.NewIterator(nil, nil)

	file, _ := os.OpenFile(path+"/experiment/sequential read result.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(file)

	db, _ := database.OpenDatabaseWithFreezerAndSwitch(&config.DefaultsEthConfig, i)
	defer db.Close()

	hash, _ := accountdb.Get([]byte(rootHash), nil)
	sdb := database.NewStateDB(common.HexToHash(string(hash)), database.NewStateCache(db), nil)

	count := 1
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())

		if key == rootHash {
			continue
		} else {
			startTime := time.Now()
			sdb.GetBalance(common.HexToAddress(value))
			log.Println(value, time.Since(startTime).Nanoseconds())
		}

		if count == number {
			break
		}

		count += 1
	}

	file.Close()
	iter.Release()
	accountdb.Close()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The processing of", name, "is completed!")
	fmt.Println("------------------------------------------------------------------------------------------------")
}

func TestLeveldb1(i int) {
	var (
		path   string
		number int
		name   string
	)

	switch i {
	case 1:
		{
			path = StateDBPath1
			number = 10000
			name = "trie_leveldb_in_1W"
		}
	case 2:
		{
			path = StateDBPath2
			number = 100000
			name = "trie_leveldb_in_10W"
		}
	case 3:
		{
			path = StateDBPath3
			number = 100000
			name = "trie_leveldb_in_100W"
		}
	case 4:
		{
			path = StateDBPath4
			number = 100000
			name = "trie_leveldb_in_2834886"
		}
	case 5:
		{
			path = StateDBPath5
			number = 100000
			name = "trie_leveldb_in_1000W"
		}
	case 6:
		{
			path = StateDBPath6
			number = 100000
			name = "trie_leveldb_in_10000W"
		}
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The processing of", name, "is started!")

	accountdb, _ := openLeveldb(path+"/accounts", true)
	iter := accountdb.NewIterator(nil, nil)

	file1, _ := os.OpenFile(path+"/experiment/sequential read result1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(file1)

	db, _ := database.OpenDatabaseWithFreezerAndSwitch(&config.DefaultsEthConfig, i)
	defer db.Close()

	hash, _ := accountdb.Get([]byte(rootHash), nil)
	sdb := database.NewStateDB(common.HexToHash(string(hash)), database.NewStateCache(db), nil)

	count := 1
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())

		if key == rootHash {
			continue
		} else {
			startTime := time.Now()
			sdb.GetBalance(common.HexToAddress(value))
			log.Println(value, time.Since(startTime).Nanoseconds())
		}

		if count == number {
			break
		}

		count += 1
	}

	file2, _ := os.OpenFile(path+"/experiment/sequential read result2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(file2)

	count = 1
	iter = accountdb.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())

		if key == rootHash {
			continue
		} else {
			startTime := time.Now()
			sdb.GetBalance(common.HexToAddress(value))
			log.Println(value, time.Since(startTime).Nanoseconds())
		}

		if count == number {
			break
		}

		count += 1
	}

	file1.Close()
	file2.Close()
	iter.Release()
	accountdb.Close()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The processing of", name, "is completed!")
	fmt.Println()
}
