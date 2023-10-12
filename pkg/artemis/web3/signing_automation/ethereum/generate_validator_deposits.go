package signing_automation_ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	bls_signer "github.com/zeus-fyi/zeus/pkg/aegis/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/aegis/crypto/ssz"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type DepositDataParams struct {
	*spec.DepositData  `json:"deposit_data"`
	DepositDataRoot    [32]byte      `json:"deposit_data_root"`
	DepositMessageRoot [32]byte      `json:"deposit_message_root"`
	ForkVersion        *spec.Version `json:"fork_version"`
}

type DepositDataJSON struct {
	Pubkey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Signature             string `json:"signature"`
	Amount                int    `json:"amount"`
	DepositDataRoot       string `json:"deposit_data_root"`
	DepositMessageRoot    string `json:"deposit_message_root"`
	ForkVersion           string `json:"fork_version"`
}

func (dd *DepositDataParams) GetValidatorDepositParamsStringValues() ValidatorDepositParams {
	return ValidatorDepositParams{
		Pubkey:                strings.TrimPrefix(fmt.Sprintf("%#x", dd.PublicKey), "0x"),
		WithdrawalCredentials: strings.TrimPrefix(fmt.Sprintf("%#x", dd.WithdrawalCredentials), "0x"),
		Signature:             strings.TrimPrefix(fmt.Sprintf("%#x", dd.Signature), "0x"),
		DepositDataRoot:       strings.TrimPrefix(fmt.Sprintf("%#x", dd.DepositDataRoot), "0x"),
	}
}

func PrintJSONSlice(p filepaths.Path, dpParamSlice []*DepositDataParams, network string) {
	p.FnOut = fmt.Sprintf("deposit_data-%s-%d.json", network, time.Now().UTC().Unix())

	dpJSON := make([]DepositDataJSON, len(dpParamSlice))

	for i, val := range dpParamSlice {
		dpJSON[i] = val.FormatJSON()
	}
	b, err := json.MarshalIndent(dpJSON, "", "\t")
	if err != nil {
		panic(err)
	}
	err = p.WriteToFileOutPath(b)
	if err != nil {
		panic(err)
	}
}
func DepositDataParamsJSON(dpParamSlice []*DepositDataParams) []DepositDataJSON {
	dpJSON := make([]DepositDataJSON, len(dpParamSlice))

	for i, val := range dpParamSlice {
		dpJSON[i] = val.FormatJSON()
	}
	return dpJSON
}

func (dd *DepositDataParams) FormatJSON() DepositDataJSON {
	if dd.DepositData == nil || dd.ForkVersion == nil {
		panic(errors.New("deposit params are empty"))
	}
	vpStrValues := dd.GetValidatorDepositParamsStringValues()
	pubkey := vpStrValues.Pubkey
	wc := vpStrValues.WithdrawalCredentials
	sig := vpStrValues.Signature
	ddRoot := vpStrValues.DepositDataRoot
	ddMsgRoot := strings.TrimPrefix(fmt.Sprintf("%#x", dd.DepositMessageRoot), "0x")
	fv := strings.TrimPrefix(fmt.Sprintf("%#x", *dd.ForkVersion), "0x")

	djson := DepositDataJSON{
		Pubkey:                pubkey,
		WithdrawalCredentials: wc,
		Signature:             sig,
		Amount:                int(dd.Amount),
		DepositDataRoot:       ddRoot,
		DepositMessageRoot:    ddMsgRoot,
		ForkVersion:           fv,
	}
	return djson
}

func (w *Web3SignerClient) GenerateEphemeryDepositDataWithDefaultWd(ctx context.Context, vdg ValidatorDepositGenerationParams) ([]*DepositDataParams, error) {
	w.Dial()
	defer w.Close()
	fv, err := GetEphemeralForkVersion(ctx)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return w.GenerateDepositDataWithDefaultWd(ctx, vdg, fv)
}

