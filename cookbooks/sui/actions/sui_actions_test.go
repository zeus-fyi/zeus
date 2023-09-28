package sui_actions

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	pods_client "github.com/zeus-fyi/zeus/zeus/z_client/workloads/pods"
)

var ctx = context.Background()

type SuiActionsCookbookTestSuite struct {
	test_suites.BaseTestSuite
	su SuiActionsClient
}

func (t *SuiActionsCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	zc := zeus_client.NewDefaultZeusClient(tc.Bearer)
	cookbooks.ChangeToCookbookDir()
	t.su = InitSuiClient(pods_client.NewPodsClientFromZeusClient(zc))
}

func TestSuiActionsCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(SuiActionsCookbookTestSuite))
}
