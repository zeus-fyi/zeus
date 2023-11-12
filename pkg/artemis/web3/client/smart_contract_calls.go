package web3_actions

import (
	"context"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

const (
	Owner = "owner"
)

func (w *Web3Actions) SetBalanceAtSlotNumber(ctx context.Context, scAddr, userAddr string, slotNum int, value *big.Int) error {
	w.Dial()
	defer w.Close()
	slotHex, err := GetSlot(userAddr, new(big.Int).SetUint64(uint64(slotNum)))
	if err != nil {
		return err
	}
	newBalance := common.LeftPadBytes(value.Bytes(), 32)
	err = w.HardhatSetStorageAt(ctx, scAddr, slotHex, common.BytesToHash(newBalance).Hex())
	if err != nil {
		return err
	}
	return nil
}

func (w *Web3Actions) HardhatSetStorageAt(ctx context.Context, addr, slot, value string) error {
	w.Dial()
	defer w.Close()
	err := w.SetStorageAt(ctx, addr, slot, value)
	if err != nil {
		return err
	}
	return err
}

func GetSlot(userAddress string, slot *big.Int) (string, error) {
	// compute keccak256 hash
	addr := common.HexToAddress(userAddress)
	hash := crypto.Keccak256Hash(
		common.LeftPadBytes(addr.Bytes(), 32),
		common.LeftPadBytes(slot.Bytes(), 32),
	)
	// return hex string of the hash
	return hash.Hex(), nil
}

func (w *Web3Actions) GetOwner(ctx context.Context, abiFile *abi.ABI, contractAddress string) (common.Address, error) {
	w.Dial()
	defer w.C.Close()
	payload := SendContractTxPayload{
		SmartContractAddr: contractAddress,
		ContractABI:       abiFile,
		SendEtherPayload:  SendEtherPayload{},
		MethodName:        Owner,
	}
	payload.Params = []interface{}{}
	owner, err := w.GetContractConst(ctx, &payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetOwner")
		return common.Address{}, err
	}
	return owner[0].(common.Address), err
}

// CallFunctionWithArgs submits a transaction to execute a smart contract function call.
func (w *Web3Actions) CallFunctionWithArgs(ctx context.Context, payload *SendContractTxPayload) (*types.Transaction, error) {
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, payload)
	if err != nil {
		log.Err(err).Msg("CallFunctionWithData: GetSignedTxToCallFunctionWithArgs")
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
		log.Err(err).Msg("CallFunctionWithData: SendRawTransaction")
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
