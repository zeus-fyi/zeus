package web3_actions

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
	"github.com/zeus-fyi/gochain/web3/assets"
)

func (w *Web3Actions) UpgradeContract(ctx context.Context, contractAddress, newTargetAddress string, amount *big.Int, timeoutInSeconds uint64) error {
	w.Dial()
	defer w.C.Close()

	myabi, err := abi.JSON(strings.NewReader(assets.UpgradeableProxyABI))
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("UpgradeContract: Cannot initialize ABI: %v", myabi)
		return err
	}
	gp := GasPriceLimits{
		GasPrice: nil,
		GasLimit: 100000,
	}
	payload := SendContractTxPayload{
		SmartContractAddr: contractAddress,
		MethodName:        Upgrade,
		SendEtherPayload: SendEtherPayload{
			TransferArgs: TransferArgs{
				Amount:    amount,
				ToAddress: accounts.Address{},
			},
			GasPriceLimits: gp,
		},
		Params: []interface{}{newTargetAddress},
	}
	tx, err := w.CallTransactFunction(ctx, &payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("tx", tx).Msg("UpgradeContract: Cannot upgrade the contract")
		return err
	}
	ctx, cancelFn := context.WithTimeout(ctx, time.Duration(timeoutInSeconds)*time.Second)
	defer cancelFn()
	receipt, err := w.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("tx", tx).Msgf("UpgradeContract: Cannot get the receipt for transaction with hash %s", tx.Hash().Hex())
		return err
	}
	log.Ctx(ctx).Info().Msgf("Transaction address: %s", receipt.TxHash.Hex())
	return err
}
