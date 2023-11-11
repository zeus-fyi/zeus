package smart_contract_library

import "math/big"

func EtherMultiple(multiple int) *big.Int {
	return new(big.Int).Mul(big.NewInt(int64(multiple)), new(big.Int).SetUint64(1e18))
}
