package topology_workloads

type TopologyCreateClassResponse struct {
	ClusterClassName string `json:"clusterClassName,omitempty"`
	ClassID          int    `json:"classID"`
	Status           string `json:"status,omitempty"`
}
