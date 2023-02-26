package ethereum_beacon_cookbooks

import (
	"errors"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	consensusClient                   = "zeus-consensus-client"
	consensusStorageDiskName          = "consensus-client-storage"
	consensusStorageDiskSizeEphemeral = "2Gi"

	LighthouseEphemeral         = "lighthouseEphemeral"
	downloadLighthouseEphemeral = "downloadLighthouseEphemeral"

	lighthouseDockerImage        = "sigp/lighthouse:v3.3.0-modern"
	lighthouseDockerImageCapella = "sigp/lighthouse:capella"
)

func EphemeralConsensusClientLighthouseConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.ConfigMap != nil {
		m := make(map[string]string)
		m = inf.ConfigMap.Data
		vSrc, ok := inf.ConfigMap.Data[LighthouseEphemeral+".sh"]
		if ok {
			m[start+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[downloadLighthouseEphemeral+".sh"]
		if ok {
			m[download+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[herculesEphemeral+".sh"]
		if ok {
			m[hercules+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
	}
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == consensusClient {
				inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = lighthouseDockerImageCapella
			}
		}
		for i, v := range inf.StatefulSet.Spec.VolumeClaimTemplates {
			if v.Name == consensusStorageDiskName {
				q, err := resource.ParseQuantity(consensusStorageDiskSizeEphemeral)
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
