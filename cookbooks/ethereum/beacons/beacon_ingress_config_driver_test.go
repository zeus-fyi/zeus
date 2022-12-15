package beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

func (t *BeaconCookbookTestSuite) TestBeaconIngressConfigDriver() {
	p := ingressChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Require().NotEmpty(inf)

	EphemeralIngressConfig(inf)

	t.Require().NotEmpty(inf.Ingress)

	err = inf.PrintWorkload(p)
	t.Require().Nil(err)
}
