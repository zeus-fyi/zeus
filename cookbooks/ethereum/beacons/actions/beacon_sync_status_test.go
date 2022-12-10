package beacon_actions

import client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"

func (t *BeaconActionsTestSuite) TestSyncStatusRequest() {
	t.ConsensusClient = client_consts.Lighthouse

	resp, err := t.GetConsensusClientSyncStatus(ctx)
	t.Assert().Nil(err)
	t.Assert().NotEmpty(resp)
}
