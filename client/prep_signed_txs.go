package web3_actions

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

func (w *Web3Actions) GetSignedSendTx(ctx context.Context, params SendEtherPayload) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()
	nonce, err := w.GetNonce(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: GetNonce")
		return nil, err
	}
	chainID, err := w.C.ChainID(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: GetChainID")
		return nil, fmt.Errorf("couldn't get chain ID: %v", err)
	}
	scAddr := common.HexToAddress(params.ToAddress.Hex())
	payload := &SendContractTxPayload{
		SmartContractAddr: scAddr.String(),
		SendEtherPayload: SendEtherPayload{
			TransferArgs: TransferArgs{
				ToAddress: params.ToAddress,
				Amount:    params.Amount,
			},
			GasPriceLimits: params.GasPriceLimits,
		},
	}
	err = w.SuggestAndSetGasPriceAndLimitForTx(ctx, payload, common.HexToAddress(params.ToAddress.Hex()))
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: SuggestAndSetGasPriceAndLimitForTx")
		return nil, err
	}
	baseTx := &types.DynamicFeeTx{
		To:        &scAddr,
		Nonce:     nonce,
		GasFeeCap: params.GasFeeCap,
		GasTipCap: params.GasTipCap,
		Gas:       params.GasLimit,
		Value:     params.Amount,
	}
	tx := types.NewTx(baseTx)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), w.EcdsaPrivateKey())
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: SignTx")
		return nil, fmt.Errorf("cannot sign transaction: %v", err)
	}
	return signedTx, err
}
