package sui_actions

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	GetLatestCheckpointSequenceNumberReq = JsonRpcReq{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []interface{}{},
	}
	GetCheckpointReq = JsonRpcReq{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "sui_getCheckpoint",
		Params:  []interface{}{},
	}
	// GetCheckpointsReq https://docs.sui.io/sui-jsonrpc#sui_getCheckpoints
	GetCheckpointsReq = JsonRpcReq{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "sui_getCheckpoints",
		Params:  []interface{}{},
	}
)

type SuiCheckpointsRange struct {
	Cursor          string `json:"cursor,omitempty"`
	Limit           int    `json:"limit,omitempty"`
	DescendingOrder bool   `json:"descendingOrder,omitempty"`
}

func (s *SuiActionsClient) GetCheckpoints(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs, ckpRange SuiCheckpointsRange) ([]SuiCheckpointsPageResponse, error) {
	req := GetCheckpointsReq
	req.Params = []interface{}{ckpRange.Cursor, ckpRange.Limit, ckpRange.DescendingOrder}
	resp, err := s.SendRpcPayload(ctx, cloudCtxNs, req)
	if err != nil {
		return []SuiCheckpointsPageResponse{}, err
	}
	ss := make([]SuiCheckpointsPageResponse, len(resp.ReplyBodies))
	i := 0
	for _, v := range resp.ReplyBodies {
		err = json.Unmarshal(v, &ss[i])
		if err != nil {
			log.Err(err).Msg("SuiActionsClient: GetCheckpoints")
			return ss, err
		}
		i += 1
	}
	return ss, err
}

func (s *SuiActionsClient) GetCheckpoint(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs, ckpId string) ([]SuiCheckpointResponse, error) {
	req := GetCheckpointReq
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

func (s *SuiActionsClient) GetLatestCheckpointSeqNumber(ctx context.Context, cloudCtxNs zeus_common_types.CloudCtxNs) ([]JsonRpcResult, error) {
	req := GetLatestCheckpointSequenceNumberReq
	resp, err := s.SendRpcPayload(ctx, cloudCtxNs, req)
	if err != nil {
		return []JsonRpcResult{}, err
	}
	ss := make([]JsonRpcResult, len(resp.ReplyBodies))
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
