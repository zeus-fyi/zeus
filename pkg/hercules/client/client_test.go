package hercules_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
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
	_ = bls_signer.InitEthBLS()
	t.Tc = configs.InitLocalTestConfigs()
	t.HerculesTestClient = NewLocalHerculesClient(t.Tc.Bearer)
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()
}

func TestHerculesClientTestSuite(t *testing.T) {
	suite.Run(t, new(HerculesClientTestSuite))
}
