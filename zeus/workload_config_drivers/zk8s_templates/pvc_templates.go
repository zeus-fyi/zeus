package zk8s_templates

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPvcTemplate(pvcTemplate PVCTemplate) v1.PersistentVolumeClaim {
	storageReq := v1.ResourceList{"storage": resource.MustParse(pvcTemplate.StorageSizeRequest)}
	accessMode := pvcTemplate.AccessMode
	if len(accessMode) == 0 {
		accessMode = string(v1.ReadWriteOnce)
	}
	return v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: pvcTemplate.Name,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{v1.PersistentVolumeAccessMode(pvcTemplate.AccessMode)},
			Resources: v1.VolumeResourceRequirements{
				Requests: storageReq,
			},
			StorageClassName: pvcTemplate.StorageClassName,
		},
	}
}
