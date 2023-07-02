package system_config_drivers

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type MatrixConfigTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *MatrixConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestMatrixConfigTestSuite(t *testing.T) {
	suite.Run(t, new(MatrixConfigTestSuite))
}
