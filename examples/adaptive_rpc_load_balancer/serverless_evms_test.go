package adaptive_rpc_load_balancer_examples

import (
	"github.com/zeus-fyi/zeus/examples/adaptive_rpc_load_balancer/smart_contract_library"
)

func (s *AdaptiveRpcLoadBalancerExamplesTestSuite) TestHardhatLocalNetwork() {
	erc20Abi := smart_contract_library.MustLoadERC20AbiPayload(ctx)
	s.Assert().NotEmpty(erc20Abi)

	//newAccount, err := accounts.ParsePrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	//s.Assert().Nil(err)
}
