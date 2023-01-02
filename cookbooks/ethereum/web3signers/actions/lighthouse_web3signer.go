package ethereum_web3signer_actions

import (
	"context"

	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	zeus_pods_resp "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/pods"
)

type LighthouseWeb3SignerRequest struct {
	Enable                bool   `json:"enable"`
	Description           string `json:"description"`
	SuggestedFeeRecipient string `json:"suggested_fee_recipient"`
	VotingPublicKey       string `json:"voting_public_key"`
	Graffiti              string `json:"graffiti,omitempty"`
	Url                   string `json:"url,omitempty"`
	RootCertificatePath   string `json:"root_certificate_path,omitempty"`
	RequestTimeoutMs      int    `json:"request_timeout_ms,omitempty"`
}

func (w *Web3SignerActionsClient) EnableWeb3SignerLighthouse(ctx context.Context, kns zeus_common_types.CloudCtxNs, w3req []LighthouseWeb3SignerRequest) (zeus_pods_resp.ClientResp, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "POST",
		Endpoint:   client_consts.LighthouseWeb3SignerAPIEndpoint,
		Ports:      client_consts.LighthouseValidatorClientPorts,
		Payload:    w3req,
	}
	filter := strings_filter.FilterOpts{Contains: "consensus-client"}
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
