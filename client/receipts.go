package web3_actions

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

var NotFoundErr = errors.New("not found")

func (w *Web3Actions) GetTxReceipt(ctx context.Context, txhash, contractFile string) error {
	w.Dial()
	defer w.C.Close()
	_, err := w.C.TransactionReceipt(ctx, common.HexToHash(txhash))
	if err != nil {
		err = fmt.Errorf("failed to get transaction receipt: %v", err)
		log.Ctx(ctx).Err(err).Msg("GetTransactionReceipt: GetTransactionReceipt")
		return err
	}
	if verbose {
		fmt.Println("Transaction Receipt Details:")
	}
	return err
}

// WaitForReceipt polls for a transaction receipt until it is available, or ctx is cancelled.
func (w *Web3Actions) WaitForReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	w.Dial()
	defer w.C.Close()
	for {
		receipt, err := w.C.TransactionReceipt(ctx, hash)
		if err == nil {
			return receipt, nil
		}
		if err != NotFoundErr {
			log.Ctx(ctx).Err(err).Msg("WaitForTxReceipt")
			return nil, err
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func FindEventById(abi abi.ABI, id common.Hash) *abi.Event {
	for _, event := range abi.Events {
		if event.ID == id {
			return &event
		}
	}
	return nil
}
