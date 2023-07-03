package zeus_client

import (
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_config_map_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/config_maps"
)

var (
	DeployConsensusClientKnsReq = zeus_req_types.TopologyDeployRequest{
		TopologyID: 0,
		CloudCtxNs: BeaconCloudCtxNs,
	}
	BeaconCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "ephemeral", // set with your own namespace
		Env:           "production",
	}
)

func (t *ZeusClientTestSuite) TestConfigMapKeySwap() {

	cmr := zeus_config_map_reqs.ConfigMapActionRequest{
		TopologyDeployRequest: DeployConsensusClientKnsReq,
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
		TopologyDeployRequest: DeployConsensusClientKnsReq,
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
