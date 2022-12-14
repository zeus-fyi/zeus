package beacon_cookbooks

import (
	"errors"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

const (
	downloadLighthouse  = "downloadLighthouseEphemeral"
	lighthouseEphemeral = "lighthouseEphemeral"
	start               = "start"
	download            = "download"
)

func EphemeralConsensusClientLighthouseConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.ConfigMap != nil {
		m := make(map[string]string)
		m = inf.ConfigMap.Data
		vSrc, ok := inf.ConfigMap.Data[lighthouseEphemeral+".sh"]
		if ok {
			m[start+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[downloadLighthouse+".sh"]
		if ok {
			m[download+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
	}
}
