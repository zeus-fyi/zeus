package aws_secrets

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func (t *AwsSecretManagerTestSuite) TestCreateSecret() {
	region := "us-west-1"
	a := AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := InitSecretsManager(ctx, a)
	t.Require().Nil(err)
	t.Require().NotNil(sm)
	si := secretsmanager.CreateSecretInput{
		Name:         aws.String("exampleSecretName"),
		SecretString: aws.String(t.Tc.AgePrivKey),
	}
	err = sm.CreateNewSecret(ctx, si)
	t.Require().Nil(err)
}
