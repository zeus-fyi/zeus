package bls_signer

import (
	"encoding/hex"
	"fmt"

	"github.com/zeus-fyi/zeus/test/configs"
)

func (s *BLSTestSuite) TestEthBLSKeyGenSignAndVerify() {
	k := NewEthBLSAccount()

	tmp := k.PublicKey()
	fmt.Println(tmp)
	pubkey := k.PublicKeyString()
	fmt.Println(pubkey)

	privKey := k.PrivateKeyString()
	fmt.Println(privKey)
}

func (s *BLSTestSuite) TestEthBLSAccountRestoredFromStringKey() {
	tc := configs.InitLocalTestConfigs()

	a := NewEthSignerBLSFromExistingKey(tc.LocalBLSTestPkey)
	expKey := "83f3ad81d7b364a07e71363f5443e256b8fef62387b56a5fcf33ec0aa3fe20fa54bf4598d86d32da942412c43d3f242c"
	data, err := hex.DecodeString(expKey)
	s.Require().Nil(err)
	s.Assert().Len(data, 48)
	s.Assert().Equal(expKey, a.PublicKeyString())
}
