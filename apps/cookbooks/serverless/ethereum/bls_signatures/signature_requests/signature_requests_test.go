package bls_serverless_signatures

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
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
		SecretName:        s.Tc.ServerlessSignerFuncSecretName,
		SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
	}
	key := "0x913d41b26a157bc8f539a9f63695b87a066f5086f259673f602a85cf9be0738629e872efd94eda6b08ecfd3c229e875e"
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
