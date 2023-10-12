package signing_automation_ethereum

import (
	"testing"

	"github.com/stretchr/testify/suite"
	bls_signer "github.com/zeus-fyi/zeus/pkg/aegis/crypto/bls"
	zeus_ecdsa "github.com/zeus-fyi/zeus/pkg/aegis/crypto/ecdsa"
	artemis_client "github.com/zeus-fyi/zeus/pkg/artemis/client"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type Web3SignerClientTestSuite struct {
	test_suites.BaseTestSuite
	Web3SignerClientTestClient Web3SignerClient
	ArtemisTestClient          artemis_client.ArtemisClient

	TestBLSAccount bls_signer.EthBLSAccount

	TestHDWalletPassword string
	TestMnemonic         string

	TestAccount1 zeus_ecdsa.Account
	TestAccount2 zeus_ecdsa.Account
	NodeURL      string
}

func (t *Web3SignerClientTestSuite) SetupTest() {
	tc := configs.InitLocalTestConfigs()
	err := bls_signer.InitEthBLS()
	t.Require().Nil(err)

	t.NodeURL = tc.EphemeralNodeURL
	t.ArtemisTestClient = artemis_client.NewDefaultArtemisClient(tc.Bearer)

	pkHexString := tc.LocalEcsdaTestPkey
	t.TestAccount1 = zeus_ecdsa.NewAccount(pkHexString)
	pkHexString2 := tc.LocalEcsdaTestPkey2
	t.TestAccount2 = zeus_ecdsa.NewAccount(pkHexString2)
	t.Web3SignerClientTestClient = NewWeb3Client(t.NodeURL, t.TestAccount1.Account)
	t.TestBLSAccount = bls_signer.NewEthSignerBLSFromExistingKey(tc.LocalBLSTestPkey)
	t.TestMnemonic = tc.LocalMnemonic24Words
	t.TestHDWalletPassword = tc.HDWalletPassword
}

func TestWeb3SignerClientTestSuite(t *testing.T) {
	suite.Run(t, new(Web3SignerClientTestSuite))
}
