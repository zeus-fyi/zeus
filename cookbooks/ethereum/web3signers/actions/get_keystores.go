package ethereum_web3signer_actions

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type GetKeystoresResp struct {
	Data []struct {
		ValidatingPubkey string `json:"validating_pubkey"`
		DerivationPath   string `json:"derivation_path"`
		Readonly         bool   `json:"readonly"`
	} `json:"data"`
}

func (w *Web3SignerActionsClient) GetKeystores(ctx context.Context, kns zeus_common_types.CloudCtxNs) ([]GetKeystoresResp, error) {
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "GET",
		Endpoint:   "eth/v1/keystores",
		Ports:      Web3SignerPorts,
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
	respImport, err := w.PortForwardReqToPods(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetKeystores")
		return []GetKeystoresResp{}, err
	}
	ss := make([]GetKeystoresResp, len(respImport.ReplyBodies))
	i := 0
	for _, v := range respImport.ReplyBodies {
		err = json.Unmarshal(v, &ss[i])
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("ImportKeystores")
			return ss, err
		}
		i += 1
	}
	return ss, err
}
