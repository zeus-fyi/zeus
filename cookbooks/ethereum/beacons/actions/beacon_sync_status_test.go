package beacon_actions

func (t *BeaconActionsTestSuite) TestSyncStatusRequest() {
	t.ConsensusClient = "lighthouse"

	resp, err := t.GetConsensusClientSyncStatus(ctx)
	t.Assert().Nil(err)
	t.Assert().NotEmpty(resp)
}
