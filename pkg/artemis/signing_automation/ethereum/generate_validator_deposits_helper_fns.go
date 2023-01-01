package signing_automation_ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/wealdtech/go-ed25519hd"
	util "github.com/wealdtech/go-eth2-util"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type ValidatorDepositGenerationParams struct {
	Fp                                  filepaths.Path
	Mnemonic, Pw                        string
	ValidatorIndexOffset, NumValidators int
}

func (vd *ValidatorDepositGenerationParams) GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx context.Context) error {
	seed, serr := ed25519hd.SeedFromMnemonic(vd.Mnemonic, vd.Pw)
	if serr != nil {
		log.Ctx(ctx).Err(serr)
		return serr
	}
	for i := vd.ValidatorIndexOffset; i < vd.NumValidators; i++ {
		path := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		enc, err := EthDepositEncryptionAndAddMetadata(ctx, seed, vd.Pw, path)
		b, err := json.Marshal(enc)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return err
		}
		slashSplit := strings.Split(path, "/")
		underScoreStr := strings.Join(slashSplit, "_")
		vd.Fp.FnOut = fmt.Sprintf("keystore-ephemery-%s.json", underScoreStr)
		err = vd.Fp.WriteToFileOutPath(b)
	}
	return serr
}

func EthDepositEncryptionAndAddMetadata(ctx context.Context, seed []byte, pw, path string) (map[string]interface{}, error) {
	ks := keystorev4.New()
	sk, err := util.PrivateKeyFromSeedAndPath(seed, path)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}

	uuidVal, err := uuid.NewUUID()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	enc, err := ks.Encrypt(sk.Marshal(), pw)
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
