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
	devnet  = "devnet"

	// docker image references
	dockerImage        = "mysten/sui-node:stable"
	dockerImageTestnet = "mysten/sui-node:testnet"
	dockerImageDevnet  = "mysten/sui-node:devnet"

	hercules = "hercules"

	// mainnet workload compute resources
	cpuCores   = "16"
	memorySize = "128Gi"
	// mainnet workload disk sizes
	mainnetDiskSize = "4Ti"

	// testnet compute resources
	cpuCoresTestnet   = "7500m"
	memorySizeTestnet = "63Gi"
	// testnet workload disk sizes
	testnetDiskSize = "3Ti"
	devnetDiskSize  = "2Ti"

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

	DownloadTestnet       = "downloadTestnetNode"
	DownloadTestnetNodeDo = "downloadTestnetNodeDo"
	DownloadDevnetNodeDo  = "downloadDevnetNodeDo"

	DownloadTestnetNodeAws = "downloadTestnetNodeAws"
)

type SuiConfigOpts struct {
	DownloadSnapshot bool   `json:"downloadSnapshot"`
	Network          string `json:"network"`

	CloudProvider string `json:"cloudProvider"`
	WithLocalNvme bool   `json:"withLocalNvme"`

	WithIngress          bool `json:"withIngress"`
	WithServiceMonitor   bool `json:"withServiceMonitor"`
	WithArchivalFallback bool `json:"withArchivalFallback"`
	WithHercules         bool `json:"withHercules"`
}

func GetSuiClientNetworkConfigBase(cfg SuiConfigOpts) zeus_cluster_config_drivers.ComponentBaseDefinition {
	downloadStartup := ""
	diskSize := mainnetDiskSize
	cpuSize := cpuCores
	memSize := memorySize
	dockerImageSui := dockerImage
	entryPointScript := "entrypoint.sh"
	switch cfg.Network {
	case mainnet:
		// todo, add workload type conditional here
		cpuSize = cpuCores
		memSize = memorySize
		diskSize = mainnetDiskSize
		downloadStartup = DownloadMainnet
		dockerImageSui = dockerImage
	case testnet:
		diskSize = testnetDiskSize
		cpuSize = cpuCoresTestnet
		memSize = memorySizeTestnet
		downloadStartup = DownloadTestnet
		dockerImageSui = dockerImageTestnet
	case devnet:
		diskSize = devnetDiskSize
		cpuSize = cpuCoresTestnet
		memSize = memorySizeTestnet
		downloadStartup = DownloadTestnet
		dockerImageSui = dockerImageDevnet
		entryPointScript = "noFallBackEntrypoint.sh"
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
		case devnet:
			downloadStartup = DownloadDevnetNodeDo
		}
	}
	if !cfg.WithLocalNvme {
		dataDir = "/data"
	}
	var storageClassName *string
	if cfg.WithLocalNvme {
		storageClassName = aws.String(zeus_nvme.ConfigureCloudProviderStorageClass(cfg.CloudProvider))
	}
	if !cfg.WithArchivalFallback {
		entryPointScript = "noFallBackEntrypoint.sh"
	}
	var envAddOns []v1Core.EnvVar
	if cfg.WithArchivalFallback {
		s3AccessKey := zeus_topology_config_drivers.MakeSecretEnvVar("AWS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID", "aws-credentials")
		s3SecretKey := zeus_topology_config_drivers.MakeSecretEnvVar("AWS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "aws-credentials")
		envAddOns = []v1Core.EnvVar{s3AccessKey, s3SecretKey}
	}

	sbCfg := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: suiMasterChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
				ConfigMap: v1Core.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{Name: suiConfigMap},
					Data: map[string]string{
						"fullnode.yaml": OverrideNodeConfigDataDir(dataDir, cfg),
					},
				},
			},
			ServiceDriver: sd,
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					Sui: {
						Container: v1Core.Container{
							Name:      Sui,
							Image:     dockerImageSui,
							Resources: zeus_topology_config_drivers.CreateComputeResourceRequirementsLimit(cpuSize, memSize),
							VolumeMounts: []v1Core.VolumeMount{{
								Name:      suiDiskName,
								MountPath: dataDir,
							}},
							Command: []string{fmt.Sprintf("/scripts/%s", entryPointScript)},
						},
						AppendEnvVars: envAddOns,
					},
					"hercules": {
						Container: v1Core.Container{
							Name: "hercules",
							VolumeMounts: []v1Core.VolumeMount{{
								Name:      suiDiskName,
								MountPath: dataDir,
							}},
						},
						IsDeleteContainer: !cfg.WithHercules,
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

func OverrideNodeConfigDataDir(dataDir string, cfg SuiConfigOpts) string {
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
				"genesis-file-location": dataDir + "/genesis.blob",
			}
		}
	}
	network := cfg.Network
	if network == mainnet || network == testnet {
		p2pCfg := GetP2PTable(network)
		if p2pCfg != nil {
			m["p2p-config"] = p2pCfg
		}
		if cfg.WithArchivalFallback {
			fallbackCfg := GetArchiveFallback(network)
			if fallbackCfg != nil {
				m["state-archive-read-config"] = fallbackCfg
			}
		}
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

func GetArchiveFallback(network string) interface{} {
	p := suiMasterChartPath
	switch network {
	case "mainnet", "testnet":
		p.FnIn = "archival-fallback.yaml"
	default:
		return nil
	}
	p.DirIn = "./sui/node/sui_config"
	fip := p.FileInPath()
	fallbackCfg, err := yaml_fileio.ReadYamlConfig(fip)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = yaml.Unmarshal(fallbackCfg, &m)
	if err != nil {
		panic(err)
	}

	if stateCfgList, ok := m["state-archive-read-config"].([]interface{}); ok {
		for _, cfg := range stateCfgList {
			if cfgMap, ok2 := cfg.(map[string]interface{}); ok2 {
				if objStoreCfg, ok3 := cfgMap["object-store-config"].(map[string]interface{}); ok3 {
					objStoreCfg["aws-access-key-id"] = "<AWS_ACCESS_KEY_ID>"
					objStoreCfg["aws-secret-access-key"] = "<AWS_SECRET_ACCESS_KEY>"
					// Setting the bucket name based on the network
					bn := fmt.Sprintf("mysten-%s-archives", network)
					objStoreCfg["bucket"] = bn
				}
			}
		}
	}

	return m["state-archive-read-config"]
}
