package aws_secrets

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type AwsSecretManagerTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (t *AwsSecretManagerTestSuite) TestFetchSecrets() {
	region := "us-west-1"
	a := AuthAWS{
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
