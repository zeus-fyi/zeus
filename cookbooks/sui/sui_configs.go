package sui_cookbooks

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
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
	DownloadSnapshot bool
	WithIngress      bool
	CloudProvider    string
	Network          string
}

func GetSuiClientNetworkConfigBase(cfg SuiConfigOpts) zeus_cluster_config_drivers.ComponentBaseDefinition {
	cmConfig := ""
	downloadStartup := ""
	diskSize := mainnetDiskSize
	herculesStartup := hercules + ".sh"
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
	rr := v1Core.ResourceRequirements{
		Requests: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
		Limits: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
	}
	sd := &zeus_topology_config_drivers.ServiceDriver{}
	if cfg.WithIngress {
		sd.AddNginxTargetPort("nginx", SuiRpcPortName)
	}
	sbCfg := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: suiMasterChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
				ConfigMap: v1Core.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{Name: suiConfigMap},
				},
				SwapKeys: map[string]string{
					"start.sh":              cmConfig + ".sh",
					hercules + ".sh":        herculesStartup,
					downloadStartup + ".sh": downloadStartup,
				},
			},
			ServiceDriver: sd,
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					Sui: {
						Container: v1Core.Container{
							Name:      Sui,
							Image:     dockerImage,
							Resources: rr,
						},
					},
				},
				PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
					PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
						suiDiskName: {
							ObjectMeta: metav1.ObjectMeta{Name: suiDiskName},
							Spec: v1Core.PersistentVolumeClaimSpec{
								Resources: v1Core.ResourceRequirements{
									Requests: v1Core.ResourceList{"storage": resource.MustParse(diskSize)},
								},
								StorageClassName: aws.String(ConfigureCloudProviderStorageClass(cfg.CloudProvider)),
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
