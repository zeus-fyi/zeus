package web3_actions

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	web3_types "github.com/zeus-fyi/gochain/web3/types"
)

// GetSignedTxToCallFunctionWithData prepares the tx for broadcast
func (w *Web3Actions) GetSignedTxToCallFunctionWithData(ctx context.Context, payload *SendContractTxPayload, data []byte) (*types.Transaction, error) {
	var err error
	w.Dial()
	defer w.C.Close()
	err = w.SetGasPriceAndLimit(ctx, &payload.GasPriceLimits)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetSignedTxToCallFunctionWithData: SetGasPriceAndLimit")
		return nil, err
	}
	if payload.GasLimit == 21000 {
		payload.GasLimit = 3000000
	}

	chainID, err := w.C.ChainID(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetChainID")
		return nil, fmt.Errorf("couldn't get chain ID: %v", err)
	}
	publicKeyECDSA := w.EcdsaPublicKey()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := w.C.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetPendingTransactionCount")
		return nil, fmt.Errorf("cannot get nonce: %v", err)
	}
	scAddr := common.HexToAddress(payload.SmartContractAddr)
	baseTx := &types.DynamicFeeTx{
		To:        &scAddr,
		Nonce:     nonce,
		GasFeeCap: payload.GasPrice,
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
	myabi := payload.ContractABI
	if myabi == nil {
		abiInternal, aerr := web3_types.GetABI(payload.ContractFile)
		if aerr != nil {
			log.Ctx(ctx).Err(aerr).Msg("CallContract: GetABI")
			return nil, aerr
		}
		myabi = abiInternal
	}
	data, err := myabi.Pack(payload.MethodName, payload.Params...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithArgs")
		return nil, fmt.Errorf("failed to pack values: %v", err)
	}
	signedTx, err := w.GetSignedTxToCallFunctionWithData(ctx, payload, data)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetSignedTxToCallFunctionWithData")
		return nil, err
	}
	return signedTx, err
}
