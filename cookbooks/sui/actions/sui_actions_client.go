package sui_actions

import (
	pods_client "github.com/zeus-fyi/zeus/zeus/z_client/workloads/pods"
)

type SuiActionsClient struct {
	pods_client.PodsClient
}

// https://docs.sui.io/sui-jsonrpc#sui_getCheckpoints

const (
	GetCheckpointPayloadJsonRpc = `{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "sui_getLatestCheckpointSequenceNumber",
  "params": []
}`
)
