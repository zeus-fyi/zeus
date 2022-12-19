package artemis_client

import (
	artemis_endpoints "github.com/zeus-fyi/zeus/pkg/artemis/client/endpoints"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
	"github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
)

func (t *ArtemisClientTestSuite) TestSendEthEndpoints() {
	goerliSendEtherTxEndpoint := artemis_endpoints.EthereumGoerliSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(goerliSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumGoerli))

	mainnetSendEtherTxEndpoint := artemis_endpoints.EthereumMainnetSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(mainnetSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumMainnet))
}

func (t *ArtemisClientTestSuite) TestSendEthPayload() {
	sendEthTx := artemis_req_types.SendEtherPayload{
		Amount:         signing_automation_ethereum.Finney,
		ToAddress:      t.TestAccount2.Address(),
		GasPriceLimits: artemis_req_types.GasPriceLimits{},
	}
	resp, err := t.ArtemisTestClient.SendEther(ctx, sendEthTx, ArtemisEthereumGoerli)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}
