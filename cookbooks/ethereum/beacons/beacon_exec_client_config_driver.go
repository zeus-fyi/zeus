package beacon_cookbooks

import (
	"errors"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	execClient         = "zeus-exec-client"
	execClientDiskName = "exec-client-storage"
	execClientDiskSize = "10Gi"

	GethEphemeral         = "gethEphemeral"
	downloadGethEphemeral = "downloadGethEphemeral"
	gethDockerImage       = "ethereum/client-go:v1.10.26"
)

func EphemeralExecClientGethConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.ConfigMap != nil {
		m := make(map[string]string)
		m = inf.ConfigMap.Data
		vSrc, ok := inf.ConfigMap.Data[GethEphemeral+".sh"]
		if ok {
			m[start+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[downloadGethEphemeral+".sh"]
		if ok {
			m[download+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
	}
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == execClient {
				inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = gethDockerImage
			}
		}

		for i, v := range inf.StatefulSet.Spec.VolumeClaimTemplates {
			if v.Name == execClientDiskName {
				q, err := resource.ParseQuantity(execClientDiskSize)
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
