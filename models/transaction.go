package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Transaction struct {
	GasPrice  *big.Int
	GasLimit  uint64
	ToAddress common.Address
	Value     *big.Int
}

func (t Transaction) ToEthereum(nonce uint64, contract []byte) *types.Transaction {
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: t.GasPrice,
		Gas:      t.GasLimit,
		To:       &t.ToAddress,
		Value:    t.Value,
		Data:     contract,
	})
}
