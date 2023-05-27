package web3_actions

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

func (w *Web3Actions) TransferERC20TokenManually(ctx context.Context, payload SendContractTxPayload, wait bool, timeoutInSeconds uint64) error {
	w.Dial()
	defer w.C.Close()

	err := w.SetGasPriceAndLimit(ctx, &payload.GasPriceLimits)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3Actions: Transfer: SetGasPriceAndLimit")
		return err
	}
	return w.transferTokenManually(ctx, &payload, wait, timeoutInSeconds)
}

// transferToken requires you to place the amounts in the params, payload amount otherwise is payable
func (w *Web3Actions) transferTokenManually(ctx context.Context, payload *SendContractTxPayload, wait bool, timeoutInSeconds uint64) error {
	payload.ContractFile = ERC20
	payload.MethodName = Transfer
	payload.Amount = &big.Int{}
	err := w.CallContract(ctx, payload, wait, nil, timeoutInSeconds)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Transfer: CallContract")
		return err
	}
	return err
}

func (w *Web3Actions) TransferERC20Token(ctx context.Context, payload SendContractTxPayload) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()

	err := w.SetGasPriceAndLimit(ctx, &payload.GasPriceLimits)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3Actions: Transfer: SetGasPriceAndLimit")
		return nil, err
	}
	return w.transferToken(ctx, &payload)
}

// transferToken requires you to place the amounts in the params, payload amount otherwise is payable
func (w *Web3Actions) transferToken(ctx context.Context, payload *SendContractTxPayload) (*types.Transaction, error) {
	payload.ContractFile = ERC20
	payload.MethodName = Transfer
	payload.Amount = &big.Int{}
	tx, err := w.CallFunctionWithArgs(ctx, payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Transfer: CallContract")
		return tx, err
	}
	return tx, err
}

func (w *Web3Actions) ApproveSpenderERC20Token(ctx context.Context, payload SendContractTxPayload) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()
	err := w.SetGasPriceAndLimit(ctx, &payload.GasPriceLimits)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3Actions: Transfer: SetGasPriceAndLimit")
		return nil, err
	}
	return w.approveToken(ctx, &payload)
}

func (w *Web3Actions) approveToken(ctx context.Context, payload *SendContractTxPayload) (*types.Transaction, error) {
	payload.ContractFile = ERC20
	payload.MethodName = Approve
	payload.Amount = &big.Int{}
	tx, err := w.CallFunctionWithArgs(ctx, payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Approve: CallFunctionWithArgs")
		return tx, err
	}
	return tx, err
}
