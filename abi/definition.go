package abi

const (
	minCache   = 2048
	minHandles = 2048

	//native leveldb用于保存原始交易数据
	//contract leveldb用于保存一行存储全部区块的智能合约交易
	//contract abi leveldb用于保存全部智能合约地址以及其对应的abi
	//block transaction leveldb用于保存处理后的交易数据
	nativeDBPath      = "D:/Project/leveldb/abi/native_leveldb"
	contractDBPath    = "D:/Project/leveldb/abi/contract_leveldb"
	abiDBPath         = "D:/Project/leveldb/abi/contract_abi_leveldb"
	transactionDBPath = "D:/Project/leveldb/abi/block_transaction_leveldb"

	contractLevelDBKey = "All addresses"
)

type (
	//处理后的交易数据的数据格式
	levelDBFormat struct {
		Hash     string
		To       string
		Opcode   string //交易Input Data中的操作码
		Function string //根据操作码解析出的函数名以及参数形式
	}
)
