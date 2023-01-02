package ethereum_cookbook_test_suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	artemis_client "github.com/zeus-fyi/zeus/pkg/artemis/client"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type EthereumCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient             zeus_client.ZeusClient
	ArtemisTestClient          artemis_client.ArtemisClient
	Web3SignerClientTestClient signing_automation_ethereum.Web3SignerClient

	TestAccount1          ecdsa.Account
	NodeURL               string
	CustomWeb3SignerImage string
}

func (t *EthereumCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	t.Tc = configs.InitLocalTestConfigs()
	t.CustomWeb3SignerImage = t.Tc.Web3SignerDockerImage
	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(t.Tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()

	t.NodeURL = t.Tc.EphemeralNodeURL
	pkHexString := t.Tc.LocalEcsdaTestPkey
	t.TestAccount1 = ecdsa.NewAccount(pkHexString)
	t.Web3SignerClientTestClient = signing_automation_ethereum.NewWeb3Client(t.NodeURL, t.TestAccount1.Account)
}

func TestEthereumAutomationCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(EthereumCookbookTestSuite))
}
