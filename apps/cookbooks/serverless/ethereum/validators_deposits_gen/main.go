package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

func HandleValidatorDepositsGen(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	m := make(map[string]any)

	sr := bls_serverless_signing.EthereumValidatorDepositsGenRequests{}
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
	err = json.Unmarshal(b, &sr)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	headerValue := os.Getenv(SessionToken)
	r := resty.New()
	// gets validator mnemonic to generate deposit data
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, sr.MnemonicAndHDWalletSecretName)
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
	vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
		Mnemonic:             mnemonic.(string),
		Pw:                   hdWalletPassword.(string),
		ValidatorIndexOffset: sr.HdOffset,
		NumValidators:        sr.ValidatorCount,
		Network:              sr.Network,
	}
	var dp []*signing_automation_ethereum.DepositDataParams
	var fv *spec.Version
	var er error
	var wc []byte
	w3Client := signing_automation_ethereum.Web3SignerClient{}
	if sr.Network == "ephemery" {
		if len(sr.WithdrawalAddress) <= 0 {
			dp, er = w3Client.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
				return ApiResponse, er
			}
		} else {
			fv, er = signing_automation_ethereum.GetEphemeralForkVersion(ctx)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
				return ApiResponse, er
			}
			wc, er = json.Marshal(sr.WithdrawalAddress)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
				return ApiResponse, er
			}
			dp, er = w3Client.GenerateDepositDataWithForWdAddr(ctx, vdg, wc, fv)
		}
	} else {
		if len(sr.BeaconURL) > 0 {
			fv, er = signing_automation_ethereum.GetForkVersion(ctx, sr.BeaconURL)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
				return ApiResponse, er
			}
			sr.ForkVersion = fv
		}
		if len(sr.ForkVersion) <= 0 {
			ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
			return ApiResponse, errors.New("fork version is required")
		}
		if len(sr.WithdrawalAddress) >= 0 {
			dp, err = w3Client.GenerateDepositDataWithDefaultWd(ctx, vdg, sr.ForkVersion)
		} else {
			wc, er = json.Marshal(sr.WithdrawalAddress)
			if er != nil {
				log.Ctx(ctx).Err(er)
				ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
				return ApiResponse, er
			}
			dp, err = w3Client.GenerateDepositDataWithForWdAddr(ctx, vdg, wc, sr.ForkVersion)
		}
	}
	b, er = json.Marshal(dp)
	if er != nil {
		log.Ctx(ctx).Err(er)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, er
	}
	ApiResponse = events.APIGatewayProxyResponse{Body: string(b), StatusCode: 200}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleValidatorDepositsGen)
}
