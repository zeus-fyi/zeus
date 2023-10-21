package web3_actions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/params"
	"github.com/rs/zerolog/log"
)

func (w *Web3Actions) GetBaseFee(ctx context.Context) (*big.Int, error) {
	w.Dial()
	defer w.C.Close()
	blk, err := w.C.BlockByNumber(ctx, nil)
	if err != nil {
		log.Warn().Err(err).Msg("GetBaseFee: BlockByNumber")
		return nil, err
	}
	config := params.MainnetChainConfig
	switch w.Network {
	case "goerli", "Goerli":
		config = params.GoerliChainConfig
	case "sepolia", "Sepolia":
		config = params.SepoliaChainConfig
	case "ephemery", "Ephemery":
		// todo: add ephemery config
	default:
		config = params.MainnetChainConfig
	}
	baseFee := misc.CalcBaseFee(config, blk.Header())
	return baseFee, nil
}

func (w *Web3Actions) SuggestAndSetGasPriceAndLimitForTx(ctx context.Context, params *SendContractTxPayload, toAddr common.Address) error {
	w.Dial()
	defer w.C.Close()
	/*
		GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
		GasFeeCap  *big.Int // a.k.a. maxFeePerGas
		Gas        uint64 // a.k.a. gasLimit
	*/
	if params.GasLimit != 0 {
		// if user sets gas manually, don't override it
		return nil
	}
	if w.Account == nil {
		log.Warn().Msg("SuggestAndSetGasPriceAndLimitForTx: account is nil")
		return fmt.Errorf("account is nil")
	}
	gasTip, err := w.C.SuggestGasTipCap(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("SuggestAndSetGasPriceAndLimitForTx: SuggestGasTipCap")
		log.Ctx(ctx).Err(err).Msg("SuggestAndSetGasPriceAndLimitForTx: GetGasTip")
		return fmt.Errorf("cannot get gas tip: %v", err)
	}
	params.GasTipCap = gasTip
	baseFee, err := w.GetBaseFee(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("SuggestAndSetGasPriceAndLimitForTx: GetBaseFee")
		log.Ctx(ctx).Err(err).Msg("SuggestAndSetGasPriceAndLimitForTx: GetBaseFee")
		return err
	}

	// Max Fee = (2 * Base Fee) + Max Priority Fee
	gasBaseWithMargin := new(big.Int).Mul(baseFee, big.NewInt(2))
	params.GasFeeCap = new(big.Int).Add(gasBaseWithMargin, gasTip)
	msg := ethereum.CallMsg{
		From:      common.HexToAddress(w.Address().Hex()),
		To:        &toAddr,
		GasFeeCap: params.GasFeeCap,
		GasTipCap: params.GasTipCap,
		Data:      params.Data,
		Value:     params.Amount,
	}
	gasLimit, err := w.C.EstimateGas(ctx, msg)
	if err != nil {
		log.Warn().Err(err).Msg("SuggestAndSetGasPriceAndLimitForTx: EstimateGas")
		log.Ctx(ctx).Err(err).Msg("SuggestAndSetGasPriceAndLimitForTx: EstimateGas")
		return err
	}
	params.GasLimit = gasLimit
	return nil
}
