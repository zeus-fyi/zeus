package beacon_cookbooks

import (
	"errors"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

const (
	consensusClient             = "zeus-consensus-client"
	lighthouseDockerImage       = "sigp/lighthouse:v3.3.0-modern"
	downloadLighthouseEphemeral = "downloadLighthouseEphemeral"
	LighthouseEphemeral         = "lighthouseEphemeral"
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
	}
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == consensusClient {
				inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = lighthouseDockerImage
			}
		}
	}
}
