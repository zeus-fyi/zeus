package aws_lambda

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type AwsLambdaTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (t *AwsLambdaTestSuite) TestClientInit() {
	region := "us-west-1"
	a := AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	lm, err := InitLambdaClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(lm)
}

func TestAwsLambdaTestSuite(t *testing.T) {
	suite.Run(t, new(AwsLambdaTestSuite))
}
