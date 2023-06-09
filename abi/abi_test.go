package abi

import (
	"encoding/json"
	"fmt"
	"github.com/DarcyWep/pureData"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nanmu42/etherscan-api"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"strings"
	"testing"
)

func TestGetContract(t *testing.T) {
	db, err := openLeveldb(nativeDbPath, true) // get native transaction or merge transaction
	if err != nil {
		fmt.Println("open leveldb error,", err)
		return
	}

	file, err := openLeveldb(contractDbPath, false)
	if err != nil {
		fmt.Println("open leveldb error,", err)
		return
	}

	mapping := make(map[string]string)
	value := ""

	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i, count := min, 0; i.Cmp(max) == -1; i = i.Add(i, addSpan) {
		txs, err := pureData.GetTransactionsByNumber(db, i)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, tx := range txs {
			if tx.Contract && len(tx.Input) >= 4 && tx.To != nil {
				contract := tx.To.String()
				_, ok := mapping[contract]

				if !ok {
					mapping[contract] = ""
					count = count + 1
					value = value + contract + " "
					log.Printf("%d: Contract %s in Block %s getted successfully!\n", count, contract, i.String())
				}
			}
		}
	}

	value = value[:len(value)-1]
	// 写入键值对到 LevelDB
	err = file.Put([]byte(contractLeveldbKey), []byte(value), nil)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	db.Close()
}

func TestGetABI(t *testing.T) {
	file1, err := openLeveldb(contractDbPath, true)
	if err != nil {
		fmt.Println("open leveldb error,", err)
		return
	}

	file2, err := openLeveldb(abiDbPath, false)
	if err != nil {
		fmt.Println("open leveldb error,", err)
		return
	}

	//创建连接指定网络的客户端
	client := etherscan.New(etherscan.Mainnet, "D5YVTPXBBYJCKGGVK9VRPJQBFCNXPFZ1AK")
	contract, err := file1.Get([]byte(contractLeveldbKey), nil)
	if err != nil {
		fmt.Println("read leveldb err,", err)
		return
	}
	contracts := strings.Split(string(contract), " ")

	count := 0
	batch := new(leveldb.Batch)
	for _, address := range contracts {
		count = count + 1
		value, _ := file2.Get([]byte(address), nil)
		if len(value) != 0 {
			log.Printf("%d: Contract %s abi is already saved in levelDB!\n", count, address)
			continue
		} else {
			str, err := client.ContractABI(address)
			if err != nil {
				log.Printf("ERROR %d: Contract %s abi is not public!\n", count, address)
				batch.Put([]byte(address), []byte("Private Contract!"))
			} else {
				batch.Put([]byte(address), []byte(str))
				log.Printf("%d: Contract %s abi is saved successfully!\n", count, address)
			}
		}
		if count%100 == 0 {
			err := file2.Write(batch, nil)
			if err != nil {
				fmt.Println("levelDB write err,", err)
				return
			}
			batch.Reset()
		}
	}

	err = file2.Write(batch, nil)
	if err != nil {
		fmt.Println("levelDB write err,", err)
		return
	}
	batch.Reset()

	file1.Close()
	file2.Close()
}

func TestHandleAllBlocks(t *testing.T) {
	db, err := openLeveldb(nativeDbPath, true)
	if err != nil {
		fmt.Println("open levelDB err,", err)
		return
	}

	file1, err := openLeveldb(abiDbPath, true)
	if err != nil {
		fmt.Println("open levelDB err,", err)
		return
	}

	file2, err := openLeveldb(transactionDbPath, false)
	if err != nil {
		fmt.Println("open levelDB err,", err)
		return
	}

	defer db.Close()
	defer file1.Close()
	defer file2.Close()

	batch := new(leveldb.Batch)
	min, max, addSpan := big.NewInt(12000001), big.NewInt(12050000), big.NewInt(1)
	for i := min; i.Cmp(max) == -1; i.Add(i, addSpan) {
		txs, err := pureData.GetTransactionsByNumber(db, i)
		if err != nil {
			fmt.Println("read levelDB err,", err)
			return
		}

		_, err = file2.Get([]byte(i.String()), nil)
		if err == nil {
			log.Printf("Block %s is already saved in levelDB!", i.String())
			continue
		}

		key := []byte(i.String())
		transactions := make([]levelDBFormat, 0)
		for _, tx := range txs {
			if tx.Contract && len(tx.Input) >= 4 && tx.To != nil {
				str, err := file1.Get([]byte(tx.To.String()), nil)
				if err != nil {
					fmt.Println("read levelDB err,", err)
					return
				} else if string(str) == "Private Contract!" {
					continue
				} else {
					contractabi, err := abi.JSON(strings.NewReader(string(str)))
					if err != nil {
						continue
					}
					method, err := contractabi.MethodById(tx.Input[:4])
					if err != nil {
						transactions = append(transactions, levelDBFormat{tx.Hash.String(), tx.To.String(), "0x" + common.Bytes2Hex(tx.Input[:4]), ""})
						//log.Printf("Transaction %s in Block %s misses method!", tx.Hash.String(), i.String())
					} else {
						transactions = append(transactions, levelDBFormat{tx.Hash.String(), tx.To.String(), "0x" + common.Bytes2Hex(tx.Input[:4]), method.Sig})
						//log.Printf("Transaction %s in Block %s analyzes successfully!", tx.Hash.String(), i.String())
					}
				}
			}
		}

		value, err := json.Marshal(transactions)
		if err != nil {
			fmt.Println("Block "+i.String()+"json transition err, ", err)
			return
		}
		batch.Put(key, value)
		err = file2.Write(batch, nil)
		if err != nil {
			fmt.Println("levelDB write err,", err)
			return
		}
		batch.Reset()
		log.Printf("Block %s is saved in levelDB successfully!", i.String())
	}
}
