package avax_node_cookbooks

import (
	"context"

	choreography_cookbooks "github.com/zeus-fyi/zeus/cookbooks/microservices/choreography"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	avaxDockerImage     = "avaplatform/avalanchego:v1.9.16"
	avaxFujiDockerImage = "avaplatform/avalanchego:v1.10.0-fuji"
	avaxClient          = "zeus-avax-client"

	avaxDiskName = "avax-client-storage"

	avaxDiskSize   = "2Ti"
	avaxCPURequest = "8"
	avaxRAMRequest = "12Gi"

	avaxFujiDiskSize   = "100Gi"
	avaxFujiCPURequest = "6"
	avaxFujiRAMRequest = "10Gi"
)

var (
	AvaxNodeClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "avaxNode",
		CloudCtxNs:       AvaxNodeCloudCtxNs,
		ComponentBases:   AvaxNodeComponentBases,
	}
	AvaxNodeCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "avax", // set with your own namespace
		Env:           "production",
	}
	AvaxNodeComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"avaxIngress":              IngressComponentBase,
		"avaxClients":              AvaxNodeComponentBase,
		"serviceMonitorAvaxClient": AvaxNodeMonitoringComponentBase,
		"choreography":             choreography_cookbooks.ChoreographyComponentBase,
	}
	AvaxNodeComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"avax": AvaxClientSkeletonBaseConfig,
		},
	}
	AvaxClientSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: AvaxClientChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					avaxClient: {Container: v1Core.Container{
						Name:  avaxClient,
						Image: avaxDockerImage,
						Resources: v1Core.ResourceRequirements{
							Limits: v1Core.ResourceList{
								"cpu":    resource.MustParse(avaxCPURequest),
								"memory": resource.MustParse(avaxRAMRequest),
							},
							Requests: v1Core.ResourceList{
								"cpu":    resource.MustParse(avaxCPURequest),
								"memory": resource.MustParse(avaxRAMRequest),
							},
						},
					}},
				},
				PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
					PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
						avaxDiskName: {
							ObjectMeta: metav1.ObjectMeta{Name: avaxDiskName},
							Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
								Requests: v1Core.ResourceList{"storage": resource.MustParse(avaxDiskSize)},
							}},
						},
					}},
			},
		},
	}
	AvaxNodeMonitoringComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"serviceMonitorAvaxClient": AvaxClientMonitorSkeletonBaseConfig,
		},
	}
	AvaxClientMonitorSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: ServiceMonitorChartPath,
	}
	ServiceMonitorChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./avax/node/servicemonitor",
		DirOut:      "./avax/node/processed_servicemonitor",
		FnIn:        "servicemonitor", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	IngressComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"avaxIngress": AvaxIngressSkeletonBaseConfig,
		},
	}
	AvaxClientChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./avax/node/infra",
		DirOut:      "./avax/outputs",
		FnIn:        "avax", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	AvaxIngressSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: IngressChartPath,
	}
	IngressChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./avax/node/ingress",
		DirOut:      "./avax/node/processed_avax_ingress",
		FnIn:        "avaxIngress", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)

func ConfigureAvaxNodeClusterBase(ctx context.Context, network string) zeus_cluster_config_drivers.ComponentBaseDefinition {
	dockerImage := avaxDockerImage
	cpuRequest := avaxCPURequest
	ramRequest := avaxRAMRequest
	diskSize := avaxDiskSize
	startCmd := []string{"-c", "/scripts/start.sh"}
	if network == "fuji" {
		dockerImage = avaxFujiDockerImage
		cpuRequest = avaxFujiCPURequest
		ramRequest = avaxFujiRAMRequest
		diskSize = avaxFujiDiskSize
		startCmd = []string{"-c", "/scripts/startFuji.sh"}
	}
	sb := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: AvaxClientChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					avaxClient: {Container: v1Core.Container{
						Name:  avaxClient,
						Image: dockerImage,
						Args:  startCmd,
						Resources: v1Core.ResourceRequirements{
							Limits: v1Core.ResourceList{
								"cpu":    resource.MustParse(cpuRequest),
								"memory": resource.MustParse(ramRequest),
							},
							Requests: v1Core.ResourceList{
								"cpu":    resource.MustParse(cpuRequest),
								"memory": resource.MustParse(ramRequest),
							},
						},
					}},
				},
				PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
					PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
						avaxDiskName: {
							ObjectMeta: metav1.ObjectMeta{Name: avaxDiskName},
							Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
								Requests: v1Core.ResourceList{"storage": resource.MustParse(diskSize)},
							}},
						},
					}},
			},
		},
	}
	return zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"avax": sb,
		},
	}
}
