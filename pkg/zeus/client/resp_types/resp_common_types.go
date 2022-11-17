package resp_types

import (
	"database/sql"
	"time"
)

type TopologyCreateResponse struct {
	ID int `json:"id"`
}

type DeployStatus struct {
	TopologyID     int       `db:"topology_id" json:"topologyID"`
	TopologyStatus string    `db:"topology_status" json:"topologyStatus"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
}

type ReadTopologiesMetadata struct {
	TopologyID       int            `db:"topology_id" json:"topologyID"`
	TopologyName     string         `db:"topology_name" json:"topologyName"`
	ChartName        string         `db:"chart_name" json:"chartName"`
	ChartVersion     string         `db:"chart_version" json:"chartVersion"`
	ChartDescription sql.NullString `db:"chart_description" json:"chartDescription"`
}

type ReadTopologiesMetadataGroup struct {
	Slice []ReadTopologiesMetadata
}
