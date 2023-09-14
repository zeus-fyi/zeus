package sui_cookbooks

import (
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	dockerImage = "mysten/sui-node:stable"
	hercules    = "hercules"

	suiDiskName     = "sui-client-storage"
	mainnetDiskSize = "6Ti"
	testnetDiskSize = "2Ti"
	cpuCores        = "16"
	memorySize      = "128Gi"

	suiConfigMap = "cm-sui"

	suiNodeConfig      = "full"
	suiValidatorConfig = "validator"
)

func GetSuiClientNetworkConfigBase(workloadType, network string) zeus_cluster_config_drivers.ComponentBaseDefinition {
	cmConfig := ""
	downloadStartup := ""
	diskSize := mainnetDiskSize
	herculesStartup := hercules + ".sh"
	cpuSize := cpuCores
	memSize := memorySize
	switch network {
	case "mainnet":
		// todo, add workload type conditional here
		downloadStartup = "downloadMainnetNode"
	case "testnet":
		diskSize = testnetDiskSize
		downloadStartup = "downloadTestnetNode"
	}
	rr := v1Core.ResourceRequirements{
		Requests: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
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
							Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
								Requests: v1Core.ResourceList{"storage": resource.MustParse(diskSize)},
							}},
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
