package artemis_client

import (
	artemis_endpoints "github.com/zeus-fyi/zeus/pkg/artemis/client/endpoints"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
)

func (t *ArtemisClientTestSuite) TestSendEthEndpoints() {
	mainnetSendEtherTxEndpoint := artemis_endpoints.EthereumMainnetSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(mainnetSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumMainnet))

	goerliSendEtherTxEndpoint := artemis_endpoints.EthereumGoerliSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(goerliSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumGoerli))

	ephemeralSendEtherTxEndpoint := artemis_endpoints.EthereumEphemeralSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(ephemeralSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumEphemeral))
}

func (t *ArtemisClientTestSuite) TestSendEthPayload() {
	sendEthTx := artemis_req_types.SendEtherPayload{
		TransferArgs: artemis_req_types.TransferArgs{
			Amount:    signing_automation_ethereum.signing_automation_ethereum.ValidatorDeposit32Eth,
			ToAddress: t.TestAccount2.Address(),
		},
		GasPriceLimits: artemis_req_types.GasPriceLimits{},
	}
	resp, err := t.ArtemisTestClient.SendEther(ctx, sendEthTx, ArtemisEthereumEphemeral)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}
