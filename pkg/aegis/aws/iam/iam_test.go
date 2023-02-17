package aegis_aws_iam

import (
	"context"
	"github.com/stretchr/testify/suite"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
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
	a := aegis_aws_auth.AuthAWS{
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

func (t *AwsIAMTestSuite) TestCreateInternalRole() {
	t.TestClientInit()
	res, err := t.IAMClientAWS.CreateInternalLambdaRole(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(res)
}

func (t *AwsIAMTestSuite) TestCreateInternalPolicy() {
	t.TestClientInit()
	r, err := t.IAMClientAWS.CreateNewLambdaUserPolicy(ctx, InternalLambdaUserAndPolicy)
	t.Require().Nil(err)
	t.Assert().NotEmpty(r)
}

func (t *AwsIAMTestSuite) TestAddInternalLambdaRolePolicies() {
	t.TestClientInit()
	res, err := t.IAMClientAWS.AddInternalPolicyToLambdaRolePolicies(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(res)
}

func (t *AwsIAMTestSuite) TestCreateExternalPolicy() {
	t.TestClientInit()
	r, err := t.IAMClientAWS.CreateNewLambdaUserPolicy(ctx, ExternalLambdaUserAndPolicy)
	t.Require().Nil(err)
	t.Assert().NotEmpty(r)
}

func TestAwsIAMTestSuite(t *testing.T) {
	suite.Run(t, new(AwsIAMTestSuite))
}
