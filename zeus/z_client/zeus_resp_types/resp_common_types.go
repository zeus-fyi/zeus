package zeus_resp_types

import (
	"database/sql"
	"time"
)

type TopologyCreateResponse struct {
	TopologyID int `json:"topologyID"`
}

type DeployStatus struct {
	TopologyID     int       `json:"topologyID"`
	TopologyStatus string    `json:"topologyStatus"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type ReadTopologiesMetadata struct {
	TopologyID       int            `json:"topologyID"`
	TopologyName     string         `json:"topologyName"`
	ChartName        string         `json:"chartName"`
	ChartVersion     string         `json:"chartVersion"`
	ChartDescription sql.NullString `json:"chartDescription"`
}

type ReadTopologiesMetadataGroup struct {
	Slice []ReadTopologiesMetadata
}
