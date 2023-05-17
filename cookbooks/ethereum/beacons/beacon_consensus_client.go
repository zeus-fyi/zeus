package ethereum_beacon_cookbooks

import (
	"fmt"
	"time"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeployConsensusClientKnsReq set your own topologyID here after uploading a chart workload
var DeployConsensusClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
}

const (
	consensusClientEphemeralRequestCPU    = "2.5"
	consensusClientEphemeralRequestMemory = "4Gi"
	consensusClientMainnetCPU             = "7"
	consensusClientMainnetCPUMemory       = "13Gi"
)

func GetConsensusClientNetworkConfig(consensusClient, network string, choreographySecretsExist bool) zeus_cluster_config_drivers.ComponentBaseDefinition {
	dockerImage := ""
	cmConfig := ""
	diskSize := consensusStorageDiskSizeMainnet
	downloadStartup := download + ".sh"
	herculesStartup := hercules + ".sh"
	cpuSize := consensusClientMainnetCPU
	memSize := consensusClientMainnetCPUMemory
	var ports []v1Core.ContainerPort
	var svcDriver *zeus_topology_config_drivers.ServiceDriver
	switch consensusClient {
	case client_consts.Lodestar:
		svcDriver = &zeus_topology_config_drivers.ServiceDriver{
			Service: v1Core.Service{
				Spec: v1Core.ServiceSpec{
					Ports: []v1Core.ServicePort{
						{
							Name:       "hercules",
							Protocol:   "TCP",
							Port:       9003,
							TargetPort: intstr.IntOrString{Type: intstr.String, StrVal: "hercules"},
						},
						{
							Name:       "p2p-tcp",
							Protocol:   "TCP",
							Port:       9000,
							TargetPort: intstr.IntOrString{Type: intstr.String, StrVal: "p2p-tcp"},
						},
						{
							Name:       "p2p-udp",
							Protocol:   "UDP",
							Port:       9000,
							TargetPort: intstr.IntOrString{Type: intstr.String, StrVal: "p2p-udp"},
						},
						{
							Name:       "http-api",
							Protocol:   "TCP",
							Port:       int32(lodestarRestPort),
							TargetPort: intstr.IntOrString{Type: intstr.String, StrVal: "http-api"},
						},
					},
				},
			},
		}
		ports = []v1Core.ContainerPort{
			{
				Name:          "p2p-tcp",
				ContainerPort: 9000,
				Protocol:      "TCP",
			},
			{
				Name:          "p2p-udp",
				ContainerPort: 9000,
				Protocol:      "UDP",
			},
			{
				Name:          "http-api",
				ContainerPort: int32(lodestarRestPort),
				Protocol:      "TCP",
			},
		}

		switch network {
		case hestia_req_types.Ephemery, "ephemeral":
			dockerImage = lodestarDockerImage
			cmConfig = LodestarEphemeral
			diskSize = consensusStorageDiskSizeEphemeral
			downloadStartup = "downloadLodestarEphemeral.sh"
			herculesStartup = "herculesLodestarEphemeral.sh"
			cpuSize = consensusClientEphemeralRequestCPU
			memSize = consensusClientEphemeralRequestMemory
		}
	case client_consts.Lighthouse:
		switch network {
		case hestia_req_types.Ephemery, "ephemeral":
			dockerImage = lighthouseDockerImage
			cmConfig = LighthouseEphemeral
			diskSize = consensusStorageDiskSizeEphemeral
			downloadStartup = "downloadLighthouseEphemeral.sh"
			herculesStartup = "herculesLighthouseEphemeral.sh"
			cpuSize = consensusClientEphemeralRequestCPU
			memSize = consensusClientEphemeralRequestMemory
		case hestia_req_types.Mainnet:
			dockerImage = lighthouseDockerImage
		}

	}
	rr := v1Core.ResourceRequirements{
		Requests: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
	}
	cp := filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/beacons/infra/consensus_client",
		DirOut:      "./ethereum/outputs",
		FnIn:        consensusClient + "Hercules", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}

	initContDriver := zeus_topology_config_drivers.ContainerDriver{
		AppendEnvVars: []v1Core.EnvVar{BearerTokenSecretFromChoreography},
	}
	if choreographySecretsExist {
		initContDriver = zeus_topology_config_drivers.ContainerDriver{
			IsInitContainer: true,
			AppendEnvVars:   []v1Core.EnvVar{BearerTokenSecretFromChoreography},
		}
	}
	sbDef := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: cp,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
				ConfigMap: v1Core.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{Name: "cm-consensus-client"},
				},
				SwapKeys: map[string]string{
					"start.sh":       cmConfig + ".sh",
					download + ".sh": downloadStartup,
					hercules + ".sh": herculesStartup,
				},
			},
			ServiceDriver: svcDriver,
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					initSnapshots: initContDriver,
					zeusConsensusClient: {Container: v1Core.Container{
						Name:      consensusClient,
						Image:     dockerImage,
						Ports:     ports,
						Resources: rr,
					}},
				},
				PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
					PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
						consensusStorageDiskName: {
							ObjectMeta: metav1.ObjectMeta{Name: consensusStorageDiskName},
							Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
								Requests: v1Core.ResourceList{"storage": resource.MustParse(diskSize)},
							}},
						},
					}},
			},
		}}
	if choreographySecretsExist {
	}
	compBase := zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"consensusClient" + "Hercules": sbDef,
		},
	}
	return compBase
}

var ConsensusClientSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: BeaconConsensusClientChartPath,
	TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
		ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "cm-consensus-client"},
			},
			SwapKeys: map[string]string{
				"start.sh": LighthouseEphemeral + ".sh",
			},
		},
		StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
			ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
				zeusConsensusClient: {Container: v1Core.Container{
					Name:  zeusConsensusClient,
					Image: lighthouseDockerImageCapella,
				}},
			},
			PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
				PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
					consensusStorageDiskName: {
						ObjectMeta: metav1.ObjectMeta{Name: consensusStorageDiskName},
						Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
							Requests: v1Core.ResourceList{"storage": resource.MustParse(consensusStorageDiskSizeEphemeral)},
						}},
					},
				}},
		},
	}}

var ConsensusClientSkeletonBaseMonitoringConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: ServiceMonitorChartPath,
	TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
		ServiceMonitorDriver: &zeus_topology_config_drivers.ServiceMonitorDriver{
			ServiceMonitor: v1.ServiceMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name: "zeus-consensus-client-monitor",
					Labels: map[string]string{
						"app":     zeusConsensusClient,
						"release": "kube-prometheus-stack",
					}},
				Spec: v1.ServiceMonitorSpec{
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/instance": zeusConsensusClient,
							"app.kubernetes.io/name":     zeusConsensusClient,
						},
					},
				},
			}},
	},
}

var ConsensusClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "lighthouseHercules",
	ChartName:         "lighthouseHercules",
	ChartDescription:  "lighthouseHercules",
	Version:           fmt.Sprintf("lighthouseHerculesv0.0.%d", time.Now().Unix()),
	ClusterClassName:  "ethereumBeacons",
	ComponentBaseName: "zeusConsensusClient",
	SkeletonBaseName:  "lighthouseHercules",
	Tag:               "latest",
}

var BeaconConsensusClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/consensus_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "lighthouseHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
