package aws_lambda

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type AwsLambdaTestSuite struct {
	LambdaClientAWS
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (t *AwsLambdaTestSuite) TestClientInit() {
	a := AuthAWS{
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

func (t *AwsLambdaTestSuite) TestKeystoresLayerCreation() {
	t.TestClientInit()

	lyOut, err := t.LambdaClientAWS.CreateServerlessBLSLambdaFnKeystoreLayer(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lyOut)
}

func TestAwsLambdaTestSuite(t *testing.T) {
	suite.Run(t, new(AwsLambdaTestSuite))
}
