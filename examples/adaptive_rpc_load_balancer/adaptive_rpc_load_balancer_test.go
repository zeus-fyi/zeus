package adaptive_rpc_load_balancer_examples

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/gochain/web3/accounts"
	"github.com/zeus-fyi/zeus/examples/adaptive_rpc_load_balancer/smart_contract_library"
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
	Tc configs.TestContainer
}

func (s *AdaptiveRpcLoadBalancerExamplesTestSuite) SetupTest() {
	// points dir to test/configs
	s.Tc = configs.InitLocalTestConfigs()
}

func CreateLocalUser(ctx context.Context, bearer, sessionID string) web3_actions.Web3Actions {
	acc, err := accounts.CreateAccount()
	if err != nil {
		panic(err)
	}
	w3a := web3_actions.NewWeb3ActionsClientWithAccount(LoadBalancerAddress, acc)
	w3a.AddAnvilSessionLockHeader(sessionID)
	w3a.AddBearerToken(bearer)
	nvB := (*hexutil.Big)(smart_contract_library.EtherMultiple(10000))
	w3a.Dial()
	defer w3a.Close()
	err = w3a.SetBalance(ctx, w3a.Address().String(), *nvB)
	if err != nil {
		panic(err)
	}
	return w3a
}

func TestAdaptiveRpcLoadBalancerExamplesTestSuite(t *testing.T) {
	suite.Run(t, new(AdaptiveRpcLoadBalancerExamplesTestSuite))
}
