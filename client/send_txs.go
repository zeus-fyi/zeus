package web3_actions

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

// Send performs a regular native coin transaction (not a contract)
func (w *Web3Actions) Send(ctx context.Context, params SendEtherPayload) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()
	signedTx, err := w.GetSignedSendTx(ctx, params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: GetSignedSendTx")
		return nil, fmt.Errorf("failed to get transaction: %v", err)
	}
	err = w.SendSignedTransaction(ctx, signedTx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: SendTransaction")
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}
	return signedTx, nil
}

// SendSignedTransaction sends the Transaction
func (w *Web3Actions) SendSignedTransaction(ctx context.Context, signedTx *types.Transaction) error {
	w.Dial()
	defer w.C.Close()
	return w.C.SendTransaction(ctx, signedTx)
}

func (w *Web3Actions) SubmitSignedTxAndReturnTxData(ctx context.Context, signedTx *types.Transaction) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()
	err := w.C.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
