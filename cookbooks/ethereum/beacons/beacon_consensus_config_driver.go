package ethereum_beacon_cookbooks

import (
	"errors"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	zeusConsensusClient               = "zeus-consensus-client"
	consensusStorageDiskName          = "consensus-client-storage"
	consensusStorageDiskSizeEphemeral = "2Gi"
	consensusStorageDiskSizeGoerli    = "100Gi"
	consensusStorageDiskSizeMainnet   = "300Gi"

	LighthouseMainnet           = "lighthouse"
	LodestarEphemeral           = "lodestar-ephemeral"
	LighthouseEphemeral         = "lighthouse-ephemeral"
	downloadLighthouseEphemeral = "downloadLighthouseEphemeral"

	lighthouseDockerImage        = "sigp/lighthouse:v4.5.0"
	lighthouseDockerImageCapella = "sigp/lighthouse:capella"

	LodestarGoerli      = "lodestar-goerli"
	lodestarDockerImage = "chainsafe/lodestar:v1.12.1"
)

const (
	consensusClientEphemeralRequestCPU    = "2.5"
	consensusClientEphemeralRequestMemory = "4Gi"
	consensusClientMainnetCPU             = "6"
	consensusClientMainnetCPUMemory       = "13Gi"
)

func GetConsensusClientNetworkConfig(beaconConfig BeaconConfig) zeus_cluster_config_drivers.ComponentBaseDefinition {
	dockerImage := ""
	cmConfig := ""
	diskSize := consensusStorageDiskSizeMainnet
	downloadStartup := download + ".sh"
	herculesStartup := hercules + ".sh"
	cpuSize := consensusClientMainnetCPU
	memSize := consensusClientMainnetCPUMemory
	var ports []v1Core.ContainerPort
	var svcDriver *config_overrides.ServiceDriver
	switch beaconConfig.ConsensusClient {
	case client_consts.Lodestar:
		svcDriver = &config_overrides.ServiceDriver{
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

		switch beaconConfig.Network {
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
		switch beaconConfig.Network {
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
			cmConfig = LighthouseMainnet
		}

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
	cp := filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/beacons/infra/consensus_client",
		DirOut:      "./ethereum/outputs",
		FnIn:        beaconConfig.ConsensusClient, // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}

	initContDriver := config_overrides.ContainerDriver{
		IsInitContainer: true,
	}
	if beaconConfig.WithChoreography {
		initContDriver.AppendEnvVars = []v1Core.EnvVar{BearerTokenSecretFromChoreography}
	}

	sbDef := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: cp,
		TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
			ConfigMapDriver: &config_overrides.ConfigMapDriver{
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
			StatefulSetDriver: &config_overrides.StatefulSetDriver{
				ContainerDrivers: map[string]config_overrides.ContainerDriver{
					initSnapshots: initContDriver,
					zeusConsensusClient: {Container: v1Core.Container{
						Name:      beaconConfig.ConsensusClient,
						Image:     dockerImage,
						Ports:     ports,
						Resources: rr,
					}},
				},
				PVCDriver: &config_overrides.PersistentVolumeClaimsConfigDriver{
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
	compBase := zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"consensus-client": sbDef,
		},
	}
	return compBase
}

var ConsensusClientSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: BeaconConsensusClientChartPath,
	TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
		ConfigMapDriver: &config_overrides.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "cm-consensus-client"},
			},
			SwapKeys: map[string]string{
				"start.sh": LighthouseEphemeral + ".sh",
			},
		},
		StatefulSetDriver: &config_overrides.StatefulSetDriver{
			ContainerDrivers: map[string]config_overrides.ContainerDriver{
				zeusConsensusClient: {Container: v1Core.Container{
					Name:  zeusConsensusClient,
					Image: lighthouseDockerImage,
				}},
			},
			PVCDriver: &config_overrides.PersistentVolumeClaimsConfigDriver{
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
	TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
		ServiceMonitorDriver: &config_overrides.ServiceMonitorDriver{
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

func EphemeralConsensusClientLighthouseConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.ConfigMap != nil {
		m := make(map[string]string)
		m = inf.ConfigMap.Data
		vSrc, ok := inf.ConfigMap.Data[LighthouseEphemeral+".sh"]
		if ok {
			m[start+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[downloadLighthouseEphemeral+".sh"]
		if ok {
			m[download+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[herculesEphemeral+".sh"]
		if ok {
			m[hercules+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
	}
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == zeusConsensusClient {
				inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = lighthouseDockerImageCapella
			}
		}
		for i, v := range inf.StatefulSet.Spec.VolumeClaimTemplates {
			if v.Name == consensusStorageDiskName {
				q, err := resource.ParseQuantity(consensusStorageDiskSizeEphemeral)
				if err != nil {
					panic(err)
				}
				for j, _ := range inf.StatefulSet.Spec.VolumeClaimTemplates[i].Spec.Resources.Requests {
					tmp := inf.StatefulSet.Spec.VolumeClaimTemplates[i].Spec.Resources.Requests[j]
					tmp.Set(q.Value())
					inf.StatefulSet.Spec.VolumeClaimTemplates[i].Spec.Resources.Requests[j] = tmp
				}
			}
		}
	}
}
