package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
)

type StatefulSetDriver struct {
	ReplicaCount     *int32
	ContainerDrivers map[string]ContainerDriver
	PVCDriver        *PersistentVolumeClaimsConfigDriver
}

func NewStatefulSetDriver() StatefulSetDriver {
	return StatefulSetDriver{ContainerDrivers: make(map[string]ContainerDriver)}
}

func (s *StatefulSetDriver) SetStatefulSetConfigs(sts *v1.StatefulSet) {
	if sts == nil {
		return
	}

	if s.ReplicaCount != nil {
		sts.Spec.Replicas = s.ReplicaCount
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
	// pvcs
	if s.PVCDriver != nil {
		tmp := sts.Spec.VolumeClaimTemplates
		s.PVCDriver.CustomPVCS(tmp)
		sts.Spec.VolumeClaimTemplates = tmp
	}
}
