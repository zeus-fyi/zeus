package bls_serverless_signatures

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	bls_signer "github.com/zeus-fyi/zeus/pkg/aegis/crypto/bls"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ServerlessInMemFSTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx context.Context

func (s *ServerlessInMemFSTestSuite) TestServerlessSigningFunc() {
	r := resty.New()
	r.SetBaseURL(s.Tc.ServerlessSignerFuncBLS)
	respMsgMap := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	signedEventResponse := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: respMsgMap,
	}
	sr := bls_serverless_signing.SignatureRequests{
		SecretName:        "ageEncryptionKeyEphemery",
		SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
	}
	key := "0x8488f29fcbabea243a9d65923cd6e865c450860a8e4e91b08318f564ac40480214debaa6f71d29b05d5443bceca62d01"
	hexMessage, err := aegis_inmemdbs.RandomHex(10)
	s.Require().Nil(err)
	signMsg := aegis_inmemdbs.EthereumBLSKeySignatureRequest{Message: hexMessage}
	sr.SignatureRequests.Map[key] = signMsg
	auth := aegis_aws_auth.AuthAWS{
		AccountNumber: "",
		Region:        "us-west-1",
		AccessKey:     s.Tc.AwsAccessKeyLambdaInvoke,
		SecretKey:     s.Tc.AwsSecretKeyLambdaInvoke,
	}
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", s.Tc.ServerlessSignerFuncBLS, sr)
	s.Require().Nil(err)

	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetResult(&signedEventResponse).
		SetBody(sr).Post("/")
	s.Require().Nil(err)
	respCode := resp.StatusCode()
	s.Require().Equal(200, respCode)
	s.Assert().NotEmpty(respMsgMap)

	err = bls_signer.InitEthBLS()
	s.Require().Nil(err)
	verified, err := signedEventResponse.VerifySignatures(ctx, sr.SignatureRequests, true)
	s.Require().Nil(err)

	s.Assert().Len(verified, 1)
	s.Require().Equal(key, verified[0])
}

func TestServerlessInMemFSTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessInMemFSTestSuite))
}
