package aegis_aws_secretmanager

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/suite"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	aegis_random "github.com/zeus-fyi/zeus/pkg/crypto/random"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"k8s.io/apimachinery/pkg/util/rand"
)

type AwsSecretManagerTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (t *AwsSecretManagerTestSuite) TestUpdateSecret() {
	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := InitSecretsManager(ctx, a)
	t.Require().Nil(err)

	secretName := "ageEncryptionKey"
	pubkey, privkey := age_encryption.GenerateNewKeyPair()
	m := make(map[string]string)
	m[pubkey] = privkey

	b, err := json.Marshal(m)
	t.Require().Nil(err)
	si := secretsmanager.UpdateSecretInput{
		SecretBinary: b,
		SecretId:     aws.String(secretName),
	}
	err = sm.UpdateSecretAWS(ctx, si)
	t.Require().Nil(err)

	secInfo := SecretInfo{
		Name: secretName,
	}
	sk, err := sm.GetSecretBinary(ctx, secInfo)
	t.Require().Nil(err)
	ms := make(map[string]string)

	err = json.Unmarshal(sk, &ms)
	t.Require().Equal(m, ms)

	secretName = "mnemonicAndHDWallet"
	m = make(map[string]string)
	m["hdWalletPassword"] = rand.String(32)
	mn, err := aegis_random.GenerateMnemonic()
	t.Require().Nil(err)
	m["mnemonic"] = mn
	b, err = json.Marshal(m)
	t.Require().Nil(err)
	si = secretsmanager.UpdateSecretInput{
		SecretBinary: b,
		SecretId:     aws.String(secretName),
	}
}

func (t *AwsSecretManagerTestSuite) TestFetchSecrets() {
	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := InitSecretsManager(ctx, a)
	t.Require().Nil(err)
	t.Require().NotNil(sm)
	secretName := "ethereum/bls/0x8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	key := "0x8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	si := SecretInfo{
		Region: region,
		Name:   secretName,
		Key:    key,
	}
	sk, err := sm.GetSecret(ctx, si)
	t.Require().Nil(err)
	acc := bls_signer.NewEthSignerBLSFromExistingKey(sk)
	expPubKey := "8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	t.Assert().Equal(expPubKey, acc.PublicKeyString())
}

func TestAwsSecretManagerTestSuite(t *testing.T) {
	suite.Run(t, new(AwsSecretManagerTestSuite))
}
