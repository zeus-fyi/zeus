package aegis_aws_iam

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type AwsIAMTestSuite struct {
	IAMClientAWS
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
	t.IAMClientAWS = iamClient
}

func (t *AwsIAMTestSuite) TestCreateUsers() {
	t.TestClientInit()
	err := t.IAMClientAWS.CreateLambdaUser(ctx, InternalLambdaUserAndPolicy)
	t.Require().Nil(err)
	err = t.IAMClientAWS.CreateLambdaUser(ctx, ExternalLambdaUserAndPolicy)
	t.Require().Nil(err)
}

func (t *AwsIAMTestSuite) TestCreateInternalRolePolicy() {
	t.TestClientInit()
	res, err := t.IAMClientAWS.CreateLambdaRole(ctx, internalLambdaUserTemplateName)
	t.Require().Nil(err)
	t.Assert().NotEmpty(res)
}

func (t *AwsIAMTestSuite) TestCreateInternalPolicy() {
	t.TestClientInit()
	err := t.IAMClientAWS.CreateLambdaUserPolicy(ctx, InternalLambdaUserAndPolicy, "")
	t.Require().Nil(err)
}

func (t *AwsIAMTestSuite) TestCreateExternalPolicy() {
	t.TestClientInit()
	fnName := "testFn"
	err := t.IAMClientAWS.CreateLambdaUserPolicy(ctx, ExternalLambdaUserAndPolicy, fnName)
	t.Require().Nil(err)
}

func TestAwsIAMTestSuite(t *testing.T) {
	suite.Run(t, new(AwsIAMTestSuite))
}
