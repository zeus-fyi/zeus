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
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	aws_lambda "github.com/zeus-fyi/zeus/pkg/cloud/aws/lambda"
	"github.com/zeus-fyi/zeus/test/configs"

	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ServerlessDepositsGenTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (s *ServerlessDepositsGenTestSuite) TestDepositsGenFn() {
	s.Tc = configs.InitLocalTestConfigs()
	r := resty.New()
	auth := aegis_aws_auth.AuthAWS{
		Region:    "us-west-1",
		AccessKey: s.Tc.AccessKeyAWS,
		SecretKey: s.Tc.SecretKeyAWS,
	}
	fnUrl, err := serverless_aws_automation.GetLambdaFunctionUrl(ctx, auth, aws_lambda.EthereumCreateValidatorsDepositsFunctionName)
	s.Require().Nil(err)
	s.Require().NotEmpty(fnUrl)
	r.SetBaseURL(fnUrl)

	validatorCount := 3
	dgReq := bls_serverless_signing.EthereumValidatorDepositsGenRequests{
		MnemonicAndHDWalletSecretName: "mnemonicAndHDWalletEphemery",
		ValidatorCount:                validatorCount,
		HdOffset:                      0,
		Network:                       "ephemery",
	}
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", fnUrl, dgReq)
	s.Require().Nil(err)

	depParams := make([]signing_automation_ethereum.DepositDataJSON, validatorCount)
	resp, err := r.R().
		SetResult(&depParams).
		SetHeaderMultiValues(req.Header).
		SetBody(dgReq).Post("/")

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode())
	s.Require().Equal(validatorCount, len(depParams))
	//fmt.Println("response json")
	//respJSON := pretty.Pretty(resp.Body())
	//respJSON = pretty.Color(respJSON, pretty.TerminalStyle)
	//fmt.Println(string(respJSON))
}

func TestServerlessDepositsGenTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessDepositsGenTestSuite))
}
