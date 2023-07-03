package beacon_actions

import "github.com/zeus-fyi/zeus/zeus/client/zeus_common_types"

func (t *BeaconActionsTestSuite) TestConsensusClientPodKill() {
	t.BeaconActionsClient.BeaconKnsReq.CloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "athena-beacon-goerli",
		Env:           "",
	}
	_, err := t.RestartConsensusClientPods(ctx)
	t.Assert().Nil(err)
}

func (t *BeaconActionsTestSuite) TestExecClientPodKill() {
	t.BeaconActionsClient.BeaconKnsReq.CloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "athena-beacon-goerli",
		Env:           "",
	}
	_, err := t.RestartExecClientPods(ctx)
	t.Assert().Nil(err)
}

func (t *BeaconActionsTestSuite) TestRestartBeaconPodKill() {
	_, err := t.RestartExecClientPods(ctx)
	t.Assert().Nil(err)

	_, err = t.RestartConsensusClientPods(ctx)
	t.Assert().Nil(err)
}
