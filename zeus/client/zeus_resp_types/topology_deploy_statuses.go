package zeus_resp_types

import (
	"time"

	"github.com/zeus-fyi/zeus/zeus/client/zeus_common_types"
)

type TopologyDeployStatuses struct {
	Slice []TopologyDeployStatus
}

type TopologyDeployStatus struct {
	TopologyID     int       `json:"topologyID"`
	TopologyName   string    `json:"topologyName"`
	TopologyStatus string    `json:"topologyStatus"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CloudCtxNs     zeus_common_types.CloudCtxNs
}

type ClusterStatus struct {
	ClusterName string `json:"clusterName"`
	Status      string `json:"status"`
}
