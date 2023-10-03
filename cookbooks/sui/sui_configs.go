package sui_cookbooks

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ghodss/yaml"
	yaml_fileio "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/yaml"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme"
	aws_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/aws"
	do_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/do"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	v1Core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// networks
	mainnet = "mainnet"
	testnet = "testnet"

	// docker image references
	dockerImage = "mysten/sui-node:stable"
	hercules    = "hercules"

	// mainnet workload compute resources
	cpuCores   = "16"
	memorySize = "128Gi"
	// mainnet workload disk sizes
	mainnetDiskSize = "4Ti"

	// testnet compute resources
	cpuCoresTestnet   = "7500m"
	memorySizeTestnet = "63Gi"
	// testnet workload disk sizes
	testnetDiskSize = "2Ti"

	// workload label, name, or k8s references
	suiDiskName  = "sui-client-storage"
	suiConfigMap = "cm-sui"

	// workload type
	suiNodeConfig      = "full"
	suiValidatorConfig = "validator"

	SuiRpcPortName = "http-rpc"

	DownloadMainnet        = "downloadMainnetNode"
	DownloadMainnetNodeDo  = "downloadMainnetNodeDo"
	DownloadMainnetNodeAws = "downloadMainnetNodeAws"

	DownloadTestnet        = "downloadTestnetNode"
	DownloadTestnetNodeDo  = "downloadTestnetNodeDo"
	DownloadTestnetNodeAws = "downloadTestnetNodeAws"
)

type SuiConfigOpts struct {
	DownloadSnapshot bool   `json:"downloadSnapshot"`
	Network          string `json:"network"`

	WithIngress        bool   `json:"withIngress"`
	WithServiceMonitor bool   `json:"withServiceMonitor"`
	CloudProvider      string `json:"cloudProvider"`
	WithLocalNvme      bool   `json:"withLocalNvme"`
}

func GetSuiClientNetworkConfigBase(cfg SuiConfigOpts) zeus_cluster_config_drivers.ComponentBaseDefinition {
	downloadStartup := ""
	diskSize := mainnetDiskSize
	cpuSize := cpuCores
	memSize := memorySize
	switch cfg.Network {
	case mainnet:
		// todo, add workload type conditional here
		cpuSize = cpuCores
		memSize = memorySize
		diskSize = mainnetDiskSize
		downloadStartup = DownloadMainnet
	case testnet:
		diskSize = testnetDiskSize
		cpuSize = cpuCoresTestnet
		memSize = memorySizeTestnet
		downloadStartup = DownloadTestnet
	}

	sd := &zeus_topology_config_drivers.ServiceDriver{}
	if cfg.WithIngress {
		sd.AddNginxTargetPort("nginx", SuiRpcPortName)
	}

	dataDir := "/data"
	switch cfg.CloudProvider {
	case "aws":
		dataDir = aws_nvme.AwsNvmePath
		switch cfg.Network {
		case mainnet:
			downloadStartup = DownloadMainnetNodeAws
		case testnet:
			downloadStartup = DownloadTestnetNodeAws
		}
	case "gcp":
		// todo, add gcp nvme path
	case "do":
		dataDir = do_nvme.DoNvmePath
		switch cfg.Network {
		case mainnet:
			downloadStartup = DownloadMainnetNodeDo
		case testnet:
			downloadStartup = DownloadTestnetNodeDo
		}
	}
	if !cfg.WithLocalNvme {
		dataDir = "/data"
	}
	var storageClassName *string
	if cfg.WithLocalNvme {
		storageClassName = aws.String(zeus_nvme.ConfigureCloudProviderStorageClass(cfg.CloudProvider))
	}
	sbCfg := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: suiMasterChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
				ConfigMap: v1Core.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{Name: suiConfigMap},
					Data: map[string]string{
						"fullnode.yaml": OverrideNodeConfigDataDir(dataDir, cfg.Network),
					},
				},
			},
			ServiceDriver: sd,
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					Sui: {
						Container: v1Core.Container{
							Name:      Sui,
							Image:     dockerImage,
							Resources: zeus_topology_config_drivers.CreateComputeResourceRequirementsLimit(cpuSize, memSize),
							VolumeMounts: []v1Core.VolumeMount{{
								Name:      suiDiskName,
								MountPath: dataDir,
							}},
						},
					},
					"hercules": {
						Container: v1Core.Container{
							Name: "hercules",
							VolumeMounts: []v1Core.VolumeMount{{
								Name:      suiDiskName,
								MountPath: dataDir,
							}},
						},
					},
					"init-snapshots": {
						Container: v1Core.Container{
							Name: "init-snapshots",
							Args: []string{"-c", fmt.Sprintf("/scripts/%s.sh", downloadStartup)},
							VolumeMounts: []v1Core.VolumeMount{{
								Name:      suiDiskName,
								MountPath: dataDir,
							}},
						},
						IsInitContainer:   true,
						IsDeleteContainer: !cfg.DownloadSnapshot,
					},
					"init-chown-data": {
						Container: v1Core.Container{
							Name:    "init-chown-data",
							Command: []string{"chown", "-R", "10001:10001", dataDir},
							VolumeMounts: []v1Core.VolumeMount{{
								Name:      suiDiskName,
								MountPath: dataDir,
							}},
						},
						IsInitContainer: true,
					},
				},
				PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
					PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
						suiDiskName: {
							ObjectMeta: metav1.ObjectMeta{Name: suiDiskName},
							Spec: v1Core.PersistentVolumeClaimSpec{
								Resources:        zeus_topology_config_drivers.CreateDiskResourceRequirementsLimit(diskSize),
								StorageClassName: storageClassName,
							},
						},
					}},
			},
		}}
	suiCompBase := zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			Sui: sbCfg,
		},
	}
	return suiCompBase
}

func OverrideNodeConfigDataDir(dataDir, network string) string {
	p := suiMasterChartPath
	p.FnIn = "fullnode.yaml"
	p.DirIn = "./sui/node/sui_config"
	fip := p.FileInPath()
	nodeCfg, err := yaml_fileio.ReadYamlConfig(fip)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(nodeCfg, &m)
	if err != nil {
		panic(err)
	}
	for k, _ := range m {
		if k == "db-path" {
			m[k] = dataDir
		}
		if k == "genesis" {
			m[k] = map[string]interface{}{
				"genesis-file-location": "genesis.blob",
			}
		}
	}
	p2pCfg := GetP2PTable(network)
	if p2pCfg != nil {
		m["p2p-config"] = p2pCfg
	}
	b, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetP2PTable(network string) interface{} {
	p := suiMasterChartPath
	switch network {
	case mainnet:
		p.FnIn = "p2p-mainnet.yaml"
	case testnet:
		p.FnIn = "p2p-testnet.yaml"
	default:
		return nil
	}
	p.DirIn = "./sui/node/sui_config"
	fip := p.FileInPath()
	p2pCfg, err := yaml_fileio.ReadYamlConfig(fip)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = yaml.Unmarshal(p2pCfg, &m)
	if err != nil {
		panic(err)
	}
	return m["p2p-config"]
}
