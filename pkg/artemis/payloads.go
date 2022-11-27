package artemis

import (
	"math/big"

	"github.com/gochain/gochain/v4/common"
	"github.com/gochain/gochain/v4/core/types"
)

type SignedTxPayload struct {
	*types.Transaction
}

type SendEtherPayload struct {
	Amount    *big.Int
	ToAddress common.Address
	GasPriceLimits
}

type GasPriceLimits struct {
	GasPrice *big.Int
	GasLimit uint64
}
