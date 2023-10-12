package web3_actions

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

// GetSignedTxToCallFunctionWithData prepares the tx for broadcast
func (w *Web3Actions) GetSignedTxToCallFunctionWithData(ctx context.Context, payload *SendContractTxPayload, data []byte) (*types.Transaction, error) {
	var err error
	w.Dial()
	defer w.C.Close()
	if payload == nil {
		return nil, fmt.Errorf("payload is nil")
	}
	scAddr := common.HexToAddress(payload.SmartContractAddr)
	if data != nil {
		payload.Data = data
	}
	err = w.SuggestAndSetGasPriceAndLimitForTx(ctx, payload, scAddr)
	if err != nil {
		log.Warn().Err(err).Msg("Send: SuggestAndSetGasPriceAndLimitForTx")
		log.Ctx(ctx).Err(err).Msg("Send: SuggestAndSetGasPriceAndLimitForTx")
		return nil, err
	}
	var chainID *big.Int
	switch strings.ToLower(w.Network) {
	case "mainnet":
		chainID = new(big.Int).SetInt64(1)
	case "goerli":
		chainID = new(big.Int).SetInt64(5)
	default:
		chainID, err = w.C.ChainID(ctx)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetChainID")
			return nil, fmt.Errorf("couldn't get chain ID: %v", err)
		}
	}
	publicKeyECDSA := w.EcdsaPublicKey()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := w.C.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetPendingTransactionCount")
		return nil, fmt.Errorf("cannot get nonce: %v", err)
	}
	nonceOffset := GetNonceOffset(ctx)
	baseTx := &types.DynamicFeeTx{
		To:        &scAddr,
		Nonce:     nonce + nonceOffset,
		GasFeeCap: payload.GasFeeCap,
		GasTipCap: payload.GasTipCap,
		Gas:       payload.GasLimit,
		Value:     payload.Amount,
		Data:      data,
	}
	tx := types.NewTx(baseTx)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), w.EcdsaPrivateKey())
	if err != nil {
		err = fmt.Errorf("cannot sign transaction: %v", err)
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: SignTx")
		return nil, err
	}
	return signedTx, err
}

// GetSignedTxToCallFunctionWithArgs prepares the tx for broadcast
func (w *Web3Actions) GetSignedTxToCallFunctionWithArgs(ctx context.Context, payload *SendContractTxPayload) (*types.Transaction, error) {
	w.Dial()
	defer w.C.Close()
	err := payload.GenerateBinDataFromParamsAbi(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithArgs: GetDataPayload")
		return nil, err
	}
	signedTx, err := w.GetSignedTxToCallFunctionWithData(ctx, payload, payload.Data)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetSignedTxToCallFunctionWithData")
		return nil, err
	}
	return signedTx, err
}
