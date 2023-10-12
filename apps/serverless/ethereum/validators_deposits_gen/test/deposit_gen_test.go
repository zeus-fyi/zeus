package serverless_keygen

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/pretty"
	serverless_aws_automation "github.com/zeus-fyi/zeus/builds/serverless/aws_automation"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	zeus_ecdsa "github.com/zeus-fyi/zeus/pkg/aegis/crypto/ecdsa"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/web3/signing_automation/ethereum"
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
	var expectedVersion spec.Version
	copy(expectedVersion[:], []byte{0x00, 0x00, 0x10, 0x20})
	validatorCount := 3
	pkHexString := s.Tc.LocalEcsdaTestPkey
	eth1Account := zeus_ecdsa.NewAccount(pkHexString)
	dgReq := bls_serverless_signing.EthereumValidatorDepositsGenRequests{
		MnemonicAndHDWalletSecretName: "mnemonicAndHDWalletGoerli",
		WithdrawalAddress:             eth1Account.PublicKey(),
		ValidatorCount:                validatorCount,
		HdOffset:                      1,
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
	fmt.Println("response json")
	respJSON := pretty.Pretty(resp.Body())
	respJSON = pretty.Color(respJSON, pretty.TerminalStyle)
	fmt.Println(string(respJSON))
}

func TestServerlessDepositsGenTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessDepositsGenTestSuite))
}
