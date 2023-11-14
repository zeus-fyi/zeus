package config_overrides

import (
	"encoding/json"

	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaimsConfigDriver struct {
	AppendPVC                    map[string]bool
	PersistentVolumeClaimDrivers map[string]v1.PersistentVolumeClaim
}

func (p *PersistentVolumeClaimsConfigDriver) CustomPVCS(pvcs []v1.PersistentVolumeClaim) []v1.PersistentVolumeClaim {
	if p.PersistentVolumeClaimDrivers == nil {
		return pvcs
	}
	for j, pvc := range pvcs {
		if customPVC, ok := p.PersistentVolumeClaimDrivers[pvc.Name]; ok {
			if customPVC.Spec.StorageClassName != nil {
				pvcs[j].Spec.StorageClassName = customPVC.Spec.StorageClassName
			}
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
			if customPVC.Spec.Resources.Limits != nil {
				b, err := json.Marshal(customPVC.Spec.Resources.Limits)
				if err != nil {
					panic(err)
				}
				pvcs[j].Spec.Resources.Limits = nil
				err = json.Unmarshal(b, &pvcs[j].Spec.Resources.Limits)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	for k, _ := range p.PersistentVolumeClaimDrivers {
		if v, ok := p.AppendPVC[k]; ok && v {
			pvcs = append(pvcs, p.PersistentVolumeClaimDrivers[k])
		}
	}
	return pvcs
}
