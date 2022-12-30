package signing_automation_ethereum

import "math/big"

var (
	Gwei   = big.NewInt(1e9)
	Szabo  = big.NewInt(1e12)
	Finney = big.NewInt(1e15)
	Ether  = big.NewInt(1e18)

	OneHundred                       = big.NewInt(100)
	ThirtyTwo                        = big.NewInt(32)
	ValidatorDeposit32Eth            = big.NewInt(1).Mul(Ether, ThirtyTwo)
	ValidatorDeposit32EthInGweiUnits = big.NewInt(1).Mul(Gwei, ThirtyTwo)
)

func MultiplyEtherUnit(mul int64, unit *big.Int) *big.Int {
	multiplier := big.NewInt(mul)
	return big.NewInt(1).Mul(multiplier, unit)
}
