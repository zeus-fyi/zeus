package serverless_keygen

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	serverless_aws_automation "github.com/zeus-fyi/zeus/builds/serverless/aws_automation"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
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
		MnemonicAndHDWalletSecretName: "testSecretName",
		AgeSecretName:                 "testAgeSecretName",
	}
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", fnUrl, kgReq)
	s.Require().Nil(err)
	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetBody(kgReq).Post("/")

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode())

	region := "us-west-1"
	a := aegis_aws_auth.AuthAWS{
		AccessKey: s.Tc.AccessKeyAWS,
		SecretKey: s.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, a)
	s.Require().Nil(err)
	s.Require().NotNil(sm)

	secretInfo := aegis_aws_secretmanager.SecretInfo{
		Region: region,
		Name:   "testSecretName",
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	newM := make(map[string]any)
	err = json.Unmarshal(b, &newM)
	s.Require().Nil(err)

	secretInfo = aegis_aws_secretmanager.SecretInfo{
		Region: region,
		Name:   "testAgeSecretName",
	}
	b, err = sm.GetSecretBinary(ctx, secretInfo)
	newM = make(map[string]any)
	err = json.Unmarshal(b, &newM)
	s.Require().Nil(err)
}

func TestServerlessKeygenTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessKeygenTestSuite))
}
