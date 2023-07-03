package ethereum_web3signer_actions

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type Web3SignerActionsClient struct {
	zeus_client.ZeusClient
}

var Web3SignerPorts = []string{"9000:9000"}

type KeystoreImportResp struct {
	Data []struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

func (w *Web3SignerActionsClient) ImportKeystores(ctx context.Context, kns zeus_common_types.CloudCtxNs, p filepaths.Path, pw string) ([]KeystoreImportResp, error) {
	ks := Web3SignerKeystores{}
	ks.ReadKeystoreDirAndAppendPw(ctx, p, pw)

	tmp := Web3SignerKeystores{
		Keystores:          []string{},
		Passwords:          []string{},
		SlashingProtection: "",
	}
	for _, k := range ks.Keystores {
		if len(tmp.Keystores) > 99 {
			cliReq := zeus_pods_reqs.ClientRequest{
				MethodHTTP: "POST",
				Endpoint:   "eth/v1/keystores",
				Ports:      Web3SignerPorts,
				Payload:    tmp,
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
				log.Ctx(ctx).Err(err).Msg("ImportKeystores")
				return []KeystoreImportResp{}, err
			}
			ss := make([]KeystoreImportResp, len(respImport.ReplyBodies))
			j := 0
			for _, v := range respImport.ReplyBodies {
				err = json.Unmarshal(v, &ss[j])
				if err != nil {
					log.Ctx(ctx).Err(err).Msg("ImportKeystores")
					return ss, err
				}
				j += 1
			}
			tmp = Web3SignerKeystores{
				Keystores:          []string{},
				Passwords:          []string{},
				SlashingProtection: "",
			}
			time.Sleep(5 * time.Second)
		} else {
			tmp.Keystores = append(tmp.Keystores, k)
			tmp.Passwords = append(tmp.Passwords, ks.Passwords[0])
		}
	}
	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "POST",
		Endpoint:   "eth/v1/keystores",
		Ports:      Web3SignerPorts,
		Payload:    tmp,
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
		log.Ctx(ctx).Err(err).Msg("ImportKeystores")
		return []KeystoreImportResp{}, err
	}
	ss := make([]KeystoreImportResp, len(respImport.ReplyBodies))
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
