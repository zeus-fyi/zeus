package signing_automation_ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/wealdtech/go-ed25519hd"
	types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type ValidatorDepositGenerationParams struct {
	Fp                   filepaths.Path
	Mnemonic             string `json:"mnemonic"`
	Pw                   string `json:"hdWalletPassword"`
	ValidatorIndexOffset int    `json:"hdOffset"`
	NumValidators        int    `json:"validatorCount"`

	// used for looping derivations
	WithdrawalKeyIndexOffset int `json:"WdHdOffset"`
	NumWithdrawalKeys        int `json:"numWdKeys"`

	// used for network selection variables
	Network string `json:"network"`
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
	wdPubkey := make([]string, vd.NumWithdrawalKeys)
	count := 0
	for i := vd.WithdrawalKeyIndexOffset; i < vd.NumWithdrawalKeys+vd.WithdrawalKeyIndexOffset; i++ {
		derPath := fmt.Sprintf("m/12381/3600/0/%d", i)
		sk, err := vd.DerivedKey(ctx, derPath)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return []string{}, err
		}
		wdPubkey[count] = bls_signer.ConvertBytesToString(sk.PublicKey().Marshal())
		count++
	}
	return wdPubkey, nil
}

func (vd *ValidatorDepositGenerationParams) GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx context.Context) error {
	initErr := bls_signer.InitEthBLS()
	if initErr != nil {
		log.Ctx(ctx).Err(initErr)
		return initErr
	}
	for i := vd.ValidatorIndexOffset; i < vd.NumValidators+vd.ValidatorIndexOffset; i++ {
		path := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		enc, err := vd.EthDepositEncryptionAndAddMetadata(ctx, path)
		b, err := json.MarshalIndent(enc, "", "\t")
		if err != nil {
			log.Ctx(ctx).Err(err)
			return err
		}
		slashSplit := strings.Split(path, "/")
		underScoreStr := strings.Join(slashSplit, "_")
		vd.Fp.FnOut = fmt.Sprintf("keystore-%s-%s.json", vd.Network, underScoreStr)
		err = vd.Fp.WriteToFileOutPath(b)
	}
	return nil
}

// GenerateAgeEncryptedValidatorKeysInMemZipFile generates a zip file of validator keys encrypted with age, the unzipped contents are keystores/keystores.tar.gz.age
func (vd *ValidatorDepositGenerationParams) GenerateAgeEncryptedValidatorKeysInMemZipFile(ctx context.Context, inMemFs memfs.MemFS, age age_encryption.Age) (*bytes.Buffer, error) {
	initErr := bls_signer.InitEthBLS()
	if initErr != nil {
		log.Ctx(ctx).Err(initErr)
		return nil, initErr
	}
	p := filepaths.Path{DirIn: "./keystores", DirOut: "./gzip", FnOut: "keystores.tar.gz"}
	for i := vd.ValidatorIndexOffset; i < vd.NumValidators+vd.ValidatorIndexOffset; i++ {
		derPath := fmt.Sprintf("m/12381/3600/%d/0/0", i)

		sk, err := vd.DerivedKey(ctx, derPath)
		if err != nil {
			panic(err)
		}
		acc := bls_signer.NewEthSignerBLSFromExistingKey(bls_signer.ConvertBytesToString(sk.Marshal()))
		p.FnIn = strings_filter.AddHexPrefix(acc.PublicKeyString())
		err = inMemFs.MakeFileIn(&p, []byte(bls_signer.ConvertBytesToString(acc.BLSPrivateKey.Marshal())))
		if err != nil {
			return nil, err
		}
	}
	b, err := compression.GzipDirectoryToMemoryFS(p, inMemFs)
	if err != nil {
		return nil, err
	}
	err = inMemFs.MakeFileOut(&p, b)
	if err != nil {
		return nil, err
	}
	err = inMemFs.RemoveAll("./keystores")
	if err != nil {
		return nil, err
	}
	p.FnIn = "keystores.tar.gz"
	p.DirIn = "./gzip"
	p.DirOut = "./gzip"
	err = age.EncryptFromInMemFS(inMemFs, &p)
	if err != nil {
		return nil, err
	}
	zipBytes, err := compression.ZipKeystoreFileInMemory(p, inMemFs)
	if err != nil {
		return nil, err
	}
	return zipBytes, err
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
	m := make(map[string]interface{})
	ks := keystorev4.New()
	enc, err := ks.Encrypt(sk.Marshal(), vd.Pw)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return enc, err
	}
	m["crypto"] = enc
	m["uuid"] = uuidVal.String()
	m["path"] = path
	m["pubkey"] = bls_signer.ConvertBytesToString(sk.PublicKey().Marshal())
	m["version"] = "4"
	return m, err
}
func forceDirToTestSuite() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}
