package signing_automation_ethereum

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
	signing_automation_ethereum_smart_contracts "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum/smart_contracts"
)

const (
	validatorDepositMethodName   = "deposit"
	validatorAbiFileLocation     = "eth_deposit_contract.json"
	EphemeralDepositContractAddr = "0x4242424242424242424242424242424242424242"
	MainnetDepositContractAddr   = "0x00000000219ab540356cBB839Cbe05303d7705Fa"
	GoerliDepositContractAddr    = "0xff50ed3d0ec03aC01D4C79aAd74928BFF48a7b2b"
	EphemeralBeacon              = "https://eth.ephemeral.zeus.fyi"
	BeaconGenesisPath            = "/eth/v1/beacon/genesis"
	BeaconForkPath               = "/eth/v1/beacon/states/head/fork"
	BeaconVoluntaryExitPath      = "/eth/v1/beacon/pool/voluntary_exits"
)

func (w *Web3SignerClient) SignValidatorDepositTxToBroadcastFromJSON(ctx context.Context, depositParams ExtendedDepositParams) (*types.Transaction, error) {
	w.Dial()
	defer w.Close()
	params, err := GetValidatorDepositPayloadV2(ctx, depositParams)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDepositTxToBroadcast")
		return nil, err
	}
	gasPrice, err := w.Web3Actions.C.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	params.GasPrice = gasPrice

	params.GasLimit = 500000
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDepositTxToBroadcast")
		return nil, err
	}
	return signedTx, err
}

func (w *Web3SignerClient) SignValidatorDepositTxToBroadcast(ctx context.Context, depositParams *DepositDataParams) (*types.Transaction, error) {
	w.Dial()
	defer w.Close()
	params, err := GetValidatorDepositPayload(ctx, depositParams)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDepositTxToBroadcast")
		return nil, err
	}

	gasPrice, err := w.Web3Actions.C.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	params.GasPrice = gasPrice
	params.GasLimit = 500000
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDepositTxToBroadcast")
		return nil, err
	}
	return signedTx, err
}

func fromHexStringTo32Byte(str string) ([32]byte, error) {
	// Convert the input string to a byte slice
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return [32]byte{}, err
	}

	// Copy the byte slice to a [32]byte array
	var result [32]byte
	copy(result[:], bytes)

	return result, nil
}

func GetValidatorDepositPayloadV2(ctx context.Context, depositParams ExtendedDepositParams) (web3_actions.SendContractTxPayload, error) {
	abiFile, err := ReadAbi(ctx, strings.NewReader(signing_automation_ethereum_smart_contracts.ValidatorDepositABI))
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit: ReadAbi")
		return web3_actions.SendContractTxPayload{}, err
	}

	pubkey, err := hex.DecodeString(strings.TrimPrefix(depositParams.Pubkey, "0x"))
	if err != nil {
		panic(err)
	}
	sig, err := hex.DecodeString(strings.TrimPrefix(depositParams.Signature, "0x"))
	if err != nil {
		panic(err)
	}

	wd, err := hex.DecodeString(strings.TrimPrefix(depositParams.WithdrawalCredentials, "0x"))
	if err != nil {
		panic(err)
	}

	ddr, err := fromHexStringTo32Byte(depositParams.DepositDataRoot)
	if err != nil {
		panic(err)
	}
	scAddr := MainnetDepositContractAddr
	switch depositParams.NetworkName {
	case "goerli":
		scAddr = GoerliDepositContractAddr
	case "ephemery", "ephemeral":
		scAddr = EphemeralDepositContractAddr
	}
	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: scAddr,
		ContractABI:       abiFile,
		MethodName:        validatorDepositMethodName,
		SendEtherPayload: web3_actions.SendEtherPayload{
			TransferArgs: web3_actions.TransferArgs{
				Amount: ValidatorDeposit32Eth,
			},
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},
		Params: []interface{}{pubkey, wd, sig, ddr},
	}
	return params, err
}

func GetValidatorDepositPayload(ctx context.Context, depositParams *DepositDataParams) (web3_actions.SendContractTxPayload, error) {
	ForceDirToEthSigningDirLocation()
	abiContents := signing_automation_ethereum_smart_contracts.ValidatorDepositABI
	abiFile, err := ABIOpenFile(ctx, abiContents)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit: ABIOpenFile")
		return web3_actions.SendContractTxPayload{}, err
	}

	pubkey, err := hex.DecodeString(strings.TrimPrefix(depositParams.PublicKey.String(), "0x"))
	if err != nil {
		panic(err)
	}
	sig, err := hex.DecodeString(strings.TrimPrefix(depositParams.Signature.String(), "0x"))
	if err != nil {
		panic(err)
	}

	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: EphemeralDepositContractAddr,
		ContractABI:       abiFile,
		MethodName:        validatorDepositMethodName,
		SendEtherPayload: web3_actions.SendEtherPayload{
			TransferArgs: web3_actions.TransferArgs{
				Amount: ValidatorDeposit32Eth,
			},
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},
		Params: []interface{}{pubkey, depositParams.WithdrawalCredentials, sig, depositParams.DepositDataRoot},
	}
	return params, err
}
