package signing_automation_ethereum

import "math/big"

var (
	Gwei   = big.NewInt(1e9)
	Finney = big.NewInt(1e15)
	Ether  = big.NewInt(1e18)

	ThirtyTwo             = big.NewInt(32)
	ValidatorDeposit32Eth = big.NewInt(1).Mul(Ether, ThirtyTwo)
)
