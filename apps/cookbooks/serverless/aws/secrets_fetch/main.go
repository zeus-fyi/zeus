package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

type SecretsRequest struct {
	SecretName        string                          `json:"secretName"`
	SignatureRequests EthereumBLSKeySignatureRequests `json:"signatureRequests"`
}

type EthereumBLSKeySignatureRequests struct {
	Map map[string]EthereumBLSKeySignatureRequest `json:"map"`
}

type EthereumBLSKeySignatureRequest struct {
	Message string `json:"message"`
}

func HandleSecretsRequest(ctx context.Context, event SecretsRequest) ([]string, error) {
	headerValue := os.Getenv(SessionToken)
	r := resty.New()
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, event.SecretName)
	resp, err := r.R().
		SetHeader(SecretsHeader, headerValue).
		Get(url)
	svo := &secretsmanager.GetSecretValueOutput{}
	err = json.Unmarshal(resp.Body(), &svo)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}

	ss := *svo.SecretString
	m := make(map[string]any)
	err = json.Unmarshal([]byte(ss), &m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}

	tmp := []string{}
	for k, _ := range m {
		tmp = append(tmp, k)
	}

	return tmp, err
}

func HandleSecretsRequestAPI(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	m := make(map[string]any)

	sr := SecretsRequest{}
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
	}
	err = json.Unmarshal(b, &sr)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
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

	tmp := []string{}
	for k, _ := range m {
		v := sr.SignatureRequests.Map[k]
		tmp = append(tmp, k)
		tmp = append(tmp, v.Message)
	}
	ApiResponse = events.APIGatewayProxyResponse{Body: strings.Join(tmp, ","), StatusCode: 200}
	return ApiResponse, err
}

func main() {
	lambda.Start(HandleSecretsRequestAPI)
}
