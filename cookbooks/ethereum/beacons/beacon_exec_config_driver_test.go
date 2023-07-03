package ethereum_beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (t *BeaconCookbookTestSuite) TestExecClientBeaconConfigDriver() {
	p := BeaconExecClientChartPath
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Require().NotEmpty(inf)

	EphemeralExecClientGethConfig(inf)

	t.Require().NotEmpty(inf.ConfigMap)
	t.Assert().Equal(inf.ConfigMap.Data["start.sh"], inf.ConfigMap.Data[GethEphemeral+".sh"])

	t.Require().NotEmpty(inf.StatefulSet)

	count := 0 // verifies execClient key is found
	for _, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
		if c.Name == zeusExecClient {
			t.Assert().Equal(gethDockerImage, c.Image)
			count += 1
		}
	}
	t.Require().Equal(1, count)

	t.Assert().Equal(inf.ConfigMap.Data[download+".sh"], inf.ConfigMap.Data[downloadGethEphemeral+".sh"])
	p.DirOut = "./ethereum/beacons/infra/processed_exec_client"

	err = inf.PrintWorkload(p)
	t.Require().Nil(err)

	q, err := resource.ParseQuantity(execClientDiskSizeEphemeral)
	if err != nil {
		panic(err)
	}

	for _, v := range inf.StatefulSet.Spec.VolumeClaimTemplates {
		if v.Name == "exec-client-storage" {
			t.Assert().True(v.Spec.Resources.Requests.Storage().Equal(q))
		}
	}
}
