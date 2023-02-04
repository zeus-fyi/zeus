package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
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
	SecretName string `json:"secretName"`
	aegis_inmemdbs.EthereumBLSKeySignatureRequests
}

func HandleEthSignRequestBLS(ctx context.Context, event SignRequestsEvent) (aegis_inmemdbs.EthereumBLSKeySignatureResponses, error) {
	sigResponsesMap := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	batchSigResponses := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: sigResponsesMap,
	}
	headerValue := os.Getenv(SessionToken)
	r := resty.New()
	respJSON := make(map[string]any)
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, event.SecretName)
	_, err := r.R().
		SetHeader(SecretsHeader, headerValue).
		SetResult(&respJSON).
		Get(url)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return batchSigResponses, err
	}

	for pubkey, msg := range event.Map {
		sk, ok := respJSON[pubkey]
		if !ok {
			log.Ctx(ctx).Warn().Interface("key", pubkey).Msg("no value found for secret key")
		} else {
			acc := bls_signer.NewEthSignerBLSFromExistingKey(sk.(string))
			sig := acc.Sign([]byte(msg.Message)).Marshal()
			batchSigResponses.Map[pubkey] = aegis_inmemdbs.EthereumBLSKeySignatureResponse{Signature: bls_signer.ConvertBytesToString(sig)}
		}
	}

	return batchSigResponses, nil
}

func main() {
	lambda.Start(HandleEthSignRequestBLS)
}
