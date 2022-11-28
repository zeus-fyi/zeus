package artemis_client

import (
	artemis_endpoints "github.com/zeus-fyi/zeus/pkg/artemis/client/endpoints"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
)

func (t *ArtemisClientTestSuite) TestSignedTxEndpoints() {
	goerliSendEtherTxEndpoint := artemis_endpoints.EthereumGoerliSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(goerliSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumGoerli))

	goerliSendSignedTxEndpoint := artemis_endpoints.EthereumGoerliSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(goerliSendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumGoerli))

	mainnetSendEtherTxEndpoint := artemis_endpoints.EthereumMainnetSendEtherOrchestrationV1BetaPath
	t.Assert().Equal(mainnetSendEtherTxEndpoint, getSendEtherEndpoint(ArtemisEthereumMainnet))

	mainnetSendSignedTxEndpoint := artemis_endpoints.EthereumMainnetSendSignedTxOrchestrationV1BetaPath
	t.Assert().Equal(mainnetSendSignedTxEndpoint, getSendSignedTxEndpoint(ArtemisEthereumMainnet))
}

func (t *ArtemisClientTestSuite) TestSignedTxPayload() {
	sendSignedTx := artemis_req_types.SignedTxPayload{}
	resp, err := t.ArtemisTestClient.SendSignedTx(ctx, sendSignedTx, ArtemisEthereumGoerli)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}
