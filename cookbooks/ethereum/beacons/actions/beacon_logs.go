package beacon_actions

import (
	"context"
	"path"

	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/client/zeus_req_types/pods"
	v1 "k8s.io/api/core/v1"
)

func (b *BeaconActionsClient) PrintConsensusClientPodLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	b.PrintReqJson(par)
	par.ContainerName = "zeus-consensus-client"
	filter := strings_filter.FilterOpts{Contains: client_consts.ZeusConsensusClient}
	logOpts := &v1.PodLogOptions{Container: b.ConsensusClient}
	par.LogOpts = logOpts
	par.FilterOpts = &filter

	resp, err := b.GetPodLogs(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("PrintConsensusClientPodLogs: GetPodLogs")
		return nil, err
	}
	b.PrintPath.FnOut = b.ConsensusClient + "_logs"
	b.PrintPath.DirOut = path.Join(b.PrintPath.DirIn, "/consensus_client")
	err = b.PrintPath.Print(resp, "json")
	return resp, err
}

func (b *BeaconActionsClient) PrintExecClientPodLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	b.PrintReqJson(par)
	par.ContainerName = b.ExecClient
	logOpts := &v1.PodLogOptions{Container: b.ExecClient}
	par.LogOpts = logOpts

	filter := strings_filter.FilterOpts{Contains: client_consts.ZeusExecClient}
	par.FilterOpts = &filter
	resp, err := b.GetPodLogs(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("PrintExecClientPodLogs: GetPodLogs")
		return nil, err
	}
	b.PrintPath.FnOut = b.ExecClient + "_logs"
	b.PrintPath.DirOut = path.Join(b.PrintPath.DirIn, "/exec_client")
	err = b.PrintPath.Print(resp, "logs")
	return resp, err
}
