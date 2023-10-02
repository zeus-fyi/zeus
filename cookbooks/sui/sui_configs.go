package sui_cookbooks

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme"
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

	DownloadMainnet = "downloadMainnetNode"
	DownloadTestnet = "downloadTestnetNode"
	NoDownload      = "noDownload"
)

type SuiConfigOpts struct {
	DownloadSnapshot bool   `json:"downloadSnapshot"`
	Network          string `json:"network"`

	WithIngress        bool   `json:"withIngress"`
	WithServiceMonitor bool   `json:"withServiceMonitor"`
	CloudProvider      string `json:"cloudProvider"`
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
	if !cfg.DownloadSnapshot {
		downloadStartup = NoDownload
	}
	sd := &zeus_topology_config_drivers.ServiceDriver{}
	if cfg.WithIngress {
		sd.AddNginxTargetPort("nginx", SuiRpcPortName)
	}
	sbCfg := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: suiMasterChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			ServiceDriver: sd,
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					Sui: {
						Container: v1Core.Container{
							Name:      Sui,
							Image:     dockerImage,
							Resources: zeus_topology_config_drivers.CreateComputeResourceRequirementsLimit(cpuSize, memSize),
						},
					},
					"init-snapshots": {
						Container: v1Core.Container{
							Name: "init-snapshots",
							Args: []string{"-c", downloadStartup + ".sh"},
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
								StorageClassName: aws.String(zeus_nvme.ConfigureCloudProviderStorageClass(cfg.CloudProvider)),
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
