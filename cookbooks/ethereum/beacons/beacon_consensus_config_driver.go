package ethereum_beacon_cookbooks

import (
	"errors"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	zeusConsensusClient               = "zeus-consensus-client"
	consensusStorageDiskName          = "consensus-client-storage"
	consensusStorageDiskSizeEphemeral = "2Gi"
	consensusStorageDiskSizeGoerli    = "100Gi"
	consensusStorageDiskSizeMainnet   = "300Gi"

	LighthouseMainnet           = "lighthouse"
	LodestarEphemeral           = "lodestarEphemeral"
	LighthouseEphemeral         = "lighthouseEphemeral"
	downloadLighthouseEphemeral = "downloadLighthouseEphemeral"

	lighthouseDockerImage        = "sigp/lighthouse:v4.2.0"
	lighthouseDockerImageCapella = "sigp/lighthouse:capella"

	LodestarGoerli      = "lodestarGoerli"
	lodestarDockerImage = "chainsafe/lodestar:v1.7.2"
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
			if c.Name == zeusConsensusClient {
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
