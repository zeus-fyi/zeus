package adaptive_rpc_load_balancer_examples

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	web3_actions "github.com/zeus-fyi/zeus/pkg/artemis/web3/client"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

const (
	LoadBalancerAddress = "https://iris.zeus.fyi/v1/router"
)

var ctx = context.Background()

type AdaptiveRpcLoadBalancerExamplesTestSuite struct {
	test_suites.BaseTestSuite
	Tc          configs.TestContainer
	Web3Actions web3_actions.Web3Actions
}

func (s *AdaptiveRpcLoadBalancerExamplesTestSuite) SetupTest() {
	// points dir to test/configs
	s.Tc = configs.InitLocalTestConfigs()
}

func TestAdaptiveRpcLoadBalancerExamplesTestSuite(t *testing.T) {
	suite.Run(t, new(AdaptiveRpcLoadBalancerExamplesTestSuite))
}
