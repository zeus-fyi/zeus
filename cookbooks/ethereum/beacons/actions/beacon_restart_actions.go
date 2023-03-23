package beacon_actions

import (
	"context"

	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

func (b *BeaconActionsClient) RestartConsensusClientPods(ctx context.Context) ([]byte, error) {
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		PodName:               b.ConsensusClient,
		ContainerName:         b.ConsensusClient,
	}
	b.PrintReqJson(par)
	resp, err := b.DeletePods(ctx, par)
	b.PrintRespJson(resp)
	return resp, err
}

func (b *BeaconActionsClient) RestartExecClientPods(ctx context.Context) ([]byte, error) {
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		PodName:               b.ExecClient,
		ContainerName:         b.ExecClient,
	}
	b.PrintReqJson(par)
	resp, err := b.DeletePods(ctx, par)
	b.PrintRespJson(resp)
	return resp, err
}
