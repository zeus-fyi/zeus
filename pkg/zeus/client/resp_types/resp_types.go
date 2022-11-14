package resp_types

import (
	"database/sql"
	"time"

	v1 "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
)

type TopologyChartWorkload struct {
	*v1core.Service       `json:"service"`
	*v1core.ConfigMap     `json:"configMap"`
	*v1.Deployment        `json:"deployment"`
	*v1.StatefulSet       `json:"statefulSet"`
	*v1networking.Ingress `json:"ingress"`
}

type DeployStatus struct {
	TopologyID     int       `db:"topology_id" json:"topologyID"`
	TopologyStatus string    `db:"topology_status" json:"topologyStatus"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
}

type ReadTopologiesMetadata struct {
	TopologyID       int            `db:"topology_id" json:"topology_id"`
	TopologyName     string         `db:"topology_name" json:"topology_name"`
	ChartName        string         `db:"chart_name" json:"chart_name"`
	ChartVersion     string         `db:"chart_version" json:"chart_version"`
	ChartDescription sql.NullString `db:"chart_description" json:"chart_description"`
}

type ReadTopologiesMetadataGroup struct {
	Slice []ReadTopologiesMetadata
}
