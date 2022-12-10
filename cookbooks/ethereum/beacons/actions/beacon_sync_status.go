package beacon_actions

import (
	"context"

	beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	zeus_pods_resp "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/pods"
)

func (b *BeaconActionsClient) GetConsensusClientSyncStatus(ctx context.Context) (zeus_pods_resp.ClientResp, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "GET",
		Endpoint:   "eth/v1/node/syncing",
		Ports:      []string{"5052:5052"},
	}
	filter := strings_filter.FilterOpts{Contains: b.ConsensusClient}
	routeHeader := beacon_cookbooks.DeployConsensusClientKnsReq
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: routeHeader,
		Action:                zeus_pods_reqs.PortForwardToAllMatchingPods,
		ClientReq:             &cliReq,
		FilterOpts:            &filter,
	}

	resp, err := b.ZeusClient.PortForwardReqToPods(ctx, par)
	return resp, err
}
