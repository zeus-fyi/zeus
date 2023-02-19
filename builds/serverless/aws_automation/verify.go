package serverless_aws_automation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/zeus-fyi/zeus/builds"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func VerifyLambdaSigner(ctx context.Context, auth aegis_aws_auth.AuthAWS, keystoresPath filepaths.Path, funcUrl string, ageEncryptionSecretNameInSecretManager string) {
	r := resty.New()
	r.SetBaseURL(funcUrl)
	respMsgMap := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	signedEventResponse := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: respMsgMap,
	}
	sr := bls_serverless_signing.SignatureRequests{
		SecretName:        ageEncryptionSecretNameInSecretManager,
		SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
	}

	fmt.Println(ageEncryptionSecretNameInSecretManager)
	builds.ChangeToBuildsDir()
	filter := &strings_filter.FilterOpts{StartsWith: "deposit_data", DoesNotInclude: []string{"keystores.tar.gz.age", ".DS_Store"}}
	keystoresPath.FilterFiles = filter
	dpSlice, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(ctx, keystoresPath)
	if err != nil {
		panic(err)
	}
	for _, dp := range dpSlice {
		hexMessage, herr := aegis_inmemdbs.RandomHex(10)
		if herr != nil {
			panic(herr)
		}
		sr.SignatureRequests.Map[strings_filter.AddHexPrefix(dp.Pubkey)] = aegis_inmemdbs.EthereumBLSKeySignatureRequest{Message: hexMessage}
	}

	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", funcUrl, sr)
	if err != nil {
		panic(err)
	}

	// the first request make timeout, since it may have a cold start latency
	r.SetTimeout(3 * time.Second)
	r.SetRetryCount(5)
	r.SetRetryWaitTime(500 * time.Millisecond)

	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetResult(&signedEventResponse).
		SetBody(sr).Post("/")

	if err != nil {
		panic(err)
	}
	respCode := resp.StatusCode()
	if respCode != 200 {
		err = errors.New("status code not 200")
		panic(err)
	}

	err = bls_signer.InitEthBLS()
	verified, err := signedEventResponse.VerifySignatures(ctx, sr.SignatureRequests, true)
	if err != nil {
		panic(err)
	}
	for _, key := range verified {
		fmt.Println("verified key: ", key)
	}

	if len(verified) != len(dpSlice) {
		err = errors.New("not all signatures verified")
		panic(err)
	}
}
