package serverless_keygen

import (
	"context"
	"net/http"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
	serverless_aws_automation "github.com/zeus-fyi/zeus/builds/serverless/aws_automation"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	aws_lambda "github.com/zeus-fyi/zeus/pkg/cloud/aws/lambda"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/test/configs"

	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ServerlessEncKeysZipGenTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (s *ServerlessEncKeysZipGenTestSuite) TestServerlessSigningFunc() {
	s.Tc = configs.InitLocalTestConfigs()
	r := resty.New()
	auth := aegis_aws_auth.AuthAWS{
		Region:    "us-west-1",
		AccessKey: s.Tc.AccessKeyAWS,
		SecretKey: s.Tc.SecretKeyAWS,
	}
	fnUrl, err := serverless_aws_automation.GetLambdaFunctionUrl(ctx, auth, aws_lambda.EthereumValidatorsEncryptedSecretsZipGenFunctionName)
	s.Require().Nil(err)
	s.Require().NotEmpty(fnUrl)
	r.SetBaseURL(fnUrl)

	dgReq := bls_serverless_signing.EthereumValidatorEncryptedZipKeysRequests{
		AgeSecretName:                 "ageEncryptionKeyEphemery",
		MnemonicAndHDWalletSecretName: "mnemonicAndHDWalletEphemery",
		ValidatorCount:                10,
		HdOffset:                      0,
	}
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", fnUrl, dgReq)
	s.Require().Nil(err)
	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetBody(dgReq).Post("/")

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode())
	s.Require().NotZero(len(resp.Body()))
	forceDirToTestSuite()
	fp := filepaths.Path{DirOut: ".", FnOut: "keystores.zip"}
	err = fp.WriteToFileOutPath(resp.Body())
	s.Require().NoError(err)
}

func TestServerlessKeygenTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessEncKeysZipGenTestSuite))
}

func forceDirToTestSuite() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}
