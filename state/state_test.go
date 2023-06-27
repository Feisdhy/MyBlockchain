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
	"reflect"
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
			if tx != nil && len(tx.Transfers) > 0 {
				for _, trs := range tx.Transfers {
					if trs.GetLabel() == 0 {
						value, _ := trs.(*transaction.StateTransition)

						if value.From != nil {
							_, ok := mapping[value.From.Address.String()]
							if !ok {
								mapping[value.From.Address.String()] = ""
								batch.Put([]byte(value.From.Address.String()), []byte("0"))
							}
						}

						if value.To != nil {
							_, ok := mapping[value.To.Address.String()]
							if !ok {
								mapping[value.To.Address.String()] = ""
								batch.Put([]byte(value.To.Address.String()), []byte("0"))
							}
						}

					} else {
						value, _ := trs.(*transaction.StorageTransition)
						_, ok := mapping[value.Contract.String()]
						if !ok {
							mapping[value.Contract.String()] = ""
							batch.Put([]byte(value.Contract.String()), []byte("0"))
						}
					}
				}
			}
		}

		db2.Write(batch, nil)
		batch.Reset()
		fmt.Println("Block", i.String(), ": All accounts is saved in leveldb!")
	}

	// The number of accounts in all blocks is 2715684
	// The number of accounts in all blocks is 2834886
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
	sdb := database.NewStateDB(common.HexToHash("0xeec5a7aed79ef01126db0d292ff4abd68eab16db3ff2df431e2f8254f82ab378"), database.NewStateCache(db), nil)

	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	//sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
	//hash, _ := sdb.Commit(false)
	//fmt.Println(hash)

	balance := sdb.GetBalance(common.HexToAddress("0xfffffffFf15AbF397dA76f1dcc1A1604F45126DB"))
	fmt.Println(balance)

	//sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)
}

func TestShowAccountsAndRoot(t *testing.T) {
	db, _ := openLeveldb(StateDBPath4+"/accounts", true)
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

func TestStateDBForTenThousand(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	//sdb := database.NewStateDB(common.HexToHash("0xf9f0f433e2ea6ec0e88884355cfb519b3289a09cb81d4499de48afd8c064ba69"), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	accountdb, _ := openLeveldb(StateDBPath1+"/accounts", false)
	defer accountdb.Close()

	batch := new(leveldb.Batch)
	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i, count := min, 1; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
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
						//fmt.Println(value.String())

						if value.Type == 1 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}

							balance = sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
								count += 1
							}
						} else if value.Type == 2 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}
						} else {
							balance := sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
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
						//fmt.Println(value.String())

						balance := sdb.GetBalance(value.Contract)
						if balance.Cmp(big.NewInt(0)) == 0 && count <= 10000 {
							sdb.GetOrNewStateObject(value.Contract)
							sdb.SetBalance(value.Contract, Balance)
							fmt.Println(count, "Address", value.Contract, "Balance", Balance, "Slot", value.Slot, "Value", value.PreValue)
							batch.Put([]byte(value.Contract.String()), []byte(""))
							count += 1
						}

						hash := sdb.GetState(value.Contract, value.Slot)
						if reflect.DeepEqual(hash, common.Hash{}) {
							sdb.SetState(value.Contract, value.Slot, value.PreValue)
							fmt.Println("Address", value.Contract, "Slot", value.Slot, "Value", value.PreValue)
						}
					}
				}
			} else if count > 10000 {
				break
			}
		}

		if count > 10000 {
			break
		} else {
			accountdb.Write(batch, nil)
			batch.Reset()
		}
	}

	accountdb.Write(batch, nil)
	batch.Reset()

	sdb.Commit(false)
	sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)

	accountdb.Put([]byte(rootHash), []byte(sdb.IntermediateRoot(false).String()), nil)
}

