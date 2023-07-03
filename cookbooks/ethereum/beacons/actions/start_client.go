package beacon_actions

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	zeus_config_map_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/config_maps"
)

func (b *BeaconActionsClient) StartConsensusClient(ctx context.Context) ([]byte, error) {
	return b.StartClient(ctx, client_consts.ZeusConfigMapConsensusClient, client_consts.Lighthouse)
}

func (b *BeaconActionsClient) StartExecClient(ctx context.Context) ([]byte, error) {
	return b.StartClient(ctx, client_consts.ZeusConfigMapExecClient, client_consts.Geth)
}

func (b *BeaconActionsClient) StartClient(ctx context.Context, cmName, clientName string) ([]byte, error) {
	cmReq := zeus_config_map_reqs.ConfigMapActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		Action:                zeus_config_map_reqs.SetOrCreateKeyFromExisting,
		ConfigMapName:         cmName,
		Keys: zeus_config_map_reqs.KeySwap{
			KeyOne: fmt.Sprintf("%s.sh", clientName),
			KeyTwo: "start.sh",
		},
		FilterOpts: nil,
	}
	respCm, err := b.SetOrCreateKeyFromConfigMapKey(ctx, cmReq)
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("configMap", respCm).Msg("PauseConsensusClient: SetOrCreateKeyFromConfigMapKey")
		return nil, err
	}
	if client_consts.IsConsensusClient(clientName) {
		b.ConsensusClient = clientName
		resp, cerr := b.RestartConsensusClientPods(ctx)
		if cerr != nil {
			log.Ctx(ctx).Err(cerr).Msg("PauseConsensusClient: RestartConsensusClientPods")
			return nil, cerr
		}
		return resp, cerr
	}
	if client_consts.IsExecClient(clientName) {
		b.ExecClient = clientName
		resp, cerr := b.RestartExecClientPods(ctx)
		if cerr != nil {
			log.Ctx(ctx).Err(cerr).Msg("PauseConsensusClient: RestartExecClientPods")
			return nil, cerr
		}
		return resp, cerr
	}
	return nil, errors.New("invalid consensus exec client supplied")
}
