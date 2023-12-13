package ethereum_infra_examples

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

type EthereumInfraExamplesTestSuite struct {
	test_suites.BaseTestSuite
	Tc             configs.TestContainer
	ZeusTestClient zeus_client.ZeusClient
}

func (s *EthereumInfraExamplesTestSuite) SetupTest() {
	s.Tc = configs.InitLocalTestConfigs()
	s.ZeusTestClient = zeus_client.NewDefaultZeusClient(s.Tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestEthereumInfraExamplesTestSuite(t *testing.T) {
	suite.Run(t, new(EthereumInfraExamplesTestSuite))
}
