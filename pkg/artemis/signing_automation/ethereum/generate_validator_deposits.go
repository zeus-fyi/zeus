package signing_automation_ethereum

import (
	"context"
	"fmt"
	"strings"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ssz"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type DepositDataParams struct {
	*spec.DepositData
	DepositDataRoot    [32]byte
	DepositMessageRoot [32]byte
	ForkVersion        *spec.Version
}

func (dd *DepositDataParams) PrintJSON(p filepaths.Path) {
	if dd.DepositData == nil || dd.ForkVersion == nil {
		panic(errors.New("deposit params are empty"))
	}
	pubkey := strings.TrimPrefix(fmt.Sprintf("%#x", dd.PublicKey), "0x")
	wc := strings.TrimPrefix(fmt.Sprintf("%#x", dd.WithdrawalCredentials), "0x")
	sig := strings.TrimPrefix(fmt.Sprintf("%#x", dd.PublicKey), "0x")
	ddRoot := strings.TrimPrefix(fmt.Sprintf("%#x", dd.DepositDataRoot), "0x")
	ddMsgRoot := strings.TrimPrefix(fmt.Sprintf("%#x", dd.DepositMessageRoot), "0x")
	fv := strings.TrimPrefix(fmt.Sprintf("%#x", *dd.ForkVersion), "0x")

	output := fmt.Sprintf(`{"pubkey":"%s","withdrawal_credentials":"%s","signature":"%s","amount":%d,"deposit_data_root":"%s","deposit_message_root":"%s","fork_version":"%s"}`,
		pubkey,
		wc,
		sig,
		dd.Amount,
		ddRoot,
		ddMsgRoot,
		fv,
	)
	err := p.WriteToFileOutPath([]byte(output))
	if err != nil {
		panic(err)
	}
}

func GenerateEphemeralDepositData(ctx context.Context, blsSigner bls_signer.EthBLSAccount, withdrawalAddress []byte) (*DepositDataParams, error) {
	fv, err := GetEphemeralForkVersion(ctx)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return GenerateDepositData(ctx, blsSigner, withdrawalAddress, fv)
}

func GenerateDepositData(ctx context.Context, blsSigner bls_signer.EthBLSAccount, withdrawalAddress []byte, forkVersion *spec.Version) (*DepositDataParams, error) {
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
