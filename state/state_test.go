package state

import (
	"testing"
)

func TestStateDB(t *testing.T) {
	//db := rawdb.NewMemoryDatabase()
	//sdb, _ := state.New(types.EmptyRootHash, state.NewDatabase(db), nil)
	////fmt.Println(sdb.Database().TrieDB().Commit(types.EmptyRootHash, true))
	//
	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F97"))
	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F98"))
	//sdb.GetOrNewStateObject(common.HexToAddress("0xEf8801eaf234ff82801821FFe2d78D60a0237F99"))
	//hash, _ := sdb.Commit(false)
	//sdb.Database().TrieDB().Commit(hash, true)
	////sdb.Commit(false)
	//
	//file, _ := leveldb.New(testDbPath, minCache, minHandles, "diskDB", false)
	//defer file.Close()
}
