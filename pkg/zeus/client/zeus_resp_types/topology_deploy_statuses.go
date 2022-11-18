package zeus_resp_types

import (
	"time"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
)

type TopologyDeployStatuses struct {
	Slice []TopologyDeployStatus
}

type TopologyDeployStatus struct {
	TopologyID     int       `db:"topology_id" json:"topologyID"`
	TopologyName   string    `db:"topology_name" json:"topologyName"`
	TopologyStatus string    `db:"topology_status" json:"topologyStatus"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
	CloudCtxNs     zeus_common_types.CloudCtxNs
}
