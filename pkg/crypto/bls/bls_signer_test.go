package bls_signer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BLSTestSuite struct {
	suite.Suite
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
}

func TestBLSTestSuite(t *testing.T) {
	suite.Run(t, new(BLSTestSuite))
}
