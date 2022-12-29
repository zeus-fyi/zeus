package signing_automation_ethereum

import (
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ssz"
)

func GenerateEphemeralDepositData(blsSigner bls_signer.Account, withdrawalAddress []byte) (*spec.DepositData, error) {
	fv, err := GetEphemeralForkVersion()
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return GenerateDepositData(blsSigner, withdrawalAddress, fv)
}

func GenerateDepositData(blsSigner bls_signer.Account, withdrawalAddress []byte, forkVersion *spec.Version) (*spec.DepositData, error) {
	var pubKey spec.BLSPubKey
	copy(pubKey[:], blsSigner.PublicKey.Serialize())
	depositMessage := &spec.DepositMessage{
		PublicKey:             pubKey,
		WithdrawalCredentials: withdrawalAddress,
		Amount:                spec.Gwei(ValidatorDeposit32Eth.Uint64()),
	}
	root, err := depositMessage.HashTreeRoot()
	if err != nil {
		log.Err(err)
		return nil, errors.Wrap(err, "failed to generate deposit message root")
	}
	var depositMessageRoot spec.Root
	copy(depositMessageRoot[:], root[:])
	domain, err := generateDepositDomain(forkVersion)
	if err != nil {
		log.Err(err)
		return nil, errors.Wrap(err, "failed to generate deposit domain")
	}
	container := &ssz.Container{
		Root:   root[:],
		Domain: domain[:],
	}
	signingRoot, err := container.HashTreeRoot()
	if err != nil {
		log.Err(err)
		return nil, errors.Wrap(err, "failed to generate hash tree root")
	}
	var blsFormatted spec.BLSSignature
	sig := blsSigner.Sign(signingRoot[:])
	copy(blsFormatted[:], sig.Serialize())
	depositData := &spec.DepositData{
		PublicKey:             pubKey,
		WithdrawalCredentials: withdrawalAddress,
		Amount:                spec.Gwei(ValidatorDeposit32Eth.Uint64()),
		Signature:             blsFormatted,
	}
	return depositData, err
}

func generateDepositDomain(forkVersion *spec.Version) (*spec.Domain, error) {
	domainData := &spec.Domain{}
	res, err := e2types.ComputeDomain(e2types.DomainDeposit, forkVersion[:], e2types.ZeroGenesisValidatorsRoot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate domain value")
	}
	copy(domainData[:], res)
	return domainData, err
}
