package sui_actions

import (
	"context"
	"fmt"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	pods_client "github.com/zeus-fyi/zeus/zeus/z_client/workloads/pods"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
	zeus_pods_resp "github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/pods"
)

const (
	DefaultSuiRpcPortNumber = 9000
)

type SuiActionsClient struct {
	pods_client.PodsClient

	RpcPortNumber int

	PrintPath filepaths.Path
}

type JsonRpcReq struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type JsonRpcResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func InitSuiClient(p pods_client.PodsClient) SuiActionsClient {
	return SuiActionsClient{
		PodsClient:    p,
		RpcPortNumber: DefaultSuiRpcPortNumber,
	}
}

func (s *SuiActionsClient) SendRpcPayload(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs, payload any) (zeus_pods_resp.ClientResp, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "POST",
		Endpoint:   "/",
		Ports:      []string{fmt.Sprintf("%d", s.RpcPortNumber)},
		Payload:    payload,
	}
	filter := strings_filter.FilterOpts{Contains: "sui"}
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: zeus_req_types.TopologyDeployRequest{
			CloudCtxNs: cloudCtxNs,
		},
		Action:     zeus_pods_reqs.PortForwardToAllMatchingPods,
		FilterOpts: &filter,
		ClientReq:  &cliReq,
	}
	return s.PortForwardReqToPods(ctx, par)
}
