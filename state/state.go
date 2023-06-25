package state

import (
	"MyBlockchain/state/config"
	"MyBlockchain/state/database"
	"fmt"
	"github.com/DarcyWep/pureData"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
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
