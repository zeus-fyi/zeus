package cluster_node_resources

import v1 "k8s.io/api/core/v1"

type Node struct {
	ResourceID  int          `json:"resourceID"`
	NodeDetails *NodeDetails `json:"nodeDetails,omitempty"`
	Taints      []v1.Taint   `json:"taints,omitempty"`
}

type NodeDetails struct {
	Memory        int     `json:"memory"`
	Vcpus         float64 `json:"vcpus"`
	Disk          int     `json:"disk"`
	DiskUnits     string  `json:"diskUnits"`
	PriceHourly   float64 `json:"priceHourly"`
	Region        string  `json:"region"`
	CloudProvider string  `json:"cloudProvider"`
	Description   string  `json:"description"`
	Slug          string  `json:"slug"`
	MemoryUnits   string  `json:"memoryUnits"`
	PriceMonthly  float64 `json:"priceMonthly"`
	Gpus          int     `json:"gpus"`
	GpuType       string  `json:"gpuType"`
}

type NodesGroup struct {
	NodeMap map[string]Node
}
