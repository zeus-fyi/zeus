package sui_actions

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

// https://docs.sui.io/sui-jsonrpc#sui_getCheckpoints

var (
	GetLatestCheckpointSequenceNumberReq = JsonRpcReq{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []interface{}{},
	}
	CheckpointReq = JsonRpcReq{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "sui_getCheckpoint",
		Params:  []interface{}{},
	}
)

func (s *SuiActionsClient) GetCheckpoint(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs, ckpId string) ([]SuiCheckpointResponse, error) {
	req := CheckpointReq
	req.Params = append(req.Params, ckpId)
	resp, err := s.SendRpcPayload(ctx, cloudCtxNs, req)
	if err != nil {
		return []SuiCheckpointResponse{}, err
	}
	ss := make([]SuiCheckpointResponse, len(resp.ReplyBodies))
	i := 0
	for _, v := range resp.ReplyBodies {
		err = json.Unmarshal(v, &ss[i])
		if err != nil {
			log.Err(err).Msg("SuiActionsClient: GetCheckpoint")
			return ss, err
		}
		i += 1
	}
	return ss, err
}

func (s *SuiActionsClient) GetLatestCheckpoint(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs) ([]SuiCheckpointResponse, error) {
	req := GetLatestCheckpointSequenceNumberReq
	resp, err := s.SendRpcPayload(ctx, cloudCtxNs, req)
	if err != nil {
		return []SuiCheckpointResponse{}, err
	}
	ss := make([]SuiCheckpointResponse, len(resp.ReplyBodies))
	i := 0
	for _, v := range resp.ReplyBodies {
		err = json.Unmarshal(v, &ss[i])
		if err != nil {
			log.Err(err).Msg("SuiActionsClient: GetCheckpoint")
			return ss, err
		}
		i += 1
	}
	return ss, err
}
