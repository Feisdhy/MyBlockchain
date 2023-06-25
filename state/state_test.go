package state

import (
	"MyBlockchain/state/config"
	"MyBlockchain/state/database"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

func TestGetAllAccounts(t *testing.T) {

}

func TestStateDB(t *testing.T) {
	db, _ := database.OpenDatabaseWithFreezer(&config.DefaultsEthConfig)
	sdb := database.NewStateDB(types.EmptyRootHash, database.NewStateCache(db), nil)
	//sdb := database.NewStateDB(common.HexToHash("0x9ae8603e271652576a83b33908facc1780e237e553eb602b43c7183116d7bd51"), database.NewStateCache(db), nil)

	sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	sdb.SetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"), big.NewInt(1000))
	hash, _ := sdb.Commit(false)
	fmt.Println(hash)

	balance := sdb.GetBalance(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	fmt.Println(balance)

	sdb.Database().TrieDB().Commit(sdb.IntermediateRoot(false), false)
	db.Close()
}
