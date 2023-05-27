package web3_actions

import (
	"math/big"

	"github.com/zeus-fyi/gochain/v4/accounts/abi"
	"github.com/zeus-fyi/gochain/v4/common"
)

type SendContractTxPayload struct {
	SmartContractAddr string
	SendEtherPayload  // payable would be an amount, otherwise for tokens use the params field
	ContractFile      string
	ContractABI       *abi.ABI // this has first priority, if nil will check default contracts using contract file
	MethodName        string   // name of the smart contract function
	Params            []interface{}
}

type SendEtherPayload struct {
	TransferArgs
	GasPriceLimits
}

func (tx *SendEtherPayload) EffectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return dst.Set(tx.GasFeeCap)
	}
	tip := dst.Sub(tx.GasFeeCap, baseFee)
	if tip.Cmp(tx.GasTipCap) > 0 {
		tip.Set(tx.GasTipCap)
	}
	return tip.Add(tip, baseFee)
}

type TransferArgs struct {
	Amount    *big.Int
	ToAddress common.Address
}

type GasPriceLimits struct {
	GasPrice  *big.Int
	GasLimit  uint64
	GasTipCap *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap *big.Int // a.k.a. maxFeePerGas
}

type CallMsg struct {
	From      *common.Address // the sender of the 'transaction'
	To        *common.Address // the destination contract (nil for contract creation)
	Gas       uint64          // if 0, the call executes with near-infinite gas
	GasPrice  *big.Int        // wei <-> gas exchange ratio
	GasTipCap *big.Int        // a.k.a. maxPriorityFeePerGas
	GasFeeCap *big.Int        // a.k.a. maxFeePerGas
	Value     *big.Int        // amount of wei sent along with the call
	Data      []byte          // input data, usually an ABI-encoded contract method invocation
}
