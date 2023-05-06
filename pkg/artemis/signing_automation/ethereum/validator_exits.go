package signing_automation_ethereum

import (
	"context"
	"fmt"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ssz"
)

// GenerateVoluntaryExit TODO generateVoluntaryExit needs current fork version
func (w *Web3SignerClient) GenerateVoluntaryExit(ctx context.Context, blsSigner bls_signer.EthBLSAccount, forkVersion *spec.Version) error {
	exitMessage := &spec.VoluntaryExit{
		Epoch:          spec.Epoch(0),
		ValidatorIndex: spec.ValidatorIndex(0),
	}
	root, err := exitMessage.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return errors.Wrap(err, "failed to generate voluntary exit message root")
	}
	var exitMessageRoot spec.Root
	copy(exitMessageRoot[:], root[:])
	domain, err := generateVoluntaryExitDomain(ctx, forkVersion)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return errors.Wrap(err, "failed to generate voluntary exit domain")
	}
	container := &ssz.Container{
		Root:   root[:],
		Domain: domain[:],
	}
	signingRoot, err := container.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return errors.Wrap(err, "failed to generate hash tree root")
	}
	var blsFormatted spec.BLSSignature
	// TODO inject from remote serverless signer
	sig := blsSigner.Sign(signingRoot[:])
	copy(blsFormatted[:], sig.Marshal())
	// TODO get validator index
	voluntaryExitData := &spec.VoluntaryExit{
		Epoch:          0,
		ValidatorIndex: 0,
	}
	ht, err := voluntaryExitData.HashTreeRoot()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return errors.Wrap(err, "failed to generate hash tree root")
	}
	// TODO submit exit to beacon
	fmt.Println("voluntary exit data hash tree root:", ht)
	return err
}

// TODO generateVoluntaryExitDomain needs current fork version
func generateVoluntaryExitDomain(ctx context.Context, forkVersion *spec.Version) (*spec.Domain, error) {
	domainData := &spec.Domain{}
	res, err := e2types.ComputeDomain(e2types.DomainVoluntaryExit, forkVersion[:], e2types.ZeroGenesisValidatorsRoot)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, errors.Wrap(err, "failed to generate domain value")
	}
	copy(domainData[:], res)
	return domainData, err
}
