package signing_automation_ethereum

import (
	"context"
	"strconv"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ssz"
)

// GenerateVoluntaryExit TODO generateVoluntaryExit needs current fork version
func (w *Web3SignerClient) GenerateVoluntaryExit(ctx context.Context, blsSigner bls_signer.EthBLSAccount, forkVersion *spec.Version, genesisForkVersion [32]byte, validatorIndex string) (*spec.SignedVoluntaryExit, error) {
	vi, err := strconv.Atoi(validatorIndex)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to convert validator index to int")
	}
	exitMessage := &spec.VoluntaryExit{
		Epoch:          spec.Epoch(162304),
		ValidatorIndex: spec.ValidatorIndex(uint64(vi)),
	}
	root, err := exitMessage.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate voluntary exit message root")
	}
	var exitMessageRoot spec.Root
	copy(exitMessageRoot[:], root[:])
	domain, err := generateVoluntaryExitDomain(ctx, forkVersion, genesisForkVersion)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate voluntary exit domain")
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
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate hash tree root")
	}
	signedExit := &spec.SignedVoluntaryExit{
		Message:   exitMessage,
		Signature: blsFormatted,
	}
	return signedExit, err
}

type VoluntaryExitMessage struct {
	Message struct {
		Epoch          string `json:"epoch"`
		ValidatorIndex string `json:"validator_index"`
	} `json:"message"`
	Signature string `json:"signature"`
}

func generateVoluntaryExitDomain(ctx context.Context, forkVersion *spec.Version, genesisForkVersion [32]byte) (*spec.Domain, error) {
	domainData := &spec.Domain{}
	res, err := e2types.ComputeDomain(e2types.DomainVoluntaryExit, forkVersion[:], genesisForkVersion[:])
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate domain value")
	}
	copy(domainData[:], res)
	return domainData, err
}
