package beacon_actions

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types"
)

func (b *BeaconActionsClient) ReplaceAndRestartConfigsConsensusClient(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	_, err := b.ReplaceConfigsConsensusClient(ctx, par.TopologyDeployRequest)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReplaceAndRestartConfigsConsensusClient: ReplaceConfigsConsensusClient")
		return nil, err
	}
	resp, err := b.RestartConsensusClientPods(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReplaceAndRestartConfigsConsensusClient: RestartConsensusClientPods")
		return nil, err
	}
	return resp, err
}

func (b *BeaconActionsClient) ReplaceAndRestartConfigsExecClient(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	_, err := b.ReplaceConfigsExecClient(ctx, par.TopologyDeployRequest)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReplaceAndRestartConfigsExecClient: ReplaceConfigsExecClient")
		return nil, err
	}
	resp, err := b.RestartExecClientPods(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReplaceAndRestartConfigsExecClient: RestartExecClientPods")
		return nil, err
	}
	return resp, err
}

func (b *BeaconActionsClient) ReplaceConfigsConsensusClient(ctx context.Context, tar zeus_req_types.TopologyDeployRequest) (zeus_resp_types.TopologyDeployStatus, error) {
	b.ConfigPaths.FnIn = b.ConsensusClient
	b.ConfigPaths.DirIn += "/consensus_client/alt_configs"
	resp, err := b.DeployReplace(ctx, b.ConfigPaths, tar)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReplaceConfigsConsensusClient")
		return zeus_resp_types.TopologyDeployStatus{}, err
	}
	return resp, err
}

func (b *BeaconActionsClient) ReplaceConfigsExecClient(ctx context.Context, tar zeus_req_types.TopologyDeployRequest) (zeus_resp_types.TopologyDeployStatus, error) {
	b.ConfigPaths.FnIn = b.ExecClient
	b.ConfigPaths.DirIn += "/exec_client/alt_configs"
	resp, err := b.DeployReplace(ctx, b.ConfigPaths, tar)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ReplaceConfigsExecClient")
		return zeus_resp_types.TopologyDeployStatus{}, err
	}
	return resp, err
}
