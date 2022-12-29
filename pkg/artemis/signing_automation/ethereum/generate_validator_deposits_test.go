package signing_automation_ethereum

import bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"

func (t *Web3SignerClientTestSuite) TestEphemeralDepositGenerator() {
	s := bls_signer.NewSignerBLS()
	wd, err := ValidateAndReturnEcdsaPubkeyBytes(t.TestAccount1.PublicKey())
	t.Require().Nil(err)
	dd, err := GenerateEphemeralDepositData(s, wd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(dd)
}
