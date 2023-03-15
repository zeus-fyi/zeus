package aws_lambda

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type AwsLambdaTestSuite struct {
	LambdaClientAWS
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (t *AwsLambdaTestSuite) TestClientInit() {
	a := aws_aegis_auth.AuthAWS{
		AccountNumber: t.Tc.AwsAccountNumber,
		AccessKey:     t.Tc.AccessKeyAWS,
		SecretKey:     t.Tc.SecretKeyAWS,
		Region:        region,
	}
	lm, err := InitLambdaClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lm)
	t.LambdaClientAWS = lm
}

func (t *AwsLambdaTestSuite) TestLambdaFnCreation() {
	t.TestClientInit()
	lf, err := t.LambdaClientAWS.CreateServerlessBLSLambdaFn(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lf)
}

func (t *AwsLambdaTestSuite) TestUpdateSignerLambdaFnBinary() {
	t.TestClientInit()
	lf, err := t.LambdaClientAWS.UpdateServerlessBLSLambdaFnBinary(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lf)
}

func (t *AwsLambdaTestSuite) TestKeystoresLayerCreation() {
	t.TestClientInit()
	keystoresZipBin := make([]byte, 0)
	lyOut, err := t.LambdaClientAWS.CreateServerlessBLSLambdaFnKeystoreLayer(ctx, "blsKeystores", keystoresZipBin)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lyOut)
}

func (t *AwsLambdaTestSuite) TestMakeLambdaFnURL() {
	t.TestClientInit()
	lf, err := t.LambdaClientAWS.MakeEthereumSignerURL(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lf)
	fmt.Print(lf.FunctionUrl)
}

func (t *AwsLambdaTestSuite) TestMakeLambdaURLPublic() {
	t.TestClientInit()
	lf, err := t.LambdaClientAWS.MakeEthereumSignerFuncPublic(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lf)
}

func (t *AwsLambdaTestSuite) TestUpdateLambdaLayer() {
	t.TestClientInit()
	lf, err := t.LambdaClientAWS.UpdateServerlessBLSLambdaFnKeystoreLayer(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lf)
}

func (t *AwsLambdaTestSuite) TestGetExtLambdaFnInfo() {
	t.TestClientInit()
	lf, err := t.LambdaClientAWS.GetExternalLambdaFuncInfo(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lf)

	lurl, err := t.LambdaClientAWS.GetExternalLambdaSignerConfigURL(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lurl)
}

func TestAwsLambdaTestSuite(t *testing.T) {
	suite.Run(t, new(AwsLambdaTestSuite))
}
