package aegis_aws_secretmanager

import (
	"encoding/json"
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

func (t *AwsSecretManagerTestSuite) TestCreateSecretBinary() {
	region := "us-west-1"
	a := AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := InitSecretsManager(ctx, a)
	t.Require().Nil(err)
	t.Require().NotNil(sm)

	m := make(map[string]string)

	m["key1"] = "value1"
	b, err := json.Marshal(m)
	t.Require().Nil(err)

	sn := "exampleSecretBinaryName"
	si := secretsmanager.CreateSecretInput{
		Name:         aws.String(sn),
		SecretBinary: b,
	}
	err = sm.CreateNewSecret(ctx, si)
	t.Require().Nil(err)

	secretInfo := SecretInfo{
		Region: region,
		Name:   sn,
	}
	b, err = sm.GetSecretBinary(ctx, secretInfo)

	newM := make(map[string]string)
	err = json.Unmarshal(b, &newM)
	t.Require().Nil(err)

	t.Assert().Equal(m, newM)
}
