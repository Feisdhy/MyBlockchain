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
	db, _ := rawdb.NewLevelDBDatabase(
		DefaultLevelDBConfig.File,
		DefaultLevelDBConfig.Cache,
		DefaultOpenOptions.Handles,
		DefaultLevelDBConfig.Namespace,
		DefaultLevelDBConfig.Readonly)
	defer db.Close()

	//db, _ := rawdb.Open(DefaultOpenOptions)
	//defer db.Close()

	sdb, _ := state.New(types.EmptyRootHash, state.NewDatabase(db), nil)
	defer sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)

	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	//sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
	//hash, _ := sdb.Commit(false)
	//fmt.Println(hash)

	balance := sdb.GetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	fmt.Println(balance)

	//sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
}
