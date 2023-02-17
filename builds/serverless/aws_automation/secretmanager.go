package serverless_aws_automation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
)

func AddMnemonicHDWalletSecretInAWSSecretManager(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, mnemonicAndHDWalletSecretName string, hdWalletPassword string, mnemonic string) {
	fmt.Println("INFO: storing age keypair and hd wallet password in aws secrets manager")
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO: storing mnemonic and wallet password in aws secrets manager with secret name: ", mnemonicAndHDWalletSecretName)
	m := make(map[string]string)
	m["hdWalletPassword"] = hdWalletPassword
	m["mnemonic"] = mnemonic
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	si := secretsmanager.CreateSecretInput{
		Name:         aws.String(mnemonicAndHDWalletSecretName),
		SecretBinary: b,
	}
	err = sm.CreateNewSecret(ctx, si)
	if err != nil {
		panic(err)
	}
}

func AddAgeEncryptionKeyInAWSSecretManager(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, ageEncryptionSecretName, agePubKey, agePrivKey string) {
	fmt.Println("INFO: storing age encryption key in aws secrets manager with secret name: ", ageEncryptionSecretName)
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
	m := make(map[string]string)
	m[agePubKey] = agePrivKey
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	si := secretsmanager.CreateSecretInput{
		Name:         aws.String(ageEncryptionSecretName),
		SecretBinary: b,
	}
	err = sm.CreateNewSecret(ctx, si)
	if err != nil {
		panic(err)
	}
}
