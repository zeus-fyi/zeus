package signing_automation_ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/wealdtech/go-ed25519hd"
	types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type ValidatorDepositGenerationParams struct {
	Fp                                  filepaths.Path
	Mnemonic, Pw                        string
	ValidatorIndexOffset, NumValidators int

	// used for looping derivations
	WithdrawalKeyIndexOffset, NumWithdrawalKeys int
}

/*
EIP-2334 defines derivation path indices for withdrawal and validator keys.
For a given index i the keys will be at the following paths:

withdrawal key: m/12381/3600/i/0
validator key: m/12381/3600/i/0/0
*/

func (vd *ValidatorDepositGenerationParams) GeneratePaddedBytesDefaultDerivedBLSWithdrawalKey(ctx context.Context) ([]byte, error) {
	vd.WithdrawalKeyIndexOffset = 0
	vd.NumWithdrawalKeys = 1

	keySlice, err := vd.GenerateDerivedWithdrawalKeys(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	if len(keySlice) < 0 {
		err = errors.New("no keys were derived")
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	wd, err := ValidateAndReturnBLSPubkeyBytes(keySlice[0])
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	return wd, err
}

func (vd *ValidatorDepositGenerationParams) GenerateDerivedWithdrawalKeys(ctx context.Context) ([]string, error) {
	wdPubkey := make([]string, vd.NumWithdrawalKeys-vd.WithdrawalKeyIndexOffset)
	count := 0
	for i := vd.ValidatorIndexOffset; i < vd.NumWithdrawalKeys-vd.WithdrawalKeyIndexOffset; i++ {
		path := fmt.Sprintf("m/12381/3600/0/%d", i)
		sk, err := vd.DerivedKey(ctx, path)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return []string{}, err
		}
		wdPubkey[count] = bls_signer.ConvertBytesToString(sk.PublicKey().Marshal())
		count++
	}
	return wdPubkey, nil
}

func (vd *ValidatorDepositGenerationParams) GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx context.Context, network string) error {
	initErr := bls_signer.InitEthBLS()
	if initErr != nil {
		log.Ctx(ctx).Err(initErr)
		return initErr
	}
	for i := vd.ValidatorIndexOffset; i < vd.NumValidators; i++ {
		path := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		enc, err := vd.EthDepositEncryptionAndAddMetadata(ctx, path)
		b, err := json.MarshalIndent(enc, "", "\t")
		if err != nil {
			log.Ctx(ctx).Err(err)
			return err
		}
		slashSplit := strings.Split(path, "/")
		underScoreStr := strings.Join(slashSplit, "_")
		vd.Fp.FnOut = fmt.Sprintf("keystore-%s-%s.json", network, underScoreStr)
		err = vd.Fp.WriteToFileOutPath(b)
	}
	return nil
}

func (vd *ValidatorDepositGenerationParams) DerivedKey(ctx context.Context, path string) (*types.BLSPrivateKey, error) {
	seed, serr := ed25519hd.SeedFromMnemonic(vd.Mnemonic, vd.Pw)
	if serr != nil {
		log.Ctx(ctx).Err(serr)
		return nil, serr
	}
	sk, err := util.PrivateKeyFromSeedAndPath(seed, path)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	return sk, err
}

func (vd *ValidatorDepositGenerationParams) EthDepositEncryptionAndAddMetadata(ctx context.Context, path string) (map[string]interface{}, error) {
	sk, err := vd.DerivedKey(ctx, path)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	uuidVal, err := uuid.NewUUID()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	ks := keystorev4.New()
	enc, err := ks.Encrypt(sk.Marshal(), vd.Pw)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return enc, err
	}
	enc["uuid"] = uuidVal.String()
	enc["path"] = path
	enc["pubkey"] = bls_signer.ConvertBytesToString(sk.PublicKey().Marshal())
	enc["version"] = "4"
	return enc, err
}
