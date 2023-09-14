package sui_actions

import (
	pods_client "github.com/zeus-fyi/zeus/zeus/z_client/workloads/pods"
)

type SuiActionsClient struct {
	pods_client.PodsClient
}
