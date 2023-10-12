package web3_actions

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

func (w *Web3Actions) ReadERC20TokenBalance(ctx context.Context, contractAddress, addrHash string) (*big.Int, error) {
	w.Dial()
	defer w.C.Close()
	payload := SendContractTxPayload{
		SmartContractAddr: contractAddress,
		ContractFile:      ERC20,
		SendEtherPayload:  SendEtherPayload{},
		MethodName:        Decimals,
	}
	payload.MethodName = BalanceOf
	addrString := common.HexToAddress(addrHash).String()
	payload.Params = []interface{}{addrString}
	balance, err := w.GetContractConst(ctx, &payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReadERC20TokenBalance")
		return new(big.Int), err
	}
	return balance[0].(*big.Int), err
}

func (w *Web3Actions) ReadERC20TokenDecimals(ctx context.Context, payload SendContractTxPayload) (int32, error) {
	w.Dial()
	defer w.C.Close()
	payload.Params = []interface{}{}
	decimals, err := w.GetContractConst(ctx, &payload)
	if err != nil {
		return 0, err
	}
	return int32(decimals[0].(uint8)), err
}

func (w *Web3Actions) ReadERC20TokenName(ctx context.Context, contractAddress string) (string, error) {
	w.Dial()
	defer w.C.Close()
	payload := SendContractTxPayload{
		SmartContractAddr: contractAddress,
		ContractFile:      ERC20,
		SendEtherPayload:  SendEtherPayload{},
		MethodName:        "name",
	}
	name, err := w.GetContractConst(ctx, &payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReadERC20TokenName")
		return "", err
	}
	return name[0].(string), err
}

func (w *Web3Actions) ReadERC20Allowance(ctx context.Context, contractAddress, owner, spender string) (*big.Int, error) {
	w.Dial()
	defer w.C.Close()
	payload := SendContractTxPayload{
		SmartContractAddr: contractAddress,
		ContractFile:      ERC20,
		SendEtherPayload:  SendEtherPayload{},
		MethodName:        "allowance",
	}
	ownerStr := common.HexToAddress(owner)
	spenderStr := common.HexToAddress(spender)
	payload.Params = []interface{}{ownerStr, spenderStr}
	balance, err := w.GetContractConst(ctx, &payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReadERC20Allowance")
		return new(big.Int), err
	}
	return balance[0].(*big.Int), err
}
