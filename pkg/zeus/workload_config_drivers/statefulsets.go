package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
)

type StatefulSetDriver struct {
	ContainerDrivers map[string]ContainerDriver
}

func NewStatefulSetDriver() StatefulSetDriver {
	return StatefulSetDriver{ContainerDrivers: make(map[string]ContainerDriver)}
}

func (s *StatefulSetDriver) SetStatefulSetConfigs(sts *v1.StatefulSet) {
	if sts == nil {
		return
	}

	// init containers
	for i, c := range sts.Spec.Template.Spec.InitContainers {
		if v, ok := s.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			sts.Spec.Template.Spec.InitContainers[i] = c
		}
	}

	// containers
	for i, c := range sts.Spec.Template.Spec.Containers {
		if v, ok := s.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			sts.Spec.Template.Spec.Containers[i] = c
		}
	}
}
