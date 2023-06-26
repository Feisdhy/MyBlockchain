package transaction

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// 以下是导出交易的相关字段需求 (注: 结构体里面的数据要给json包访问 -> 需要首字母大写)

type StateTransition struct {
	Label uint8 // 0: state, 1: storage
	// type 类型(1: 转账; 2: 手续费扣除, 只有From字段; 3: 手续费添加给矿工, 只有To字段 ; 4: 合约销毁; 5: 矿工奖励, 只有To字段)
	// 类型 5(矿工奖励) 每个区块只有一个记录
	Type uint8 // (手续费扣除 2!=3 给矿工的手续费)

	From  *Balance
	To    *Balance
	Value *big.Int
}

func newStateTransition(type_ int, from, to *Balance, value *big.Int) *StateTransition {
	return &StateTransition{
		Label: 0,
		Type:  uint8(type_),

		From:  from,
		To:    to,
		Value: new(big.Int).Set(value),
	}
}

func (t *StateTransition) GetLabel() uint8 {
	return t.Label
}

type Balance struct {
	Address common.Address
	Balance *big.Int
}

func newBalance(addr common.Address, before *big.Int) *Balance {
	return &Balance{
		Address: addr,
		Balance: new(big.Int).Set(before),
	}
}

func (t *StateTransition) String() string {
	from, to := "", ""
	if t.From != nil {
		from = t.From.String()
	}
	if t.To != nil {
		to = t.To.String()
	}
	str := fmt.Sprintf("%d,%d,%v,%v,%v", t.Label, t.Type, from, to, t.Value.String())
	return str
}

func (b *Balance) String() string {
	return fmt.Sprintf("%v~%v", b.Address, b.Balance.String())
}
