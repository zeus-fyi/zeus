package nodes

import v1 "k8s.io/api/core/v1"

type NodeGroup struct {
	Taints []v1.Taint `json:"taints,omitempty"`
}

type Node struct {
	Memory        int     `db:"memory" json:"memory"`
	Vcpus         float64 `db:"vcpus" json:"vcpus"`
	Disk          int     `db:"disk" json:"disk"`
	DiskUnits     string  `db:"disk_units" json:"diskUnits"`
	DiskType      string  `db:"disk_type" json:"diskType"`
	PriceHourly   float64 `db:"price_hourly" json:"priceHourly"`
	Region        string  `db:"region" json:"region"`
	CloudProvider string  `db:"cloud_provider" json:"cloudProvider"`
	ResourceID    int     `db:"resource_id" json:"resourceID"`
	Description   string  `db:"description" json:"description"`
	Slug          string  `db:"slug" json:"slug"`
	MemoryUnits   string  `db:"memory_units" json:"memoryUnits"`
	PriceMonthly  float64 `db:"price_monthly" json:"priceMonthly"`
	Gpus          int     `db:"gpus" json:"gpus"`
	GpuType       string  `db:"gpu_type" json:"gpuType"`
}

type NodesSlice []Node
