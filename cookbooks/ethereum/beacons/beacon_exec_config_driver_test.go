package beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

func (t *BeaconCookbookTestSuite) TestExecClientBeaconConfigDriver() {
	p := beaconExecClientChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Require().NotEmpty(inf)

	EphemeralConsensusClientGethConfig(inf)

	t.Require().NotEmpty(inf.ConfigMap)
	t.Assert().Equal(inf.ConfigMap.Data["start.sh"], inf.ConfigMap.Data[gethEphemeral+".sh"])

	t.Require().NotEmpty(inf.StatefulSet)

	count := 0 // verifies execClient key is found
	for _, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
		if c.Name == execClient {
			t.Assert().Equal(gethDockerImage, c.Image)
			count += 1
		}
	}
	t.Require().Equal(1, count)

	t.Assert().Equal(inf.ConfigMap.Data[download+".sh"], inf.ConfigMap.Data[downloadGethEphemeral+".sh"])
	p.DirOut = "./ethereum/beacons/infra/processed_exec_client"

	err = inf.PrintWorkload(p)
	t.Require().Nil(err)
}
