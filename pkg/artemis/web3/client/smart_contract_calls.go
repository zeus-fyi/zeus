package web3_actions

import (
	"context"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

// CallFunctionWithArgs submits a transaction to execute a smart contract function call.
func (w *Web3Actions) CallFunctionWithArgs(ctx context.Context, payload *SendContractTxPayload) (*types.Transaction, error) {
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetSignedTxToCallFunctionWithArgs")
		return nil, err
	}
	return w.SubmitSignedTxAndReturnTxData(ctx, signedTx)
}

// CallFunctionWithData if you already have the encoded function data, then use this
func (w *Web3Actions) CallFunctionWithData(ctx context.Context, payload *SendContractTxPayload, data []byte) (*types.Transaction, error) {
	signedTx, err := w.GetSignedTxToCallFunctionWithData(ctx, payload, data)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: GetSignedTxToCallFunctionWithData")
		return nil, err
	}

	err = w.C.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CallFunctionWithData: SendRawTransaction")
		return nil, fmt.Errorf("cannot send transaction: %v", err)
	}

	return signedTx, nil
}

func convertOutputParams(params []interface{}) []interface{} {
	for ind := range params {
		p := params[ind]
		if h, ok := p.(common.Hash); ok {
			params[ind] = h
		} else if a, okAddr := p.(common.Address); okAddr {
			params[ind] = a
		} else if b, okBytes := p.(hexutil.Bytes); okBytes {
			params[ind] = b
		} else if v := reflect.ValueOf(p); v.Kind() == reflect.Array {
			if t := v.Type(); t.Elem().Kind() == reflect.Uint8 {
				ba := make([]byte, t.Len())
				bv := reflect.ValueOf(ba)
				// Copy since we can't t.Slice() unaddressable arrays.
				for i := 0; i < t.Len(); i++ {
					bv.Index(i).Set(v.Index(i))
				}
				params[ind] = hexutil.Bytes(ba)
			}
		}
	}
	return params
}
