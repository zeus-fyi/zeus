package web3_actions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
	web3_types "github.com/zeus-fyi/gochain/web3/types"
)

type SendContractTxPayload struct {
	SmartContractAddr string
	SendEtherPayload  // payable would be an amount, otherwise for tokens use the params field
	ContractFile      string
	ContractABI       *abi.ABI // this has first priority, if nil will check default contracts using contract file
	MethodName        string   // name of the smart contract function
	Params            []interface{}
	Data              []byte
}

type SendEtherPayload struct {
	TransferArgs
	GasPriceLimits
}

func (s *SendContractTxPayload) GenerateBinDataFromParamsAbi(ctx context.Context) error {
	myabi := s.ContractABI
	if myabi == nil {
		abiInternal, aerr := web3_types.GetABI(s.ContractFile)
		if aerr != nil {
			log.Err(aerr).Msg("CallContract: GetABI")
			return aerr
		}
		myabi = abiInternal
	}
	fn := myabi.Methods[s.MethodName]
	goParams, err := web3_types.ConvertArguments(fn.Inputs, s.Params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithArgs")
		return err
	}
	data, err := myabi.Pack(s.MethodName, goParams...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithArgs")
		return fmt.Errorf("failed to pack values: %v", err)
	}
	s.Data = data
	return nil
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
	ToAddress accounts.Address
}

type GasPriceLimits struct {
	GasPrice  *big.Int
	GasLimit  uint64
	GasTipCap *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap *big.Int // a.k.a. maxFeePerGas
}

type CallMsg struct {
	From      *accounts.Address // the sender of the 'transaction'
	To        *accounts.Address // the destination contract (nil for contract creation)
	Gas       uint64            // if 0, the call executes with near-infinite gas
	GasPrice  *big.Int          // wei <-> gas exchange ratio
	GasTipCap *big.Int          // a.k.a. maxPriorityFeePerGas
	GasFeeCap *big.Int          // a.k.a. maxFeePerGas
	Value     *big.Int          // amount of wei sent along with the call
	Data      []byte            // input data, usually an ABI-encoded contract method invocation
}
