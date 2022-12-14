package artemis_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type ArtemisClientTestSuite struct {
	test_suites.BaseTestSuite
	ArtemisTestClient          ArtemisClient
	Web3SignerClientTestClient signing_automation_ethereum.Web3SignerClient

	TestBLSAccount bls_signer.Account
	TestAccount1   ecdsa.Account
	TestAccount2   ecdsa.Account
	NodeURL        string
}

func (t *ArtemisClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.NodeURL = tc.NodeURL
	t.ArtemisTestClient = NewDefaultArtemisClient(tc.Bearer)
	//t.ArtemisTestClient = NewLocalArtemisClient(tc.Bearer)
	// points working dir to inside /test

	pkHexString := tc.LocalEcsdaTestPkey
	t.TestAccount1 = ecdsa.NewAccount(pkHexString)
	t.ArtemisTestClient.Account = t.TestAccount1

	pkHexString2 := tc.LocalEcsdaTestPkey2
	t.TestAccount2 = ecdsa.NewAccount(pkHexString2)

	t.Web3SignerClientTestClient = signing_automation_ethereum.NewWeb3Client(tc.EphemeralNodeURL, t.TestAccount1.Account)
	t.TestBLSAccount = bls_signer.NewSignerBLSFromExistingKey(tc.LocalBLSTestPkey)
}

func TestArtemisClientTestSuite(t *testing.T) {
	suite.Run(t, new(ArtemisClientTestSuite))
}
