package beacon_actions

func (t *BeaconActionsTestSuite) TestPrintConsensusLogs() {
	t.ConsensusClient = "lighthouse"
	_, err := t.PrintConsensusClientPodLogs(ctx, basePar)
	t.Assert().Nil(err)
}

func (t *BeaconActionsTestSuite) TestPrintExecPodsLogs() {
	t.ExecClient = "geth"
	_, err := t.PrintExecClientPodLogs(ctx, basePar)
	t.Assert().Nil(err)
}
