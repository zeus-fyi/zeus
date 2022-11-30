package artemis_client

import (
	"github.com/zeus-fyi/gochain/web3/web3_actions"
	artemis_endpoints "github.com/zeus-fyi/zeus/pkg/artemis/client/endpoints"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
)

func (t *ArtemisClientTestSuite) TestSignedTxEndpoints() {
	goerliSendSignedTxEndpoint := artemis_endpoints.EthereumGoerliSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(goerliSendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumGoerli))

	mainnetSendSignedTxEndpoint := artemis_endpoints.EthereumMainnetSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(mainnetSendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumMainnet))
}

const LinkGoerliContractAddr = "0x326C977E6efc84E512bB9C30f76E30c160eD06FB"

func (t *ArtemisClientTestSuite) TestSignedTxPayload() {
	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: LinkGoerliContractAddr,
		ContractFile:      web3_actions.ERC20,
		MethodName:        web3_actions.Transfer,
		SendEtherPayload: web3_actions.SendEtherPayload{
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},
		Params: []interface{}{t.TestAccount2.Address(), Finney},
	}
	testClient1 := NewWeb3Client(t.NodeURL, t.TestAccount1.Account)
	signedTx, err := testClient1.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	t.Require().Nil(err)
	t.Require().NotNil(signedTx)
	payload := artemis_req_types.SignedTxPayload{Transaction: signedTx}
	resp, err := t.ArtemisTestClient.SendSignedTx(ctx, payload, ArtemisEthereumGoerli)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}
