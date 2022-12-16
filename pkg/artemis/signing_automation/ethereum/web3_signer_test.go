package signing_automation_ethereum

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type Web3SignerClientTestSuite struct {
	test_suites.BaseTestSuite
	Web3SignerClientTestClient Web3SignerClient

	TestAccount1 ecdsa.Account
	TestAccount2 ecdsa.Account
	NodeURL      string
}

func (t *Web3SignerClientTestSuite) SetupTest() {
	tc := configs.InitLocalTestConfigs()
	t.NodeURL = tc.EphemeralNodeURL

	pkHexString := tc.LocalEcsdaTestPkey
	t.TestAccount1 = ecdsa.NewAccount(pkHexString)
	pkHexString2 := tc.LocalEcsdaTestPkey2
	t.TestAccount2 = ecdsa.NewAccount(pkHexString2)
	t.Web3SignerClientTestClient = NewWeb3Client(tc.EphemeralNodeURL, t.TestAccount1.Account)
}

func TestWeb3SignerClientTestSuite(t *testing.T) {
	suite.Run(t, new(Web3SignerClientTestSuite))
}
