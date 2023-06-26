package state

import (
	"MyBlockchain/state/config"
	"MyBlockchain/state/database"
	"bufio"
	"fmt"
	"github.com/DarcyWep/pureData"
	"github.com/DarcyWep/pureData/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"os"
	"testing"
)

func TestGetAllAccounts(t *testing.T) {
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

func TestShowAllAccounts(t *testing.T) {
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

func TestStateDB(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0xb8c439e2ac915b4085960c76d02148ce6c4b47ca7fa77089275e81b04047fb2c"), database.NewStateCache(db), nil)

	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	//sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
	//hash, _ := sdb.Commit(false)
	//fmt.Println(hash)

	balance := sdb.GetBalance(common.HexToAddress("0xFB5426d0BFc52a1418bC02fe34396123e0543e24"))
	fmt.Println(balance)

	//sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)
}

func TestStateDBForTenThousand(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	//sdb := database.NewStateDB(common.HexToHash("0xf9f0f433e2ea6ec0e88884355cfb519b3289a09cb81d4499de48afd8c064ba69"), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	file, _ := os.OpenFile(database.StateDBPath1+"/Root_Hash.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i, count := min, 0; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, _ := pureData.GetTransactionsByNumber(nativedb, i)

		for _, tx := range txs {
			if tx != nil && len(tx.Transfers) > 0 && count <= 10000 {
				for _, trs := range tx.Transfers {
					if trs.GetLabel() == 0 {
						/*
							type StateTransition struct {
								Label uint8 // 0: state, 1: storage
								// type 类型(1: 转账; 2: 手续费扣除, 只有From字段; 3: 手续费添加给矿工, 只有To字段 ; 4: 合约销毁; 5: 矿工奖励, 只有To字段)
								// 类型 5(矿工奖励) 每个区块只有一个记录
								Type uint8 // (手续费扣除 2!=3 给矿工的手续费)

								From  *Balance
								To    *Balance
								Value *big.Int
							}
						*/

						value, _ := trs.(*transaction.StateTransition)
						//fmt.Println(tx.Hash, value.String())

						if value.Type == 1 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From, "Balance", Balance)
								count += 1
							}

							balance = sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To, "Balance", Balance)
								count += 1
							}
						} else if value.Type == 2 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From, "Balance", Balance)
								count += 1
							}
						} else {
							balance := sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To, "Balance", Balance)
								count += 1
							}
						}
					} else {
						/*
							type StorageTransition struct {
								Label uint8 // 0: state, 1: storage

								Contract common.Address
								Slot     common.Hash // 智能合约的存储槽
								PreValue common.Hash
								NewValue *common.Hash // newValue = nil 则是 SLOAD, 否则为 SSTORE
							}
						*/
						value, _ := trs.(*transaction.StorageTransition)
						//fmt.Println(tx.Hash, value.String())

						balance := sdb.GetBalance(value.Contract)
						if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
							sdb.GetOrNewStateObject(value.Contract)
							sdb.SetBalance(value.Contract, Balance)
							sdb.SetState(value.Contract, value.Slot, value.PreValue)
							fmt.Println(count, "Address", value.Contract, "Balance", Balance, "Slot", value.Slot, "Value", value.PreValue)
							count += 1
						}
					}
				}
			} else if count > 10000 {
				break
			}
		}

		if count > 10000 {
			break
		}
	}

	sdb.Commit(false)
	sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)

	// 创建一个写入器
	writer := bufio.NewWriter(file)
	// 写入内容
	writer.WriteString(sdb.IntermediateRoot(false).String() + "\n")
	// 刷新写入器缓冲区，确保数据写入文件
	writer.Flush()
}
