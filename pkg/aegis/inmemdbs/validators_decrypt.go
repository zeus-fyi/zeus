package aegis_inmemdbs

import (
	"context"
	"encoding/json"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	"os"
)

type DecryptedValidators struct {
	Validators []Validator
	HDPassword string
}

func (dv *DecryptedValidators) ReadValidatorFromKeystore(filepath string) error {
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
