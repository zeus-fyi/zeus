package artemis_req_types

import (
	"math/big"

	"github.com/zeus-fyi/gochain/v4/accounts/abi"
	"github.com/zeus-fyi/gochain/v4/common"
	"github.com/zeus-fyi/gochain/v4/core/types"
)

type SignedTxPayload struct {
	types.Transaction `json:"tx"`
}

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

type TransferArgs struct {
	Amount    *big.Int
	ToAddress common.Address
}

type GasPriceLimits struct {
	GasPrice *big.Int
	GasLimit uint64
}
