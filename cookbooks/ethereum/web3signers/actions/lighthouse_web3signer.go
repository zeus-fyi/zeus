package ethereum_web3signer_actions

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
	zeus_pods_resp "github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/pods"
)

func (w *Web3SignerActionsClient) EnableWeb3SignerLighthouse(ctx context.Context, kns zeus_common_types.CloudCtxNs, w3req []LighthouseWeb3SignerRequest, authToken string) (zeus_pods_resp.ClientResp, error) {
	h := make(map[string]string)
	authToken = strings.TrimPrefix(authToken, "\"")
	authToken = strings.TrimSuffix(authToken, "\"\n")
	h["Authorization"] = authToken
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP:      "POST",
		Endpoint:        client_consts.LighthouseWeb3SignerAPIEndpoint,
		Ports:           client_consts.LighthouseValidatorClientPorts,
		Payload:         w3req,
		EndpointHeaders: h,
	}
	filter := strings_filter.FilterOpts{Contains: "validators"}
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: zeus_req_types.TopologyDeployRequest{
			CloudCtxNs: kns,
		},
		Action:     zeus_pods_reqs.PortForwardToAllMatchingPods,
		ClientReq:  &cliReq,
		FilterOpts: &filter,
	}
	respImport, err := w.PortForwardReqToPods(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("EnableWeb3SignerLighthouse")
		return zeus_pods_resp.ClientResp{}, err
	}
	return respImport, err
}

func (w *Web3SignerActionsClient) GetLighthouseAuth(ctx context.Context, kns zeus_common_types.CloudCtxNs) (zeus_pods_resp.ClientResp, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "GET",
		Endpoint:   client_consts.HerculesLighthouseAuthTokenEndpoint,
		Ports:      client_consts.HerculesPorts,
	}
	filter := strings_filter.FilterOpts{Contains: "validators"}
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: zeus_req_types.TopologyDeployRequest{
			CloudCtxNs: kns,
		},
		Action:     zeus_pods_reqs.PortForwardToAllMatchingPods,
		ClientReq:  &cliReq,
		FilterOpts: &filter,
	}
	respImport, err := w.PortForwardReqToPods(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("EnableWeb3SignerLighthouse")
		return zeus_pods_resp.ClientResp{}, err
	}
	return respImport, err
}

func (w *Web3SignerActionsClient) GetLighthouseRemoteKeystores(ctx context.Context, kns zeus_common_types.CloudCtxNs) (zeus_pods_resp.ClientResp, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "GET",
		Endpoint:   "eth/v1/remotekeys",
		Ports:      client_consts.LighthouseValidatorClientPorts,
	}
	filter := strings_filter.FilterOpts{Contains: "web3signer"}
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: zeus_req_types.TopologyDeployRequest{
			CloudCtxNs: kns,
		},
		Action:     zeus_pods_reqs.PortForwardToAllMatchingPods,
		ClientReq:  &cliReq,
		FilterOpts: &filter,
	}
	resp, err := w.PortForwardReqToPods(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetRemoteKeystores")
		return zeus_pods_resp.ClientResp{}, err
	}
	return resp, err
}
