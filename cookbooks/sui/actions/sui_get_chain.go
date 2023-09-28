package sui_actions

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	GetChainId = JsonRpcReq{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "sui_getChainIdentifier",
		Params:  []interface{}{},
	}
)

func (s *SuiActionsClient) GetChainId(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs) ([]JsonRpcResult, error) {
	req := GetChainId
	resp, err := s.SendRpcPayload(ctx, cloudCtxNs, req)
	if err != nil {
		return []JsonRpcResult{}, err
	}
	ss := make([]JsonRpcResult, len(resp.ReplyBodies))
	i := 0
	for _, v := range resp.ReplyBodies {
		err = json.Unmarshal(v, &ss[i])
		if err != nil {
			log.Err(err).Msg("SuiActionsClient: GetChainId")
			return ss, err
		}
		i += 1
	}
	return ss, err
}