func TestStateDBForHundredThousand(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0xeec5a7aed79ef01126db0d292ff4abd68eab16db3ff2df431e2f8254f82ab378"), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	accountdb, _ := openLeveldb(StateDBPath2+"/accounts", false)
	defer accountdb.Close()

	batch := new(leveldb.Batch)
	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i, count := min, 10001; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, _ := pureData.GetTransactionsByNumber(nativedb, i)

		for _, tx := range txs {
			if tx != nil && len(tx.Transfers) > 0 && count <= 100000 {
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
						//fmt.Println(value.String())

						if value.Type == 1 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 100000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}

							balance = sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 100000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
								count += 1
							}
						} else if value.Type == 2 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 100000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}
						} else {
							balance := sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 100000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
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
						//fmt.Println(value.String())

						balance := sdb.GetBalance(value.Contract)
						if balance.Cmp(big.NewInt(0)) == 0 && count <= 100000 {
							sdb.GetOrNewStateObject(value.Contract)
							sdb.SetBalance(value.Contract, Balance)
							fmt.Println(count, "Address", value.Contract, "Balance", Balance, "Slot", value.Slot, "Value", value.PreValue)
							batch.Put([]byte(value.Contract.String()), []byte(""))
							count += 1
						}

						hash := sdb.GetState(value.Contract, value.Slot)
						if reflect.DeepEqual(hash, common.Hash{}) {
							sdb.SetState(value.Contract, value.Slot, value.PreValue)
							fmt.Println("Address", value.Contract, "Slot", value.Slot, "Value", value.PreValue)
						}
					}
				}
			} else if count > 100000 {
				break
			}
		}

		if count > 100000 {
			break
		} else {
			accountdb.Write(batch, nil)
			batch.Reset()
		}
	}

	accountdb.Write(batch, nil)
	batch.Reset()

	sdb.Commit(false)
	sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)

	accountdb.Put([]byte(rootHash), []byte(sdb.IntermediateRoot(false).String()), nil)
}

func TestStateDBForOneMillion(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0x1ee6be866b323731bd1faa7fed30945fe46871ec12a8d8a1e86033dd64ddd642"), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	accountdb, _ := openLeveldb(StateDBPath3+"/accounts", false)
	defer accountdb.Close()

	batch := new(leveldb.Batch)
	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i, count := min, 100001; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, _ := pureData.GetTransactionsByNumber(nativedb, i)

		for _, tx := range txs {
			if tx != nil && len(tx.Transfers) > 0 && count <= 1000000 {
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
						//fmt.Println(value.String())

						if value.Type == 1 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 1000000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}

							balance = sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 1000000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
								count += 1
							}
						} else if value.Type == 2 {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 1000000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println(count, "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}
						} else {
							balance := sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 1000000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println(count, "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
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
						//fmt.Println(value.String())

						balance := sdb.GetBalance(value.Contract)
						if balance.Cmp(big.NewInt(0)) == 0 && count <= 100000 {
							sdb.GetOrNewStateObject(value.Contract)
							sdb.SetBalance(value.Contract, Balance)
							fmt.Println(count, "Address", value.Contract, "Balance", Balance, "Slot", value.Slot, "Value", value.PreValue)
							batch.Put([]byte(value.Contract.String()), []byte(""))
							count += 1
						}

						hash := sdb.GetState(value.Contract, value.Slot)
						if reflect.DeepEqual(hash, common.Hash{}) {
							sdb.SetState(value.Contract, value.Slot, value.PreValue)
							//fmt.Println("Address", value.Contract, "Slot", value.Slot, "Value", value.PreValue)
						}
					}
				}
			} else if count > 1000000 {
				break
			}
		}

		if count > 1000000 {
			break
		} else {
			accountdb.Write(batch, nil)
			batch.Reset()
		}
	}

	accountdb.Write(batch, nil)
	batch.Reset()

	sdb.Commit(false)
	sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)

	accountdb.Put([]byte(rootHash), []byte(sdb.IntermediateRoot(false).String()), nil)
}

