package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
aegis_random
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

func HandleEthValidatorKeyGenRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	m := make(map[string]any)
	err := json.Unmarshal([]byte(event.Body), &m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	keyGenRequest := bls_serverless_signing.BlsKeyGenRequests{}
	err = json.Unmarshal(b, &keyGenRequest)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
		return ApiResponse, err
	}

	if keyGenRequest.Region == "" {
		keyGenRequest.Region = "us-west-1"
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(keyGenRequest.Region))
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	sm := aegis_aws_secretmanager.SecretsManagerAuthAWS{
		Client: secretsmanager.NewFromConfig(cfg),
	}
	// if name is provided for this secret, but no mnemonic, then it will assume you want one to be generated and saved
	// will not overwrite any existing secrets to prevent accidental overwrites
	if len(keyGenRequest.MnemonicAndHDWalletSecretName) > 0 {
		if sm.DoesSecretExist(ctx, keyGenRequest.MnemonicAndHDWalletSecretName) {
			log.Info().Msg("INFO: secret already exists, skipping creation")
		} else {
			hdWalletSecrets := make(map[string]string)
			if len(keyGenRequest.Mnemonic) <= 0 {
				mn, er := aegis_random.GenerateMnemonic()
				if er != nil {
					log.Ctx(ctx).Err(er)
					ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 509}
					return ApiResponse, er
				}
				keyGenRequest.Mnemonic = mn
			}
			if len(keyGenRequest.HdWalletPassword) <= 0 {
				pw, er := aegis_random.GenerateRandomPassword(32)
				if er != nil {
					log.Ctx(ctx).Err(er)
					ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 510}
					return ApiResponse, er
				}
				keyGenRequest.HdWalletPassword = pw
			}

			hdWalletSecrets["hdWalletPassword"] = keyGenRequest.HdWalletPassword
			hdWalletSecrets["mnemonic"] = keyGenRequest.Mnemonic
			by, er := json.Marshal(hdWalletSecrets)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 511}
				return ApiResponse, er
			}
			si := secretsmanager.CreateSecretInput{
				Name:         aws.String(keyGenRequest.MnemonicAndHDWalletSecretName),
				SecretBinary: by,
			}
			err = sm.CreateNewSecret(ctx, si)
			if err != nil {
				if strings.Contains(err.Error(), "already exists") {
					fmt.Println("INFO: secret already exists, skipping creation")
				} else {
					log.Ctx(ctx).Err(er)
					ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 512}
					return ApiResponse, er
				}
			}
		}
	}

	// if name is provided for this secret, but no mnemonic, then it will assume you want one to be generated and saved
	// will not overwrite any existing secrets to prevent accidental overwrites
	if len(keyGenRequest.AgeSecretName) > 0 {
		if sm.DoesSecretExist(ctx, keyGenRequest.AgeSecretName) {
			log.Info().Msg("INFO: secret already exists, skipping creation")
		} else {
			ageSecrets := make(map[string]string)
			if len(keyGenRequest.AgePubKey) <= 0 || len(keyGenRequest.AgePrivKey) <= 0 {
				keyGenRequest.AgePubKey, keyGenRequest.AgePrivKey = age_encryption.GenerateNewKeyPair()
			}
			ageSecrets[keyGenRequest.AgePubKey] = keyGenRequest.AgePrivKey
			by, er := json.Marshal(ageSecrets)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 513}
				return ApiResponse, er
			}
			si := secretsmanager.CreateSecretInput{
				Name:         aws.String(keyGenRequest.AgeSecretName),
				SecretBinary: by,
			}
			err = sm.CreateNewSecret(ctx, si)
			if err != nil {
				if strings.Contains(err.Error(), "already exists") {
					fmt.Println("INFO: secret already exists, skipping creation")
				} else {
					log.Ctx(ctx).Err(er)
					ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 514}
					return ApiResponse, er
				}
			}
		}
	}
	ApiResponse = events.APIGatewayProxyResponse{StatusCode: 200}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleEthValidatorKeyGenRequest)
}
