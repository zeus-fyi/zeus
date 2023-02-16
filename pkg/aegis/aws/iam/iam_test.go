package aegis_aws_iam

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type AwsIAMTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (t *AwsIAMTestSuite) TestClientInit() {
	region := "us-west-1"
	a := AuthAWS{
		AccountNumber: t.Tc.AwsAccountNumber,
		AccessKey:     t.Tc.AccessKeyAWS,
		SecretKey:     t.Tc.SecretKeyAWS,
		Region:        region,
	}
	iamClient, err := InitIAMClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(iamClient)
}

func (t *AwsIAMTestSuite) TestCreateInternal() {
	region := "us-west-1"
	a := AuthAWS{
		AccountNumber: t.Tc.AwsAccountNumber,
		AccessKey:     t.Tc.AccessKeyAWS,
		SecretKey:     t.Tc.SecretKeyAWS,
		Region:        region,
	}
	iamClient, err := InitIAMClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(iamClient)

	fnName := "testFn"
	err = iamClient.CreateLambdaUser(ctx, InternalLambdaUserAndPolicy, fnName)
	t.Require().Nil(err)
}

func (t *AwsIAMTestSuite) TestCreateExt() {
	region := "us-west-1"
	a := AuthAWS{
		AccountNumber: t.Tc.AwsAccountNumber,
		AccessKey:     t.Tc.AccessKeyAWS,
		SecretKey:     t.Tc.SecretKeyAWS,
		Region:        region,
	}
	iamClient, err := InitIAMClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(iamClient)

	fnName := "testFn"
	err = iamClient.CreateLambdaUser(ctx, ExternalLambdaUserAndPolicy, fnName)
	t.Require().Nil(err)
}

func TestAwsIAMTestSuite(t *testing.T) {
	suite.Run(t, new(AwsIAMTestSuite))
}
