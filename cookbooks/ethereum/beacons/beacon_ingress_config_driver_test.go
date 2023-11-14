package ethereum_beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
)

func (t *BeaconCookbookTestSuite) TestBeaconIngressConfigDriver() {
	p := IngressChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Require().NotEmpty(inf)

	EphemeralIngressConfig(inf)

	t.Require().NotEmpty(inf.Ingress)

	err = inf.PrintWorkload(p)
	t.Require().Nil(err)
}
