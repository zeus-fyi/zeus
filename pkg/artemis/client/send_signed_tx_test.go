package artemis_client

import (
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
	artemis_endpoints "github.com/zeus-fyi/zeus/pkg/artemis/client/endpoints"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/web3/signing_automation/ethereum"
)

func (t *ArtemisClientTestSuite) TestSignedTxEndpoints() {
	mainnetSendSignedTxEndpoint := artemis_endpoints.EthereumMainnetSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(mainnetSendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumMainnet))

	goerliSendSignedTxEndpoint := artemis_endpoints.EthereumGoerliSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(goerliSendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumGoerli))

	ephemerySendSignedTxEndpoint := artemis_endpoints.EthereumEphemeralSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(ephemerySendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumEphemeral))
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
		Params: []interface{}{t.TestAccount2.Address(), signing_automation_ethereum.Finney},
	}
	testClient1 := signing_automation_ethereum.NewWeb3Client(t.NodeURL, t.TestAccount1.Account)
	signedTx, err := testClient1.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	t.Require().Nil(err)
	t.Require().NotNil(signedTx)
	payload := artemis_req_types.SignedTxPayload{Transaction: *signedTx}
	resp, err := t.ArtemisTestClient.SendSignedTx(ctx, &payload, ArtemisEthereumGoerli)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}
