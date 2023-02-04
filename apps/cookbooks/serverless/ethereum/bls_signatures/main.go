package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fastjson"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"

	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

var (
	secretCache, _ = secretcache.New()
)

type SignRequestsEvent struct {
	SecretName        string                                         `json:"secretName"`
	SignatureRequests aegis_inmemdbs.EthereumBLSKeySignatureRequests `json:"signatureRequests"`
}

func HandleEthSignRequestBLS(ctx context.Context, event events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	ApiResponse := events.APIGatewayProxyResponse{}
	// Switch for identifying the HTTP request
	switch event.HTTPMethod {
	case "GET":
	case "POST":
		//validates json and returns error if not working
		err := fastjson.Validate(event.Body)
		if err != nil {
			body := "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + event.Body
			ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
		} else {
			resp, serr := SignMessages(ctx, event.Body)
			if serr != nil {
				body := "Error: Invalid JSON payload ||| " + fmt.Sprint(serr) + " Body Obtained" + "||||" + event.Body
				ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
				return ApiResponse
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: resp, StatusCode: 200}
		}
	}
	return ApiResponse

}
func SignMessages(ctx context.Context, body string) (string, error) {
	b, err := json.Marshal(body)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return "", err
	}

	sr := SignRequestsEvent{}
	err = json.Unmarshal(b, &sr)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return "", err
	}

	if len(sr.SecretName) <= 0 {
		return "", errors.New("no secret name provided")
	}
	headerValue := os.Getenv(SessionToken)
	r := resty.New()

	var respJSON map[string]any
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, sr.SecretName)
	log.Ctx(ctx).Info().Interface("url", url)

	_, err = r.R().
		SetHeader(SecretsHeader, headerValue).
		SetResult(&respJSON).
		Get(url)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return "", err
	}
	signedResponses := aegis_inmemdbs.EthereumBLSKeySignatureResponses{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)}
	for pubkey, msg := range sr.SignatureRequests.Map {
		sk, ok := respJSON[pubkey]
		if !ok {
			log.Ctx(ctx).Warn().Interface("key", pubkey).Msg("no value found for secret key")
		} else {
			acc := bls_signer.NewEthSignerBLSFromExistingKey(sk.(string))
			sig := acc.Sign([]byte(msg.Message)).Marshal()
			signedResponses.Map[pubkey] = aegis_inmemdbs.EthereumBLSKeySignatureResponse{Signature: bls_signer.ConvertBytesToString(sig)}
		}
	}
	b, serr := json.Marshal(signedResponses)
	return string(b), serr
}

func main() {
	lambda.Start(HandleEthSignRequestBLS)
}
