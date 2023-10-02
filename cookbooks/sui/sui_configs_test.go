package sui_cookbooks

import (
	"fmt"

	do_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/do"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (t *SuiCookbookTestSuite) generateFromConfigDriverBuilder(cfg SuiConfigOpts) {
	cd := GetSuiClientClusterDef(cfg)
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
}

func (t *SuiCookbookTestSuite) TestSuiTestnetCfg() {
	cfg := SuiConfigOpts{
		DownloadSnapshot: false,
		WithIngress:      false,
		CloudProvider:    "do",
		Network:          testnet,
	}
	t.generateFromConfigDriverBuilder(cfg)

	p := suiMasterChartPath
	p.DirIn = "./sui/node/custom_sui"
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.Nil(inf.Ingress)
	t.Nil(inf.Deployment)
	t.NotNil(inf.StatefulSet)
	t.NotNil(inf.Service)
	t.NotNil(inf.ConfigMap)

	t.NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates)
	t.Len(inf.StatefulSet.Spec.VolumeClaimTemplates, 1)
	t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests)
	t.Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage(), resource.MustParse(testnetDiskSize))
	t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName)
	t.Require().Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName, do_nvme.DoStorageClass)

	seen := 0
	for _, ct := range inf.StatefulSet.Spec.Template.Spec.Containers {
		if ct.Name == Sui {
			t.Equal(*ct.Resources.Requests.Cpu(), resource.MustParse(cpuCoresTestnet))
			t.Equal(*ct.Resources.Requests.Memory(), resource.MustParse(memorySizeTestnet))
			t.Equal(*ct.Resources.Limits.Cpu(), resource.MustParse(cpuCoresTestnet))
			t.Equal(*ct.Resources.Limits.Memory(), resource.MustParse(memorySizeTestnet))
			seen++
		}
	}
	for _, ct := range inf.StatefulSet.Spec.Template.Spec.InitContainers {
		if ct.Name == "init-snapshots" {
			t.Require().Len(ct.Args, 2)
			t.Equal(NoDownload+".sh", ct.Args[1])
		}
	}
	// init-snapshots
	t.Require().Equal(seen, 1)
	for k, v := range inf.ConfigMap.Data {
		fmt.Println(k, v)
	}
}
