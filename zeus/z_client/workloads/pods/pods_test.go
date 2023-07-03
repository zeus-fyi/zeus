package pods_client

import (
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

type PodsClientTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient PodsClient
}

func (t *PodsClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient.ZeusClient = zeus_client.NewDefaultZeusClient(tc.Bearer)

	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()

	// generates outputs to /test/outputs dir
	// uses inputs from /test/mocks dir
}

func TestPodsClientTestSuite(t *testing.T) {
	suite.Run(t, new(PodsClientTestSuite))
}
