package zeus_client

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"

// TestDeployStatus will return the latest status updates for a topology
func (t *ZeusClientTestSuite) TestDeployStatus() {
	tr := zeus_req_types.TopologyRequest{}
	tr.TopologyID = deployKnsReq.TopologyID
	resp, err := t.ZeusTestClient.ReadDeployStatusUpdates(ctx, tr)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}
