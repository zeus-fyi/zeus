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

func (w *Web3Actions) SuggestAndSetGasPriceAndLimitForTx(ctx context.Context, params *SendContractTxPayload, toAddr common.Address, data []byte) error {
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
	gasTip, err := w.C.SuggestGasTipCap(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: GetGasTip")
		return fmt.Errorf("cannot get gas tip: %v", err)
	}
	params.GasTipCap = gasTip
	baseFee, err := w.GetBaseFee(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: GetBaseFee")
		return err
	}
	// Max Fee = (2 * Base Fee) + Max Priority Fee
	gasBaseWithMargin := new(big.Int).Mul(baseFee, big.NewInt(2))
	params.GasFeeCap = new(big.Int).Add(gasBaseWithMargin, gasTip)
	gasPrice, err := w.C.SuggestGasPrice(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: EstimateGas")
		return err
	}
	msg := ethereum.CallMsg{
		From:      common.HexToAddress(w.Address().Hex()),
		To:        &toAddr,
		GasPrice:  gasPrice,
		GasFeeCap: params.GasFeeCap,
		GasTipCap: params.GasTipCap,
		Data:      data,
		Value:     params.Amount,
	}
	gasLimit, err := w.C.EstimateGas(ctx, msg)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Send: EstimateGas")
		return err
	}
	params.GasLimit = gasLimit
	return nil
}
