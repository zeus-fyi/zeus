package sui_cookbooks

import (
	"encoding/json"
	"fmt"

	yaml_fileio "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/yaml"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	aws_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/aws"
	do_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/do"
	gcp_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/gcp"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (t *SuiCookbookTestSuite) generateFromConfigDriverBuilder(cfg SuiConfigOpts) []zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition {
	cd := GetSuiClientClusterDef(cfg)
	cd.DisablePrint = false
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)

	return sbDefs
}

func (t *SuiCookbookTestSuite) TestReadFullNodeConfig() {
	p := suiMasterChartPath
	p.FnIn = "fullnode.yaml"
	p.DirIn = "./sui/node/sui_config"
	fip := p.FileInPath()
	nodeCfg, err := yaml_fileio.ReadYamlConfig(fip)
	t.Require().Nil(err)
	t.Require().NotNil(nodeCfg)
	m := make(map[string]interface{})
	err = json.Unmarshal(nodeCfg, &m)
	t.Require().Nil(err)
	dataDir := "/mnt/fast-disks"
	for k, _ := range m {
		if k == "db-path" {
			m[k] = dataDir
			fmt.Println(m[k])
		}
		if k == "genesis" {
			m[k] = map[string]interface{}{
				"genesis-file-location": dataDir + "/genesis.blob",
			}
			fmt.Println(m[k])
		}
	}
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

func (t *SuiCookbookTestSuite) TestSuiMainnetCfg() {
	cfg := SuiConfigOpts{
		DownloadSnapshot: true,
		WithIngress:      false,
		CloudProvider:    "aws",
		Network:          mainnet,
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
	t.Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage(), resource.MustParse(mainnetDiskSize))
	t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName)
	t.Require().Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName, aws_nvme.AwsStorageClass)

	seen := 0
	for _, ct := range inf.StatefulSet.Spec.Template.Spec.Containers {
		if ct.Name == Sui {
			t.Equal(*ct.Resources.Requests.Cpu(), resource.MustParse(cpuCores))
			t.Equal(*ct.Resources.Requests.Memory(), resource.MustParse(memorySize))
			t.Equal(*ct.Resources.Limits.Cpu(), resource.MustParse(cpuCores))
			t.Equal(*ct.Resources.Limits.Memory(), resource.MustParse(memorySize))
			seen++
		}
	}
	for _, ct := range inf.StatefulSet.Spec.Template.Spec.InitContainers {
		if ct.Name == "init-snapshots" {
			t.Require().Len(ct.Args, 2)
			t.Equal(DownloadMainnet+".sh", ct.Args[1])
		}
	}
	// init-snapshots
	t.Require().Equal(seen, 1)
	for k, v := range inf.ConfigMap.Data {
		fmt.Println(k, v)
	}
}

func (t *SuiCookbookTestSuite) TestSuiAllOptsEnabled() {
	p := suiMasterChartPath
	p.DirIn = "./sui/node/custom_sui"
	p.DirOut = "./sui/node/custom_sui"
	cfg := SuiConfigOpts{
		DownloadSnapshot:   true,
		WithIngress:        true,
		WithServiceMonitor: true,
		CloudProvider:      "gcp",
		Network:            mainnet,
	}
	sbDefs := t.generateFromConfigDriverBuilder(cfg)

	seenIngressCount := 0
	seenServiceMonitorCount := 0
	for _, sb := range sbDefs {
		if sb.Workload.Ingress != nil {
			p.FnOut = "ing-sui.yaml"
			seenIngressCount += 1
			b, err := json.Marshal(sb.Workload.Ingress)
			t.Require().Nil(err)
			err = p.WriteToFileOutPath(b)
			t.Require().Nil(err)
		}
		if sb.Workload.ServiceMonitor != nil {
			p.FnOut = "sm-sui.yaml"
			seenServiceMonitorCount += 1
			b, err := json.Marshal(sb.Workload.ServiceMonitor)
			t.Require().Nil(err)
			err = p.WriteToFileOutPath(b)
			t.Require().Nil(err)
		}
	}

	if cfg.WithIngress {
		t.Require().Equal(seenIngressCount, 1)
	}
	if cfg.WithServiceMonitor {
		t.Require().Equal(seenServiceMonitorCount, 1)
	}
	t.Require().Len(sbDefs, 3)
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	t.NotNil(inf.Ingress)
	t.NotNil(inf.ServiceMonitor)
	t.Nil(inf.Deployment)
	t.NotNil(inf.StatefulSet)
	t.NotNil(inf.Service)
	t.NotNil(inf.ConfigMap)

	t.NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates)
	t.Len(inf.StatefulSet.Spec.VolumeClaimTemplates, 1)
	t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests)
	t.Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage(), resource.MustParse(mainnetDiskSize))
	t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName)
	t.Require().Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName, gcp_nvme.GcpStorageClass)

	seen := 0
	for _, ct := range inf.StatefulSet.Spec.Template.Spec.Containers {
		if ct.Name == Sui {
			t.Equal(*ct.Resources.Requests.Cpu(), resource.MustParse(cpuCores))
			t.Equal(*ct.Resources.Requests.Memory(), resource.MustParse(memorySize))
			t.Equal(*ct.Resources.Limits.Cpu(), resource.MustParse(cpuCores))
			t.Equal(*ct.Resources.Limits.Memory(), resource.MustParse(memorySize))
			seen++
		}
	}
	for _, ct := range inf.StatefulSet.Spec.Template.Spec.InitContainers {
		if ct.Name == "init-snapshots" {
			t.Require().Len(ct.Args, 2)
			t.Equal(DownloadMainnet+".sh", ct.Args[1])
		}
	}
	// init-snapshots
	t.Require().Equal(seen, 1)
	for k, v := range inf.ConfigMap.Data {
		fmt.Println(k, v)
	}
}
