package req_types

type TopologyCreateRequest struct {
	TopologyName     string `json:"topologyName"`
	ChartName        string `json:"chartName"`
	ChartDescription string `json:"chartDescription,omitempty"`
	Version          string `json:"version"`
}

type TopologyDeployRequest struct {
	TopologyID    int    `json:"topologyID"`
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	Context       string `json:"context"`
	Namespace     string `json:"namespace"`
	Env           string `json:"env"`
}

type TopologyRequest struct {
	TopologyID int `db:"topology_id" json:"topology_id"`
}
