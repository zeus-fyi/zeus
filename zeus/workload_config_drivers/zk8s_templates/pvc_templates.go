package zk8s_templates

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func GetPvcTemplate(pvcTemplate PVCTemplate) v1.PersistentVolumeClaim {
	storageReq := v1.ResourceList{"storage": resource.MustParse(pvcTemplate.StorageSizeRequest)}

	return v1.PersistentVolumeClaim{
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{v1.PersistentVolumeAccessMode(pvcTemplate.AccessMode)},
			Resources: v1.VolumeResourceRequirements{
				Requests: storageReq,
			},
			VolumeName: pvcTemplate.Name,
		},
	}
}
