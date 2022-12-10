package zeus_client

import (
	beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	zeus_config_map_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/config_maps"
)

func (t *ZeusClientTestSuite) TestConfigMapKeySwap() {
	cmr := zeus_config_map_reqs.ConfigMapActionRequest{
		TopologyDeployRequest: beacon_cookbooks.DeployConsensusClientKnsReq,
		Action:                zeus_config_map_reqs.KeySwapAction,
		ConfigMapName:         "cm-consensus-client",
		Keys: zeus_config_map_reqs.KeySwap{
			KeyOne: "start.sh",
			KeyTwo: "pause.sh",
		},
		FilterOpts: nil,
	}
	resp, err := t.ZeusTestClient.SwapConfigMapKeys(ctx, cmr)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *ZeusClientTestSuite) TestConfigMapSetFromExistingKey() {
	cmr := zeus_config_map_reqs.ConfigMapActionRequest{
		TopologyDeployRequest: beacon_cookbooks.DeployConsensusClientKnsReq,
		Action:                zeus_config_map_reqs.KeySwapAction,
		ConfigMapName:         "cm-consensus-client",
		Keys: zeus_config_map_reqs.KeySwap{
			KeyOne: "start.sh",
			KeyTwo: "pause.sh",
		},
		FilterOpts: nil,
	}
	// keyOne=keyToCopy, keyTwo=keyToSetOrCreateFromCopy
	resp, err := t.ZeusTestClient.SetOrCreateKeyFromConfigMapKey(ctx, cmr)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}
