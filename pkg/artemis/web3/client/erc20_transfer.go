package web3_actions

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

func (w *Web3Actions) TransferERC20Token(ctx context.Context, payload SendContractTxPayload) (*types.Transaction, error) {
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
