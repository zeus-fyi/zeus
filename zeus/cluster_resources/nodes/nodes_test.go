package nodes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

type NodesConfigTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient      zeus_client.ZeusClient
	ZeusLocalTestClient zeus_client.ZeusClient
}

var ctx = context.Background()

func (t *NodesConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	t.ZeusLocalTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestNodesConfigTestSuite(t *testing.T) {
	suite.Run(t, new(NodesConfigTestSuite))
}
