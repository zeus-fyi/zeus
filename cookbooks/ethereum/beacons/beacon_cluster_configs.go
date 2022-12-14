package beacon_cookbooks

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

func ConfigEphemeralLighthouseGethStakingBeacon(cp, ep filepaths.Path) {
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := cp.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	if err != nil {
		panic(err)
	}
	EphemeralConsensusClientLighthouseConfig(inf)
	err = inf.PrintWorkload(cp)
	if err != nil {
		panic(err)
	}

	infe := topology_workloads.NewTopologyBaseInfraWorkload()
	err = ep.WalkAndApplyFuncToFileType(".yaml", infe.DecodeK8sWorkload)
	if err != nil {
		panic(err)
	}
	EphemeralExecClientGethConfig(infe)
	err = infe.PrintWorkload(ep)
	if err != nil {
		panic(err)
	}
}
