package topology_workloads

type TopologyCreateClassResponse struct {
	ClusterName string `json:"name,omitempty"`
	ClassID     int    `json:"classID"`
	Status      string `json:"status,omitempty"`
}
