package adaptive_rpc_load_balancer_examples

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type AdaptiveRpcLoadBalancerExamplesTestSuite struct {
	test_suites.BaseTestSuite
	Tc configs.TestContainer
}

func (t *AdaptiveRpcLoadBalancerExamplesTestSuite) SetupTest() {
	// points dir to test/configs
	t.Tc = configs.InitLocalTestConfigs()
}

func TestAdaptiveRpcLoadBalancerExamplesTestSuite(t *testing.T) {
	suite.Run(t, new(AdaptiveRpcLoadBalancerExamplesTestSuite))
}
