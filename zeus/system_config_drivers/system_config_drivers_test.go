package system_config_drivers

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/client"
)

type SystemConfigTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *SystemConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestSystemConfigTestSuite(t *testing.T) {
	suite.Run(t, new(SystemConfigTestSuite))
}
