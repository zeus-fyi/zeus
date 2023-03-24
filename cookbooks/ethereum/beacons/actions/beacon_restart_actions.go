package beacon_actions

import (
	"context"

	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

func (b *BeaconActionsClient) RestartConsensusClientPods(ctx context.Context) ([]byte, error) {
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		ContainerName:         b.ConsensusClient,
		FilterOpts: &strings_filter.FilterOpts{
			StartsWith: client_consts.ZeusConsensusClient,
		},
	}
	b.PrintReqJson(par)
	resp, err := b.DeletePods(ctx, par)
	b.PrintRespJson(resp)
	return resp, err
}

func (b *BeaconActionsClient) RestartExecClientPods(ctx context.Context) ([]byte, error) {
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		ContainerName:         b.ExecClient,
		FilterOpts: &strings_filter.FilterOpts{
			StartsWith: client_consts.ZeusExecClient,
		},
	}
	b.PrintReqJson(par)
	resp, err := b.DeletePods(ctx, par)
	b.PrintRespJson(resp)
	return resp, err
}
