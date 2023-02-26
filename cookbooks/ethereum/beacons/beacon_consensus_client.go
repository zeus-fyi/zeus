package ethereum_beacon_cookbooks

import (
	"fmt"
	"time"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeployConsensusClientKnsReq set your own topologyID here after uploading a chart workload
var DeployConsensusClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
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
				consensusClient: {Container: v1Core.Container{
					Name:  consensusClient,
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
						"app":     consensusClient,
						"release": "kube-prometheus-stack",
					}},
				Spec: v1.ServiceMonitorSpec{
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/instance": consensusClient,
							"app.kubernetes.io/name":     consensusClient,
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
	ComponentBaseName: "consensusClient",
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
