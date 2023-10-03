package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
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
	if sts.Spec.Template.Spec.Containers == nil {
		sts.Spec.Template.Spec.Containers = []v1core.Container{}
	}
	if sts.Spec.Template.Spec.InitContainers == nil {
		sts.Spec.Template.Spec.InitContainers = []v1core.Container{}
	}
	for _, contDriver := range s.ContainerDrivers {
		if contDriver.IsAppendContainer {
			if contDriver.IsInitContainer {
				sts.Spec.Template.Spec.InitContainers = append(sts.Spec.Template.Spec.InitContainers, contDriver.Container)
			} else {
				sts.Spec.Template.Spec.Containers = append(sts.Spec.Template.Spec.Containers, contDriver.Container)
			}
		}
	}
	// init containers
	var ifc []v1core.Container
	for _, c := range sts.Spec.Template.Spec.InitContainers {
		if v, ok := s.ContainerDrivers[c.Name]; ok && v.IsInitContainer {
			if v.IsDeleteContainer {
				continue
			}
			v.SetContainerConfigs(&c)
			ifc = append(ifc, c)
		}
	}
	sts.Spec.Template.Spec.InitContainers = ifc
	// containers
	var fc []v1core.Container
	for _, c := range sts.Spec.Template.Spec.Containers {
		if v, ok := s.ContainerDrivers[c.Name]; ok {
			if v.IsDeleteContainer {
				continue
			}
			v.SetContainerConfigs(&c)
			fc = append(fc, c)
		}
	}
	sts.Spec.Template.Spec.Containers = fc

	// pvcs
	if sts.Spec.VolumeClaimTemplates == nil {
		sts.Spec.VolumeClaimTemplates = []v1core.PersistentVolumeClaim{}
	}
	if s.PVCDriver != nil {
		tmp := sts.Spec.VolumeClaimTemplates
		tmp = s.PVCDriver.CustomPVCS(tmp)
		sts.Spec.VolumeClaimTemplates = tmp
	}
}
