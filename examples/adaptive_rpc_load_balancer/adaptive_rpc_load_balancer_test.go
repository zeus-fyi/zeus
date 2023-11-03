package adaptive_rpc_load_balancer_examples

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/gochain/web3/accounts"
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

	// local hardhat user account default
	newAccount, err := accounts.ParsePrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	s.Assert().Nil(err)
	s.Require().NotNil(newAccount)
	pubKey := newAccount.Address().String()
	s.Require().Equal("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", pubKey)

	s.Web3Actions = web3_actions.NewWeb3ActionsClientWithAccount(LoadBalancerAddress, newAccount)
	s.Web3Actions.AddBearerToken(s.Tc.Bearer)
}

func TestAdaptiveRpcLoadBalancerExamplesTestSuite(t *testing.T) {
	suite.Run(t, new(AdaptiveRpcLoadBalancerExamplesTestSuite))
}
