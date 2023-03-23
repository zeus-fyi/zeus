package beacon_actions

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_config_map_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/config_maps"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

// set your own topologyID here after uploading a chart workload
var beaconKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: beaconCloudCtxNs,
}

var beaconCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "ethereum", // set with your own namespace
	Env:           "dev",
}

var basePar = zeus_pods_reqs.PodActionRequest{
	TopologyDeployRequest: beaconKnsReq,
	PodName:               "",
	FilterOpts:            nil,
	ClientReq:             nil,
	DeleteOpts:            nil,
}

func (b *BeaconActionsClient) PauseClient(ctx context.Context, cmName, clientName string) ([]byte, error) {
	cmr := zeus_config_map_reqs.ConfigMapActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		Action:                zeus_config_map_reqs.SetOrCreateKeyFromExisting,
		ConfigMapName:         cmName,
		Keys: zeus_config_map_reqs.KeySwap{
			KeyOne: "pause.sh",
			KeyTwo: "start.sh",
		},
		FilterOpts: nil,
	}
	respCm, err := b.SetOrCreateKeyFromConfigMapKey(ctx, cmr)
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("configMap", respCm).Msg("PauseConsensusClient: SetOrCreateKeyFromConfigMapKey")
		return nil, err
	}
	if client_consts.IsConsensusClient(clientName) {
		b.ConsensusClient = clientName
		basePar.TopologyDeployRequest = b.BeaconKnsReq
		resp, cerr := b.RestartConsensusClientPods(ctx)
		if cerr != nil {
			log.Ctx(ctx).Err(cerr).Msg("PauseConsensusClient: RestartConsensusClientPods")
			return nil, cerr
		}
		return resp, cerr
	}
	if client_consts.IsExecClient(clientName) {
		b.ExecClient = clientName
		basePar.TopologyDeployRequest = b.BeaconKnsReq
		resp, cerr := b.RestartExecClientPods(ctx)
		if cerr != nil {
			log.Ctx(ctx).Err(cerr).Msg("PauseConsensusClient: RestartExecClientPods")
			return nil, cerr
		}
		return resp, cerr
	}
	return nil, errors.New("invalid consensus exec client supplied")
}
