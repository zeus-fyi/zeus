package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

func HandleValidatorEncryptedKeystoreZipGenRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	m := make(map[string]any)

	zipReq := bls_serverless_signing.EthereumValidatorEncryptedZipKeysRequests{}
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
	err = json.Unmarshal(b, &zipReq)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	headerValue := os.Getenv(SessionToken)
	r := resty.New()
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, zipReq.AgeSecretName)
	resp, err := r.R().
		SetHeader(SecretsHeader, headerValue).
		Get(url)
	svo := &secretsmanager.GetSecretValueOutput{}
	err = json.Unmarshal(resp.Body(), &svo)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	m = make(map[string]any)
	if svo.SecretString != nil {
		m = make(map[string]any)
		ss := *svo.SecretString
		err = json.Unmarshal([]byte(ss), &m)
		if err != nil {
			log.Ctx(ctx).Err(err)
			ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
			return ApiResponse, err
		}
	} else {
		err = json.Unmarshal(svo.SecretBinary, &m)
		if err != nil {
			log.Ctx(ctx).Err(err)
			ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
			return ApiResponse, err
		}
	}

	var enc age_encryption.Age
	// m should only have one age key in it
	for pubkey, privkey := range m {
		enc = age_encryption.NewAge(privkey.(string), pubkey)
	}

	// gets validator mnemonic to generate deposit data & encrypted keystore zip file
	url = fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, zipReq.MnemonicAndHDWalletSecretName)
	resp, err = r.R().
		SetHeader(SecretsHeader, headerValue).
		Get(url)
	svo = &secretsmanager.GetSecretValueOutput{}
	err = json.Unmarshal(resp.Body(), &svo)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	if svo.SecretString != nil {
		m = make(map[string]any)
		ss := *svo.SecretString
		err = json.Unmarshal([]byte(ss), &m)
		if err != nil {
			log.Ctx(ctx).Err(err)
			ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
			return ApiResponse, err
		}
	} else {
		err = json.Unmarshal(svo.SecretBinary, &m)
		if err != nil {
			log.Ctx(ctx).Err(err)
			ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
			return ApiResponse, err
		}
	}
	mnemonic := m["mnemonic"]
	hdWalletPassword := m["hdWalletPassword"]

	// network doesn"t matter here, just need to generate encrypted key pairs
	vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
		Mnemonic:             mnemonic.(string),
		Pw:                   hdWalletPassword.(string),
		ValidatorIndexOffset: zipReq.HdOffset,
		NumValidators:        zipReq.ValidatorCount,
	}
	inMemFs := memfs.NewMemFs()
	zipFileData, err := vdg.GenerateAgeEncryptedValidatorKeysInMemZipFile(ctx, inMemFs, enc)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 501}
		return ApiResponse, err
	}
	base64EncodedData := base64.StdEncoding.EncodeToString(zipFileData)
	ApiResponse = events.APIGatewayProxyResponse{Body: base64EncodedData, StatusCode: 200, IsBase64Encoded: true}
	ApiResponse.Headers = make(map[string]string)
	ApiResponse.Headers["Content-Type"] = "application/zip"
	ApiResponse.Headers["Content-Disposition"] = "attachment; filename=keystores.zip"
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleValidatorEncryptedKeystoreZipGenRequest)
}
