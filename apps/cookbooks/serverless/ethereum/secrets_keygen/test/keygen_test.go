package serverless_keygen

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	serverless_aws_automation "github.com/zeus-fyi/zeus/builds/serverless/aws_automation"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	"github.com/zeus-fyi/zeus/test/configs"

	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ServerlessKeygenTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (s *ServerlessKeygenTestSuite) TestServerlessSigningFunc() {
	s.Tc = configs.InitLocalTestConfigs()
	r := resty.New()
	auth := aegis_aws_auth.AuthAWS{
		Region:    "us-west-1",
		AccessKey: s.Tc.AccessKeyAWS,
		SecretKey: s.Tc.SecretKeyAWS,
	}
	fnUrl, err := serverless_aws_automation.CreateLambdaFunctionSecretsKeyGen(ctx, auth)
	s.Require().Nil(err)
	s.Require().NotEmpty(fnUrl)
	r.SetBaseURL(fnUrl)

	kgReq := bls_serverless_signing.BlsKeyGenRequests{
		MnemonicAndHDWalletSecretName: "testHdWalletPw",
		AgeSecretName:                 "testAgeKey",
	}
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", fnUrl, kgReq)
	s.Require().Nil(err)
	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetBody(kgReq).Post("/")

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode())
}

func TestServerlessKeygenTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessKeygenTestSuite))
}
