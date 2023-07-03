package ethereum_validator_cookbooks

import (
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	validatorClient                = "zeus-validators"
	validatorClientStorageDiskName = "validator-client-storage"
	validatorClientStorageDiskSize = "100Mi"

	lighthouseDockerImage = "sigp/lighthouse:v3.3.0-modern"
)

func EphemeralValidatorClientLighthouseConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == validatorClient {
				inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = lighthouseDockerImage
			}
		}
		for i, v := range inf.StatefulSet.Spec.VolumeClaimTemplates {
			if v.Name == validatorClientStorageDiskName {
				q, err := resource.ParseQuantity(validatorClientStorageDiskSize)
				if err != nil {
					panic(err)
				}
				for j, _ := range inf.StatefulSet.Spec.VolumeClaimTemplates[i].Spec.Resources.Requests {
					tmp := inf.StatefulSet.Spec.VolumeClaimTemplates[i].Spec.Resources.Requests[j]
					tmp.Set(q.Value())
					inf.StatefulSet.Spec.VolumeClaimTemplates[i].Spec.Resources.Requests[j] = tmp
				}
			}
		}
	}
}
