package state

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestStateDB(t *testing.T) {
	db1, _ := rawdb.NewLevelDBDatabase(
		DefaultLevelDBConfig.File,
		DefaultLevelDBConfig.Cache,
		DefaultOpenOptions.Handles,
		DefaultLevelDBConfig.Namespace,
		DefaultLevelDBConfig.Readonly)
	defer db1.Close()

	//db, _ := rawdb.Open(DefaultOpenOptions)
	//defer db.Close()

	sdb1, _ := state.New(types.EmptyRootHash, state.NewDatabase(db1), nil)
	defer sdb1.Database().TrieDB().Commit(sdb1.IntermediateRoot(false), false)

	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	//sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))

	//hash, _ := sdb.Commit(false)
	//fmt.Println(hash)

	balance1 := sdb1.GetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	fmt.Println(balance1)

	//sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
}
