package zeus_topology_config_drivers

import (
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func CreateComputeResourceRequirementsLimit(cpuSize, memSize string) v1Core.ResourceRequirements {
	rr := v1Core.ResourceRequirements{
		Requests: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
		Limits: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
	}
	return rr
}

func CreateDiskResourceRequirementsLimit(diskSize string) v1Core.ResourceRequirements {
	rr := v1Core.ResourceRequirements{
		Requests: v1Core.ResourceList{
			"storage": resource.MustParse(diskSize),
		},
	}
	return rr
}
