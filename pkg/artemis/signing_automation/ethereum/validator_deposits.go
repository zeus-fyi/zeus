package signing_automation_ethereum

import (
	"context"

	"github.com/gochain/gochain/v4/common"
	"github.com/gochain/gochain/v4/core/types"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/web3_actions"
)

const EphemeralDepositContractAddr = "0x4242424242424242424242424242424242424242"

type ValidatorDepositParams struct {
	Pubkey                string
	WithdrawalCredentials string
	Signature             string
	DepositDataRoot       string
}

func (w *Web3SignerClient) SignValidatorDeposit(ctx context.Context, depositParams ValidatorDepositParams) (*types.Transaction, error) {
	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: EphemeralDepositContractAddr,
		ContractFile:      web3_actions.ValidatorDeposits,
		MethodName:        web3_actions.Deposit,
		SendEtherPayload: web3_actions.SendEtherPayload{
			TransferArgs: web3_actions.TransferArgs{
				Amount:    ValidatorDeposit32Eth,
				ToAddress: common.Address{},
			},
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},
		Params: []interface{}{depositParams.Pubkey, depositParams.WithdrawalCredentials, depositParams.Signature, depositParams.DepositDataRoot},
	}
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit")
	}

	return signedTx, err
}

/*
reference
   {
     "internalType": "bytes",
     "name": "pubkey",
     "type": "bytes"
   },
   {
     "internalType": "bytes",
     "name": "withdrawal_credentials",
     "type": "bytes"
   },
   {
     "internalType": "bytes",
     "name": "signature",
     "type": "bytes"
   },
   {
     "internalType": "bytes32",
     "name": "deposit_data_root",
     "type": "bytes32"
   }
*/
