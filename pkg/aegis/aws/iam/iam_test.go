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
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	iamClient, err := InitIAMClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(iamClient)
}

func (t *AwsIAMTestSuite) TestCreateInternal() {
	region := "us-west-1"
	a := AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	iamClient, err := InitIAMClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(iamClient)

	resourceID := "3323"
	err = iamClient.CreateLambdaUser(ctx, InternalLambdaUserAndPolicy, resourceID)
	t.Require().Nil(err)
}

func (t *AwsIAMTestSuite) TestCreateExt() {
	region := "us-west-1"
	a := AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	iamClient, err := InitIAMClient(ctx, a)
	t.Require().Nil(err)
	t.Assert().NotEmpty(iamClient)

	resourceID := "3323"
	err = iamClient.CreateLambdaUser(ctx, ExternalLambdaUserAndPolicy, resourceID)
	t.Require().Nil(err)
}

func TestAwsIAMTestSuite(t *testing.T) {
	suite.Run(t, new(AwsIAMTestSuite))
}
