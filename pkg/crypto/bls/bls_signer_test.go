package bls_signer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/configs"
)

type BLSTestSuite struct {
	suite.Suite
}

func (s *BLSTestSuite) TestBlsAccountFromSecretKey() {
	tc := configs.InitLocalTestConfigs()
	a := NewSignerBLSFromExistingKey(tc.LocalBLSTestPkey)
	s.Assert().NotEmpty(a)
	s.Assert().Equal(tc.LocalBLSTestPkey, a.SecretKey.String())

	expPubKey := "0fbc592db51d76b2014ec60169b8a1a3d6040432a85751517e0aaa52da04dd5f3eb99e9a1c3b4dba0f54df73698077801548a1be6abf030bdc37a0b6a3d63c7dd2a229f1db1a9ae9103f3da7532c6631d4b30c0f44df7896eefeb86d40bfbec6"
	s.Assert().Equal(expPubKey, a.PublicKey.String())
}

// https://github.com/supranational/blst/blob/master/bindings/go/blst_minsig_test.go
func (s *BLSTestSuite) TestKeyGenSignAndVerify() {
	k := NewKeyBLS()
	s.Assert().NotEmpty(k)
	s.Assert().NotEmpty(k.PublicKey)
	s.Assert().NotEmpty(k.SecretKey)

	msg := []byte("hello foo")
	sig := k.Sign(msg)
	s.Assert().True(k.Verify(*sig, msg))
	msgUnauthorized := []byte("hello bar")
	s.Assert().False(k.Verify(*sig, msgUnauthorized))

	serializedSecretKey := k.SecretKey.Serialize()
	fmt.Println(len(serializedSecretKey))
	s.Assert().Equal(32, len(serializedSecretKey))

	reGenSecKey := SecretKeyFromBytes(serializedSecretKey)
	s.Assert().Equal(k.SecretKey, reGenSecKey)

	fmt.Println(fmt.Sprintf("%x", serializedSecretKey))
	s.Assert().Equal(fmt.Sprintf("%x", serializedSecretKey), k.SecretKey.String())

	serializedPubKey := k.PublicKey.Serialize()
	reGenPubKey := PublicKeyFromBytes(serializedPubKey)
	s.Assert().Equal(k.PublicKey, reGenPubKey)

	s.Assert().Equal(96, len(serializedPubKey))
	fmt.Println(fmt.Sprintf("%x", serializedPubKey))
	s.Assert().Equal(fmt.Sprintf("%x", serializedPubKey), k.PublicKey.String())
}

func TestBLSTestSuite(t *testing.T) {
	suite.Run(t, new(BLSTestSuite))
}
