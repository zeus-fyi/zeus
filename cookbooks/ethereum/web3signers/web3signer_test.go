package web3signer_cookbooks

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type Web3SignerCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient

	CustomWeb3SignerImage string
	AuthURL               string
}

func (t *Web3SignerCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.CustomWeb3SignerImage = tc.Web3SignerDockerImage
	t.AuthURL = tc.Web3SignerAuthURL
	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestWeb3SignerCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(Web3SignerCookbookTestSuite))
}
