package main

import (
	"context"
	"encoding/json"
	"fmt"
	bls_serverless_signatures "github.com/zeus-fyi/zeus/serverless/ethereum/signatures/signature_requests"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	serverless_inmemfs "github.com/zeus-fyi/zeus/serverless/ethereum/signatures/inmemfs"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

func HandleEthSignRequestBLS(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	m := make(map[string]any)

	sr := bls_serverless_signatures.SignatureRequests{}
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
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, sr.SecretName)
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

	ss := *svo.SecretString
	m = make(map[string]any)
	err = json.Unmarshal([]byte(ss), &m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	// init inmemfs, m should only have one age key in it
	for pubkey, privkey := range m {
		enc := age_encryption.NewAge(privkey.(string), pubkey)
		err = serverless_inmemfs.ImportIntoInMemFs(ctx, enc)
		if err != nil {
			log.Ctx(ctx).Err(err)
			ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
			return ApiResponse, err
		}
	}

	signedResponses, err := serverless_inmemfs.SignValidatorMessagesFromInMemFs(ctx, sr.SignatureRequests)
	b, err = json.Marshal(signedResponses)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	ApiResponse = events.APIGatewayProxyResponse{Body: string(b), StatusCode: 200}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleEthSignRequestBLS)
}
