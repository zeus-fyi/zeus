package iris_proxy_rules_configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type IrisConfigTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *IrisConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestIrisConfigTestSuite(t *testing.T) {
	suite.Run(t, new(IrisConfigTestSuite))
}
