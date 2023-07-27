package serverless_aws_automation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	bls_signer "github.com/zeus-fyi/zeus/pkg/aegis/crypto/bls"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

// VerifyLambdaSigner requires your deposit_data_*.json files to be in your keystoresPath, since it uses it to get your public keys to verify.
func VerifyLambdaSigner(ctx context.Context, auth aegis_aws_auth.AuthAWS, keystoresPath filepaths.Path, funcUrl string, ageEncryptionSecretNameInSecretManager string) error {
	err := bls_signer.InitEthBLS()
	if err != nil {
		return err
	}
	filter := &strings_filter.FilterOpts{StartsWith: "deposit_data", DoesNotInclude: []string{"keystores.tar.gz.age", ".DS_Store", "keystore.zip"}}
	keystoresPath.FilterFiles = filter
	dpSlice, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(ctx, keystoresPath)
	if err != nil {
		return err
	}
	return VerifyLambdaSignerFromDepositDataSlice(ctx, auth, dpSlice, funcUrl, ageEncryptionSecretNameInSecretManager)
}

func VerifyLambdaSignerFromDepositDataSlice(ctx context.Context, auth aegis_aws_auth.AuthAWS, dpSlice signing_automation_ethereum.ValidatorDepositSlice, funcUrl string, ageEncryptionSecretNameInSecretManager string) error {
	err := bls_signer.InitEthBLS()
	if err != nil {
		return err
	}
	sr := bls_serverless_signing.SignatureRequests{
		SecretName:        ageEncryptionSecretNameInSecretManager,
		SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
	}
	fmt.Println(ageEncryptionSecretNameInSecretManager)
	sliceGroup := signing_automation_ethereum.ValidatorDepositSlice{}
	for _, dp := range dpSlice {
		hexMessage, herr := aegis_inmemdbs.RandomHex(10)
		if herr != nil {
			return herr
		}
		sr.SignatureRequests.Map[strings_filter.AddHexPrefix(dp.Pubkey)] = aegis_inmemdbs.EthereumBLSKeySignatureRequest{Message: hexMessage}
		sliceGroup = append(sliceGroup, dp)
		if len(sliceGroup) >= 100 {
			sliceGroup, err = sendVerify(ctx, auth, funcUrl, sr, sliceGroup)
			if err != nil {
				return err
			}
			sr = bls_serverless_signing.SignatureRequests{
				SecretName:        ageEncryptionSecretNameInSecretManager,
				SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
			}
			time.Sleep(1 * time.Second)
		}
	}
	if len(sliceGroup) > 0 {
		sliceGroup, err = sendVerify(ctx, auth, funcUrl, sr, sliceGroup)
	}
	if len(sliceGroup) > 0 {
		return errors.New("some keys were not verified")
	}
	return err
}

// returns keys if the func fails due to non-200 response to retry
func sendVerify(ctx context.Context, auth aegis_aws_auth.AuthAWS, funcUrl string, sr bls_serverless_signing.SignatureRequests,
	dpSlice signing_automation_ethereum.ValidatorDepositSlice) (signing_automation_ethereum.ValidatorDepositSlice, error) {
	r := resty.New()
	r.SetBaseURL(funcUrl)
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", funcUrl, sr)
	if err != nil {
		return signing_automation_ethereum.ValidatorDepositSlice{}, err
	}
	respMsgMap := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	signedEventResponse := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: respMsgMap,
	}
	// the first request make timeout, since it may have a cold start latency
	r.SetTimeout(12 * time.Second)
	r.SetRetryCount(5)
	r.SetRetryWaitTime(500 * time.Millisecond)

	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetResult(&signedEventResponse).
		SetBody(sr).Post("/")

	if err != nil {
		return signing_automation_ethereum.ValidatorDepositSlice{}, err
	}
	respCode := resp.StatusCode()
	if respCode != 200 {
		// cool down, too many requests most likely
		log.Warn().Msgf("resp code: %d", respCode)
		time.Sleep(10 * time.Second)
		return dpSlice, nil
	}

	verified, err := signedEventResponse.VerifySignatures(ctx, sr.SignatureRequests, true)
	if err != nil {
		panic(err)
	}
	for _, key := range verified {
		fmt.Println("verified key: ", key)
	}

	if len(verified) != len(dpSlice) {
		err = errors.New("not all signatures verified")
		return signing_automation_ethereum.ValidatorDepositSlice{}, err
	}
	return signing_automation_ethereum.ValidatorDepositSlice{}, err
}
