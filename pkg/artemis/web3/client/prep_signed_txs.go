package web3_actions

import (
	"context"
	"fmt"
	"math/big"
	"strings"

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
	var chainID *big.Int
	switch strings.ToLower(w.Network) {
	case "mainnet":
		chainID = new(big.Int).SetInt64(1)
	case "goerli":
		chainID = new(big.Int).SetInt64(5)
	case "hardhat", "local", "anvil":
		chainID = new(big.Int).SetInt64(31337)
	default:
		chainID, err = w.C.ChainID(ctx)
		if err != nil {
			log.Err(err).Msg("CallFunctionWithData: GetChainID")
			return nil, fmt.Errorf("couldn't get chain ID: %v", err)
		}
	}
	scAddr := common.HexToAddress(params.ToAddress.Hex())
	payload := &SendContractTxPayload{
		SmartContractAddr: scAddr.String(),
		SendEtherPayload: SendEtherPayload{
			TransferArgs: TransferArgs{
				ToAddress: params.ToAddress,
				Amount:    params.Amount,
			},
		},
	}
	err = w.SuggestAndSetGasPriceAndLimitForTx(ctx, payload, common.HexToAddress(params.ToAddress.Hex()))
	if err != nil {
		log.Err(err).Msg("Send: SuggestAndSetGasPriceAndLimitForTx")
		return nil, err
	}
	nonceOffset := GetNonceOffset(ctx)
	baseTx := &types.DynamicFeeTx{
		To:        &scAddr,
		Nonce:     nonce + nonceOffset,
		GasFeeCap: params.GasFeeCap,
		GasTipCap: params.GasTipCap,
		Gas:       params.GasLimit,
		Value:     params.Amount,
	}
	tx := types.NewTx(baseTx)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), w.EcdsaPrivateKey())
	if err != nil {
		log.Err(err).Msg("Send: SignTx")
		return nil, fmt.Errorf("cannot sign transaction: %v", err)
	}
	return signedTx, err
}
