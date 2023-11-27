package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	GasPrice  *big.Int
	GasLimit  uint64
	ToAddress common.Address
	Value     *big.Int
}
