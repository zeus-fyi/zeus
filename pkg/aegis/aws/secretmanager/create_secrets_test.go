package aegis_aws_secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
)

func (t *AwsSecretManagerTestSuite) TestCreateSecret() {
	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
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
func (t *AwsSecretManagerTestSuite) TestGetSecretBinary() {
	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := InitSecretsManager(ctx, a)
	t.Require().Nil(err)
	t.Require().NotNil(sm)

	secretInfo := SecretInfo{
		Region: region,
		Name:   "mnemonicAndHDWalletGoerli",
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	t.Require().Nil(err)

	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	t.Require().Nil(err)

	fmt.Println(string(b))
}

func (t *AwsSecretManagerTestSuite) TestCreateSecretBinary() {
	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
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
