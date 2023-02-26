package zeus_topology_config_drivers

import (
	"encoding/json"

	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaimsConfigDriver struct {
	PersistentVolumeClaimDrivers map[string]v1.PersistentVolumeClaim
}

func (p *PersistentVolumeClaimsConfigDriver) CustomPVCS(pvcs []v1.PersistentVolumeClaim) []v1.PersistentVolumeClaim {
	if pvcs == nil {
		return pvcs
	}
	for j, pvc := range pvcs {
		if customPVC, ok := p.PersistentVolumeClaimDrivers[pvc.Name]; ok {
			if customPVC.Spec.Resources.Requests != nil {
				b, err := json.Marshal(customPVC.Spec.Resources.Requests)
				if err != nil {
					panic(err)
				}
				pvcs[j].Spec.Resources.Requests = nil
				err = json.Unmarshal(b, &pvcs[j].Spec.Resources.Requests)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return pvcs
}