func TestStateDBForAllRealData(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	defer db.Close()

	//sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	sdb := database.NewStateDB(common.HexToHash("0x66056d2875ebef61a456dc31ef93eb16ff15e2dd9ae2bdb7c2a1235ad7bdcfa0"), database.NewStateCache(db), nil)

	nativedb, _ := openLeveldb(nativeDBPath, true)
	defer nativedb.Close()

	accountdb, _ := openLeveldb(StateDBPath4+"/accounts", false)
	defer accountdb.Close()

	count := 1000001
	batch := new(leveldb.Batch)
	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i := min; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, _ := pureData.GetTransactionsByNumber(nativedb, i)

		for _, tx := range txs {
			if tx != nil && len(tx.Transfers) > 0 && count <= 3000000 {
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

						if value.From != nil {
							balance := sdb.GetBalance(value.From.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 3000000 {
								sdb.GetOrNewStateObject(value.From.Address)
								sdb.SetBalance(value.From.Address, Balance)
								fmt.Println("Account", count, "Block", i.String(), "Address", value.From.Address, "Balance", Balance)
								batch.Put([]byte(value.From.Address.String()), []byte(""))
								count += 1
							}
						}

						if value.To != nil {
							balance := sdb.GetBalance(value.To.Address)
							if balance.Cmp(big.NewInt(0)) == 0 && count <= 3000000 {
								sdb.GetOrNewStateObject(value.To.Address)
								sdb.SetBalance(value.To.Address, Balance)
								fmt.Println("Account", count, "Block", i.String(), "Address", value.To.Address, "Balance", Balance)
								batch.Put([]byte(value.To.Address.String()), []byte(""))
								count += 1
							}
						}

					} else {
						value, _ := trs.(*transaction.StorageTransition)
						balance := sdb.GetBalance(value.Contract)
						if balance.Cmp(big.NewInt(0)) == 0 && count <= 3000000 {
							sdb.GetOrNewStateObject(value.Contract)
							sdb.SetBalance(value.Contract, Balance)
							fmt.Println("Account", count, "Block", i.String(), "Address", value.Contract, "Balance", Balance, "Slot", value.Slot, "Value", value.PreValue)
							batch.Put([]byte(value.Contract.String()), []byte(""))
							count += 1
						}

						hash := sdb.GetState(value.Contract, value.Slot)
						if reflect.DeepEqual(hash, common.Hash{}) {
							sdb.SetState(value.Contract, value.Slot, value.PreValue)
							//fmt.Println("Address", value.Contract, "Slot", value.Slot, "Value", value.PreValue)
						}
					}
				}
			} else if count > 3000000 {
				break
			}
		}

		if count > 3000000 {
			break
		} else {
			accountdb.Write(batch, nil)
			batch.Reset()

			//sdb.Database().TrieDB().Reference(sdb.IntermediateRoot(false), common.Hash{}) // metadata reference to keep trie alive
			//// If we exceeded our memory allowance, flush matured singleton nodes to disk
			//var (
			//	nodes, imgs = sdb.Database().TrieDB().Size()
			//	limit       = common.StorageSize(256) * 1024 * 1024
			//)
			//if nodes > limit || imgs > 4*1024*1024 {
			//	sdb.Database().TrieDB().Cap(limit - ethdb.IdealBatchSize)
			//}
			//sdb.Database().TrieDB().Dereference(sdb.IntermediateRoot(false))

			hash, _ := sdb.Commit(false)
			sdb.Database().TrieDB().Commit(hash, false)

			//tmp := new(big.Int).Set(i)
			//if tmp.Mod(tmp, big.NewInt(1000)).Cmp(big.NewInt(0)) == 0 {
			//	sdb = database.NewStateDB(hash, database.NewStateCache(db), nil)
			//}
		}
	}

	accountdb.Write(batch, nil)
	batch.Reset()

	hash, _ := sdb.Commit(false)
	sdb.Database().TrieDB().Commit(hash, false)

	fmt.Println("The number of Accounts has achieved", count, "!")
	accountdb.Put([]byte(rootHash), []byte(sdb.IntermediateRoot(false).String()), nil)
}
