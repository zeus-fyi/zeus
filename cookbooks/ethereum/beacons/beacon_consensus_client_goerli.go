package ethereum_beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var lodestarRestPort = 9596

var ConsensusClientGoerliSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: BeaconConsensusClientChartPath,
	TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
		ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "cm-consensus-client"},
			},
			SwapKeys: map[string]string{
				"start.sh": LodestarGoerli + ".sh",
			},
		},
		ServiceDriver: &zeus_topology_config_drivers.ServiceDriver{
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
		},
		StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
			ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
				zeusConsensusClient: {Container: v1Core.Container{
					Name:  zeusConsensusClient,
					Image: lodestarDockerImage,
					Ports: []v1Core.ContainerPort{
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
					},
				}},
			},
			PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
				PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
					consensusStorageDiskName: {
						ObjectMeta: metav1.ObjectMeta{Name: consensusStorageDiskName},
						Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
							Requests: v1Core.ResourceList{"storage": resource.MustParse(consensusStorageDiskSizeGoerli)},
						}},
					},
				}},
		},
	}}
