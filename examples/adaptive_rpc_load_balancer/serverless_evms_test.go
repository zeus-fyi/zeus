package adaptive_rpc_load_balancer_examples

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/zeus-fyi/zeus/examples/adaptive_rpc_load_balancer/smart_contract_library"
	web3_actions "github.com/zeus-fyi/zeus/pkg/artemis/web3/client"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/web3/signing_automation/ethereum"
)

func (s *AdaptiveRpcLoadBalancerExamplesTestSuite) setupMintToken(mintAmount *big.Int) (string, web3_actions.SendContractTxPayload) {
	// mintable contract
	abiDefAndByteCode := smart_contract_library.TokenJson
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(abiDefAndByteCode), &m)
	s.Assert().Nil(err)

	s.Require().NotNil(m["abi"])
	s.Require().NotNil(m["bytecode"])
	abiDef := m["abi"]

	abiBin, err := json.Marshal(abiDef)
	s.Assert().Nil(err)

	byteCode := m["bytecode"].(string)
	abiFile := signing_automation_ethereum.MustReadAbiString(ctx, string(abiBin))
	tokenPayload := web3_actions.SendContractTxPayload{
		SendEtherPayload: web3_actions.SendEtherPayload{
			GasPriceLimits: web3_actions.GasPriceLimits{
				GasLimit:  5000000,
				GasTipCap: big.NewInt(100000000),
				GasFeeCap: big.NewInt(1000000000 * 2),
			},
		},
		ContractABI: abiFile,
		Params:      []interface{}{mintAmount},
	}

	return byteCode, tokenPayload
}

func (s *AdaptiveRpcLoadBalancerExamplesTestSuite) TestHardhatLocalNetwork() {
	sessionID := fmt.Sprintf("%s-%s", "local-network-session", uuid.New().String())
	s.Web3Actions.AddAnvilSessionLockHeader(sessionID)

	defer func(sessionID string) {
		err := s.Web3Actions.EndAnvilSession()
		s.Require().Nil(err)
	}(sessionID)
	// deploy a contract
	ether := big.NewInt(1e18)
	mintAmount := new(big.Int).Mul(big.NewInt(10000000), ether)

	pubkey := s.Web3Actions.Address().String()
	etherBalance, err := s.Web3Actions.GetBalance(ctx, pubkey, nil)
	s.Require().Nil(err)
	s.Require().NotZero(etherBalance.Int64())

	byteCode, tokenPayload := s.setupMintToken(mintAmount)
	tx, err := s.Web3Actions.GetSignedDeployTxToCallFunctionWithArgs(ctx, byteCode, &tokenPayload)
	s.Require().Nil(err)
	s.Require().NotNil(tx)

	err = s.Web3Actions.SendSignedTransaction(ctx, tx)
	s.Require().Nil(err)

	rx, err := s.Web3Actions.GetTxReceipt(ctx, tx.Hash().String())
	s.Require().NotNil(rx)
	s.Require().Nil(err)

	scAddr := rx.ContractAddress.String()

	tokenBalance, err := s.Web3Actions.ReadERC20TokenBalance(ctx, scAddr, s.Web3Actions.Address().String())
	s.Require().Nil(err)
	s.Assert().NotZero(tokenBalance)
	s.Assert().Equal(mintAmount.String(), tokenBalance.String())
}
