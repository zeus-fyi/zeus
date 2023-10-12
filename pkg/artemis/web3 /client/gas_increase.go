package web3_actions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	web3_types "github.com/zeus-fyi/gochain/web3/types"
)

func (w *Web3Actions) IncreaseGas(ctx context.Context, txHash string, amountGwei string) error {
	w.Dial()
	defer w.C.Close()
	// then we'll clone the original and increase gas
	txOrig, isPending, err := w.C.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		err = fmt.Errorf("error on GetTransactionByHash: %v", err)
		log.Ctx(ctx).Err(err).Msg("IncreaseGas: Dial")
		return err
	}
	if !isPending {
		fmt.Printf("tx isn't pending, so can't increase gas")
		return err
	}
	amount, err := web3_types.ParseGwei(amountGwei)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("IncreaseGas: ParseGwei")
		log.Ctx(ctx).Warn().Msgf("IncreaseGas: failed to parse amount %q: %v\n", amountGwei, err)
		return err
	}
	newPrice := new(big.Int).Add(txOrig.GasPrice(), amount)
	_, err = w.ReplaceTx(ctx, txOrig.ChainId(), txOrig.Nonce(), *txOrig.To(), txOrig.Value(), newPrice, txOrig.Gas(), txOrig.Data())
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("IncreaseGas: ReplaceTx")
		return err
	}
	log.Ctx(ctx).Info().Msgf("IncreaseGas: Increased gas price to %v\n", newPrice)
	return err
}

func (w *Web3Actions) ReplaceTx(ctx context.Context, chainID *big.Int, nonce uint64, to common.Address, amount *big.Int,
	gasPrice *big.Int, gasLimit uint64, data []byte) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()
	if gasPrice == nil {
		gasPriceFetched, err := w.C.SuggestGasPrice(ctx)
		if err != nil {
			err = fmt.Errorf("couldn't get suggested gas price: %v", err)
			log.Ctx(ctx).Err(err).Msg("ReplaceTx: Dial")
			return nil, err
		}
		gasPrice = gasPriceFetched
		fmt.Printf("Using suggested gas price: %v\n", gasPrice)
	}

	if chainID == nil {
		fetchedChainID, err := w.C.ChainID(ctx)
		if err != nil {
			err = fmt.Errorf("couldn't get chain ID: %v", err)
			log.Ctx(ctx).Err(err).Msg("ReplaceTx: Dial")
			return nil, err
		}
		chainID = fetchedChainID
	}
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	fmt.Printf("Replacing transaction nonce: %v, gasPrice: %v, gasLimit: %v\n", nonce, gasPrice, gasLimit)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), w.EcdsaPrivateKey())
	if err != nil {
		err = fmt.Errorf("couldn't sign tx: %v", err)
		log.Ctx(ctx).Err(err).Msg("ReplaceTx: SignTx")
		return nil, err
	}
	err = w.SendSignedTransaction(ctx, signedTx)
	if err != nil {
		err = fmt.Errorf("error sending transaction: %v", err)
		log.Ctx(ctx).Err(err).Msg("ReplaceTx: SendTransaction")
		return nil, err
	}
	log.Ctx(ctx).Info().Msgf("ReplaceTx: Replaced transaction. New transaction:  %s\n", signedTx.Hash().Hex())
	return tx, err
}
