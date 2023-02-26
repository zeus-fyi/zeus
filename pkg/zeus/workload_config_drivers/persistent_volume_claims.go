package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaimConfigDriver struct {
	PersistentVolumeClaimDrivers map[string]v1.PersistentVolumeClaim
}

func (p *PersistentVolumeClaimConfigDriver) CustomPVC(pvc *v1.PersistentVolumeClaim) {
	if pvc == nil {
		return
	}
	if customPVC, ok := p.PersistentVolumeClaimDrivers[pvc.Name]; ok {
		if customPVC.Spec.Resources.Requests != nil {
			pvc.Spec.Resources.Requests = customPVC.Spec.Resources.Requests
		}
	}
}
