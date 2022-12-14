package beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

func (t *BeaconCookbookTestSuite) TestConsensusClientBeaconConfigDriver() {
	p := beaconConsensusClientChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Require().NotEmpty(inf)

	EphemeralConsensusClientLighthouseConfig(inf)

	t.Require().NotEmpty(inf.ConfigMap)
	t.Assert().Equal(inf.ConfigMap.Data["start.sh"], inf.ConfigMap.Data[lighthouseEphemeral+".sh"])

	t.Require().NotEmpty(inf.StatefulSet)

	count := 0 // verifies consensusClient key is found
	for _, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
		if c.Name == consensusClient {
			t.Assert().Equal(lighthouseDockerImage, c.Image)
			count += 1
		}
	}
	t.Require().Equal(1, count)

	t.Assert().Equal(inf.ConfigMap.Data[download+".sh"], inf.ConfigMap.Data[downloadLighthouseEphemeral+".sh"])
	p.DirOut = "./ethereum/beacons/infra/processed_consensus_client"

	err = inf.PrintWorkload(p)
	t.Require().Nil(err)
}