func (w *Web3SignerClient) GenerateDepositDataWithDefaultWd(ctx context.Context, vdg ValidatorDepositGenerationParams, fv *spec.Version) ([]*DepositDataParams, error) {
	depositSlice := make([]*DepositDataParams, vdg.NumValidators)
	initErr := bls_signer.InitEthBLS()
	if initErr != nil {
		log.Ctx(ctx).Err(initErr)
		return depositSlice, initErr
	}
	wc, werr := vdg.GeneratePaddedBytesDefaultDerivedBLSWithdrawalKey(ctx)
	if initErr != nil {
		log.Ctx(ctx).Err(werr)
		return depositSlice, werr
	}
	count := 0
	for i := vdg.ValidatorIndexOffset; i < vdg.NumValidators+vdg.ValidatorIndexOffset; i++ {
		path := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		sk, err := vdg.DerivedKey(ctx, path)
		if err != nil {
			panic(err)
		}
		acc := bls_signer.NewEthSignerBLSFromExistingKey(bls_signer.ConvertBytesToString(sk.Marshal()))
		dd, err := w.GenerateDepositData(ctx, acc, wc, fv)
		if err != nil {
			panic(err)
		}
		depositSlice[count] = dd
		count++
	}
	return depositSlice, nil
}

func (w *Web3SignerClient) GenerateDepositDataWithForWdAddr(ctx context.Context, vdg ValidatorDepositGenerationParams, wd []byte, fv *spec.Version) ([]*DepositDataParams, error) {
	depositSlice := make([]*DepositDataParams, vdg.NumValidators)
	initErr := bls_signer.InitEthBLS()
	if initErr != nil {
		log.Ctx(ctx).Err(initErr)
		return depositSlice, initErr
	}
	count := 0
	for i := vdg.ValidatorIndexOffset; i < vdg.NumValidators+vdg.ValidatorIndexOffset; i++ {
		path := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		sk, err := vdg.DerivedKey(ctx, path)
		if err != nil {
			panic(err)
		}
		acc := bls_signer.NewEthSignerBLSFromExistingKey(bls_signer.ConvertBytesToString(sk.Marshal()))
		dd, err := w.GenerateDepositData(ctx, acc, wd, fv)
		if err != nil {
			panic(err)
		}
		depositSlice[count] = dd
		count++
	}
	return depositSlice, nil
}

func (w *Web3SignerClient) GenerateDepositData(ctx context.Context, blsSigner bls_signer.EthBLSAccount, withdrawalAddress []byte, forkVersion *spec.Version) (*DepositDataParams, error) {
	dp := &DepositDataParams{ForkVersion: forkVersion}
	var pubKey spec.BLSPubKey
	copy(pubKey[:], blsSigner.PublicKey().Marshal())
	depositMessage := &spec.DepositMessage{
		PublicKey:             pubKey,
		WithdrawalCredentials: withdrawalAddress,
		Amount:                spec.Gwei(ValidatorDeposit32EthInGweiUnits.Uint64()),
	}
	root, err := depositMessage.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate deposit message root")
	}
	var depositMessageRoot spec.Root
	copy(depositMessageRoot[:], root[:])
	dp.DepositMessageRoot = depositMessageRoot
	domain, err := generateDepositDomain(ctx, forkVersion)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate deposit domain")
	}
	container := &ssz.Container{
		Root:   root[:],
		Domain: domain[:],
	}
	signingRoot, err := container.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate hash tree root")
	}
	var blsFormatted spec.BLSSignature
	sig := blsSigner.Sign(signingRoot[:])
	copy(blsFormatted[:], sig.Marshal())

	depositData := &spec.DepositData{
		PublicKey:             pubKey,
		WithdrawalCredentials: withdrawalAddress,
		Amount:                spec.Gwei(ValidatorDeposit32EthInGweiUnits.Uint64()),
		Signature:             blsFormatted,
	}
	dp.DepositData = depositData
	ht, err := depositData.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate hash tree root")
	}
	dp.DepositDataRoot = ht
	return dp, err
}

func generateDepositDomain(ctx context.Context, forkVersion *spec.Version) (*spec.Domain, error) {
	domainData := &spec.Domain{}
	res, err := e2types.ComputeDomain(e2types.DomainDeposit, forkVersion[:], e2types.ZeroGenesisValidatorsRoot)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate domain value")
	}
	copy(domainData[:], res)
	return domainData, err
}
