package beacon_actions

func (t *BeaconActionsTestSuite) TestConsensusClientPodKill() {
	t.ConsensusClient = "lighthouse"
	_, err := t.RestartConsensusClientPods(ctx, basePar)
	t.Assert().Nil(err)
}

func (t *BeaconActionsTestSuite) TestExecClientPodKill() {
	t.ExecClient = "geth"
	_, err := t.RestartExecClientPods(ctx, basePar)
	t.Assert().Nil(err)
}

func (t *BeaconActionsTestSuite) TestRestartBeaconPodKill() {
	t.ConsensusClient = "lighthouse"
	t.ExecClient = "geth"

	_, err := t.RestartExecClientPods(ctx, basePar)
	t.Assert().Nil(err)

	_, err = t.RestartConsensusClientPods(ctx, basePar)
	t.Assert().Nil(err)
}
