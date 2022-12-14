package beacon_cookbooks

import (
	"errors"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

const (
	execClient            = "zeus-exec-client"
	gethDockerImage       = "ethereum/client-go:v1.10.26"
	downloadGethEphemeral = "downloadGethEphemeral"
	GethEphemeral         = "gethEphemeral"
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
	}
}
