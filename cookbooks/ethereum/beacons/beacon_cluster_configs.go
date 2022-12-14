package beacon_cookbooks

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"

func ConfigEphemeralLighthouseGethStakingBeacon() {
	pc := beaconConsensusClientChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := pc.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	if err != nil {
		panic(err)
	}
	EphemeralConsensusClientLighthouseConfig(inf)
	pc.DirOut = "./ethereum/beacons/infra/processed_consensus_client"
	err = inf.PrintWorkload(pc)
	if err != nil {
		panic(err)
	}

	pe := beaconExecClientChartPath
	infe := topology_workloads.NewTopologyBaseInfraWorkload()
	err = pe.WalkAndApplyFuncToFileType(".yaml", infe.DecodeK8sWorkload)
	if err != nil {
		panic(err)
	}
	EphemeralExecClientGethConfig(infe)
	pe.DirOut = "./ethereum/beacons/infra/processed_exec_client"
	err = infe.PrintWorkload(pe)
	if err != nil {
		panic(err)
	}
}
