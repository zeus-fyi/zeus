package hercules_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"

	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type HerculesClientTestSuite struct {
	test_suites.BaseTestSuite
	HerculesTestClient HerculesClient
}

func (t *HerculesClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.HerculesTestClient = NewLocalHerculesClient(tc.Bearer)
	//.HerculesTestClient = NewDefaultHerculesClient(tc.Bearer)
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()
}

func TestHerculesClientTestSuite(t *testing.T) {
	suite.Run(t, new(HerculesClientTestSuite))
}
