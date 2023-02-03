package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

func HandleEthSignRequestBLS(ctx context.Context, secretName, keyName, msg string) (string, error) {
	headerValue := os.Getenv(SessionToken)
	r := resty.New()
	respJSON := make(map[string]any)
	_, err := r.R().
		SetHeader(SecretsHeader, headerValue).
		SetResult(&respJSON).
		Get(fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, secretName))
	if err != nil {
		log.Ctx(ctx).Err(err)
		return "", err
	}
	sk, ok := respJSON[keyName]
	if !ok {
		log.Ctx(ctx).Warn().Interface("key", keyName).Msg("no value found for secret key")
		return "", err
	}
	acc := bls_signer.NewEthSignerBLSFromExistingKey(sk.(string))
	sig := acc.Sign([]byte(msg))
	return bls_signer.ConvertBytesToString(sig.Marshal()), nil
}

func main() {
	lambda.Start(HandleEthSignRequestBLS)
}
