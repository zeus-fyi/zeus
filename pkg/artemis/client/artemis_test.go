package artemis_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type ArtemisClientTestSuite struct {
	test_suites.BaseTestSuite
	ArtemisTestClient ArtemisClient
}

func (t *ArtemisClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	//t.ZeusTestClient = NewDefaultZeusClient(tc.Bearer)
	t.ArtemisTestClient = NewDefaultArtemisClient(tc.Bearer)
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()
}

func TestArtemisClientTestSuite(t *testing.T) {
	suite.Run(t, new(ArtemisClientTestSuite))
}
