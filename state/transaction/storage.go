package transaction

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type StorageTransition struct {
	Label uint8 // 0: state, 1: storage

	Contract common.Address
	Slot     common.Hash // 智能合约的存储槽
	PreValue common.Hash
	NewValue *common.Hash // newValue = nil 则是 SLOAD, 否则为 SSTORE
}

func newStorageTransition(contract common.Address, slot common.Hash, preValue, newValue *common.Hash) *StorageTransition {
	st := &StorageTransition{
		Label: 1,

		Contract: contract,
		Slot:     slot,
		PreValue: *preValue,
	}
	st.NewValue = new(common.Hash)
	if newValue != nil {
		st.NewValue.SetBytes(newValue.Bytes())
	} else {
		st.NewValue = nil
	}
	return st
}

func (s *StorageTransition) GetLabel() uint8 {
	return s.Label
}

func (s *StorageTransition) String() string {
	var str string = ""
	if s.NewValue == nil {
		str = fmt.Sprintf("%d,%v,%v,%v,", s.Label, s.Contract.Hex(), s.Slot.Hex(), s.PreValue.Hex())
	} else {
		str = fmt.Sprintf("%d,%v,%v,%v,%v", s.Label, s.Contract.Hex(), s.Slot.Hex(), s.PreValue.Hex(), s.NewValue.Hex())
	}
	return str
}
