package serverless_aws_automation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
	"strings"
)

func AddMnemonicHDWalletSecretInAWSSecretManager(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, mnemonicAndHDWalletSecretName string, hdWalletPassword string, mnemonic string) {
	fmt.Println("INFO: storing mnemonic and wallet password in aws secrets manager with secret name: ", mnemonicAndHDWalletSecretName)
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
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
		if strings.Contains(err.Error(), "already exists") {
			fmt.Println("INFO: secret already exists, skipping creation")
			return
		}
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
		if strings.Contains(err.Error(), "already exists") {
			fmt.Println("INFO: secret already exists, skipping creation")
			return
		}
		panic(err)
	}
}

func AddExternalAccessKeysInAWSSecretManager(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, externalLambdaAccessKeysSecretName string, awsAuthExternal aws_aegis_auth.AuthAWS) {
	fmt.Println("INFO: storing external iam user credentials in aws secrets manager with secret name: ", externalLambdaAccessKeysSecretName)
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
	if awsAuthExternal.AccessKey == "" || awsAuthExternal.SecretKey == "" {
		panic("ERROR: external access key and secret key cannot be empty")
	}
	b, err := json.Marshal(awsAuthExternal)
	if err != nil {
		panic(err)
	}
	si := secretsmanager.CreateSecretInput{
		Name:         aws.String(externalLambdaAccessKeysSecretName),
		SecretBinary: b,
	}
	err = sm.CreateNewSecret(ctx, si)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Println("INFO: secret already exists, skipping creation")
			return
		}
		panic(err)
	}
}

func GetExternalAccessKeySecret(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, sn string) (aws_aegis_auth.AuthAWS, error) {
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
	secretInfo := aegis_aws_secretmanager.SecretInfo{
		Region: awsAuth.Region,
		Name:   sn,
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	if err != nil {
		panic(err)
	}
	extAuth := aws_aegis_auth.AuthAWS{}
	err = json.Unmarshal(b, &extAuth)
	if err != nil {
		panic(err)
	}
	return extAuth, err
}

func UpdateExternalAccessKeySecret(ctx context.Context, auth aws_aegis_auth.AuthAWS, externalLambdaAccessKeysSecretName string, extAuth aws_aegis_auth.AuthAWS) {
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, auth)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(extAuth)
	if err != nil {
		panic(err)
	}
	si := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(externalLambdaAccessKeysSecretName),
		SecretBinary: b,
	}
	_, err = sm.UpdateSecret(ctx, si)
	if err != nil {
		panic(err)
	}
}

func GetSecret(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, sn string) (map[string]string, error) {
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
	secretInfo := aegis_aws_secretmanager.SecretInfo{
		Region: awsAuth.Region,
		Name:   sn,
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	if err != nil {
		panic(err)
	}
	newM := make(map[string]string)
	err = json.Unmarshal(b, &newM)
	if err != nil {
		panic(err)
	}
	return newM, err
}

func GetExternalAccessKeySecretIfExists(ctx context.Context, awsAuth aws_aegis_auth.AuthAWS, sn string) (aws_aegis_auth.AuthAWS, error) {
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, awsAuth)
	if err != nil {
		panic(err)
	}
	secretInfo := aegis_aws_secretmanager.SecretInfo{
		Region: awsAuth.Region,
		Name:   sn,
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	if err != nil {
		if strings.Contains(err.Error(), "can't find the specified secret") {
			fmt.Println("INFO: secret doesn't exists")
			return aws_aegis_auth.AuthAWS{}, nil
		}
		panic(err)
	}
	extAuth := aws_aegis_auth.AuthAWS{}
	err = json.Unmarshal(b, &extAuth)
	if err != nil {
		panic(err)
	}
	return extAuth, err
}
