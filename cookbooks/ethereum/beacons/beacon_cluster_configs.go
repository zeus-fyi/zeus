package ethereum_beacon_cookbooks

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
)

func ConfigEphemeralLighthouseGethBeacon(cp, ep, ing filepaths.Path, withIngress bool) {
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

	if withIngress {
		ingr := topology_workloads.NewTopologyBaseInfraWorkload()
		err = ing.WalkAndApplyFuncToFileType(".yaml", ingr.DecodeK8sWorkload)
		if err != nil {
			panic(err)
		}
		EphemeralIngressConfig(ingr)
	}
}
