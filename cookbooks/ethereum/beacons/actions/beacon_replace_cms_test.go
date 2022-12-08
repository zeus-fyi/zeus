package beacon_actions

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"

var deployConsensusKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 1669171061885689000,
	CloudCtxNs: beaconCloudCtxNs,
}

func (t *BeaconActionsTestSuite) TestReplaceConsensusConfigs() {
	t.ConsensusClient = "lighthouse"
	_, err := t.ReplaceConfigsConsensusClient(ctx, deployConsensusKnsReq)
	t.Assert().Nil(err)
}

var deployExecKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 1669171045611326000,
	CloudCtxNs: beaconCloudCtxNs,
}

func (t *BeaconActionsTestSuite) TestReplaceExecConfigs() {
	t.ExecClient = "geth"
	_, err := t.ReplaceConfigsExecClient(ctx, deployExecKnsReq)
	t.Assert().Nil(err)
}
