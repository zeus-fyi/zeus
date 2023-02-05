package aegis_inmemdbs

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"os"
)

type DecryptedValidators struct {
	Validators  []Validator
	HDPassword  string
	DecryptPath string
}

func (dv *DecryptedValidators) ReadValidatorsFromKeystores(filepath string) error {
	if dv.Validators == nil {
		dv.Validators = []Validator{}
	}
	jsonByteArray, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	input := make(map[string]interface{})
	err = json.Unmarshal(jsonByteArray, &input)
	if err != nil {
		panic(err)
	}
	acc, err := signing_automation_ethereum.DecryptKeystoreCipherIntoEthSignerBLS(context.Background(), input, dv.HDPassword)
	if err != nil {
		panic(err)
	}
	v := NewValidator(acc)
	dv.Validators = append(dv.Validators, v)
	return nil
}

// ReadValidatorsFromLightweightKeystores uses pubkey as filename, sk as the file contents
func (dv *DecryptedValidators) ReadValidatorsFromLightweightKeystores(filepath string) error {
	if dv.Validators == nil {
		dv.Validators = []Validator{}
	}
	sk, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	acc := bls_signer.NewEthSignerBLSFromExistingKeyBytes(sk)
	v := NewValidator(acc)
	dv.Validators = append(dv.Validators, v)
	return nil
}

// ReadValidatorFromKeystoreAndGenerateRawKeyfiles provides a lightweight format decrypted output
func (dv *DecryptedValidators) ReadValidatorFromKeystoreAndGenerateRawKeyfiles(filepath string) error {
	if dv.Validators == nil {
		dv.Validators = []Validator{}
	}
	jsonByteArray, err := os.ReadFile(filepath)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	input := make(map[string]interface{})
	err = json.Unmarshal(jsonByteArray, &input)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	acc, err := signing_automation_ethereum.DecryptKeystoreCipherIntoEthSignerBLS(context.Background(), input, dv.HDPassword)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	p := filepaths.Path{
		DirOut: dv.DecryptPath,
		FnOut:  strings_filter.AddHexPrefix(acc.PublicKeyString()),
	}
	err = p.WriteToFileOutPath(acc.BLSPrivateKey.Marshal())
	if err != nil {
		log.Err(err)
		panic(err)
	}
	return nil
}
