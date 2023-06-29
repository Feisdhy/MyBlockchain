package state

import (
	"MyBlockchain/state/config"
	"MyBlockchain/state/database"
	"fmt"
	"github.com/DarcyWep/pureData"
	"github.com/DarcyWep/pureData/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	//sdb := database.NewStateDB(common.HexToHash("0xa18bd9c951b0d0fd85e3692716a2e60cde7037044aaa886c3be2e501e7378264"), database.NewStateCache(db), nil)

	accountdb, _ := openLeveldb(StateDBPath6+"/accounts", true)
	defer accountdb.Close()

	file, _ := os.OpenFile(StateDBPath6+"/output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 创建迭代器
	iter := accountdb.NewIterator(nil, nil)
	defer iter.Release()

	log.Println("The number of accounts has achieved", 0)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of accounts has achieved", 0)

	// 遍历键值对
	count := 1
	for iter.Next() {
		key := string(iter.Key())
		sdb.SetBalance(common.HexToAddress(key), Balance)

		if count%100000 == 0 {
			sdb.Commit(false)
			log.Println("The number of accounts has achieved", count)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "The number of accounts has achieved", count)
		}

		if count%1000000 == 0 {
			sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)
		}

		count += 1
	}

	// 检查迭代器错误
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}

	hash, _ := sdb.Commit(false)
	sdb.Database().TrieDB().Commit(hash, false)
	//accountdb.Put([]byte(rootHash), []byte(hash.String()), nil)
	log.Println(rootHash, hash.String())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), rootHash, hash.String())
}

func MPTForHundredMillionTwo() {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0x5be16d8197509546ee5ef661a2c8c06ca088d1daa1528696dfa91bf8f7935b68"), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	file, _ := os.OpenFile(StateDBPath6+"/output1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

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

		hash, _ := sdb.Commit(false)
		sdb.Database().TrieDB().Commit(hash, false)

		log.Println("Block", i.String(), "is completed!")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Block", i.String(), "is completed!")
	}

	hash, _ := sdb.Commit(false)
	sdb.Database().TrieDB().Commit(hash, false)
	log.Println(rootHash, hash.String())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), rootHash, hash.String())
}

func TestForHundredMillionOne() {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0x5be16d8197509546ee5ef661a2c8c06ca088d1daa1528696dfa91bf8f7935b68"), database.NewStateCache(db), nil)

	accountdb, _ := openLeveldb(StateDBPath6+"/accounts", true)
	defer accountdb.Close()

	file, _ := os.OpenFile(StateDBPath6+"/output2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
