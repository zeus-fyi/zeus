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
}

func (t *BeaconCookbookTestSuite) TestExecClientBeaconConfigDriver() {
	p := beaconExecClientChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Require().NotEmpty(inf)
}
