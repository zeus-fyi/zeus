package beacon_actions

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

func (b *BeaconActionsClient) GetConsensusClientSyncStatus(ctx context.Context) ([]client_consts.ConsensusClientSyncStatus, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "GET",
		Endpoint:   "eth/v1/node/syncing",
		Ports:      client_consts.GetClientBeaconPortsHTTP(b.ConsensusClient),
	}
	filter := strings_filter.FilterOpts{Contains: b.ConsensusClient}
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		Action:                zeus_pods_reqs.PortForwardToAllMatchingPods,
		ClientReq:             &cliReq,
		FilterOpts:            &filter,
	}
	resp, err := b.ZeusClient.PortForwardReqToPods(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetConsensusClientSyncStatus")
		return []client_consts.ConsensusClientSyncStatus{}, err
	}
	ss := make([]client_consts.ConsensusClientSyncStatus, len(resp.ReplyBodies))
	i := 0
	for _, v := range resp.ReplyBodies {
		err = json.Unmarshal(v, &ss[i])
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("GetConsensusClientSyncStatus")
			return ss, err
		}
		i += 1
	}
	return ss, err
}

func (b *BeaconActionsClient) GetExecClientSyncStatus(ctx context.Context) ([]client_consts.ExecClientSyncStatus, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP:      "POST",
		Endpoint:        "/",
		Ports:           client_consts.GetAnyClientApiPorts(b.ExecClient),
		EndpointHeaders: headers,
		Payload:         `{"method":"eth_syncing","params":[],"id":1,"jsonrpc":"2.0"}`,
	}
	filter := strings_filter.FilterOpts{Contains: b.ExecClient}
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: b.BeaconKnsReq,
		Action:                zeus_pods_reqs.PortForwardToAllMatchingPods,
		ClientReq:             &cliReq,
		FilterOpts:            &filter,
	}
	resp, err := b.ZeusClient.PortForwardReqToPods(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetExecClientSyncStatus")
		return []client_consts.ExecClientSyncStatus{}, err
	}
	es := make([]client_consts.ExecClientSyncStatus, len(resp.ReplyBodies))
	i := 0
	for _, v := range resp.ReplyBodies {
		err = json.Unmarshal(v, &es[i])
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("GetExecClientSyncStatus")
			return es, err
		}
		i += 1
	}
	return es, err
}
