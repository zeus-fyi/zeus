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

// DeployExecClientKnsReq set your own topologyID here after uploading a chart workload
var DeployExecClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
}

var ExecClientSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: BeaconExecClientChartPath,
	TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
		ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "cm-exec-client"},
			},
			SwapKeys: map[string]string{
				"start.sh": GethEphemeral + ".sh",
			},
		},
		StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
			ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
				execClient: {
					Container: v1Core.Container{
						Name:  execClient,
						Image: gethDockerImageCapella,
					},
				},
			},
			PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
				PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
					execClientDiskName: {
						ObjectMeta: metav1.ObjectMeta{Name: execClientDiskName},
						Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
							Requests: v1Core.ResourceList{"storage": resource.MustParse(execClientDiskSizeEphemeral)},
						}},
					},
				}},
		},
	}}

var ExecClientSkeletonBaseMonitoringConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: ServiceMonitorChartPath,
	TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
		ServiceMonitorDriver: &zeus_topology_config_drivers.ServiceMonitorDriver{
			ServiceMonitor: v1.ServiceMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name: "zeus-exec-client-monitor",
					Labels: map[string]string{
						"app":     execClient,
						"release": "kube-prometheus-stack",
					}},
				Spec: v1.ServiceMonitorSpec{
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/instance": execClient,
							"app.kubernetes.io/name":     execClient,
						},
					},
				},
			}},
	},
}

var ExecClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "gethHercules",
	ChartName:         "gethHercules",
	ChartDescription:  "gethHercules",
	Version:           fmt.Sprintf("gethHerculesv0.0.%d", time.Now().Unix()),
	ClusterClassName:  "ethereumBeacons",
	ComponentBaseName: "executionClient",
	SkeletonBaseName:  "gethHercules",
	Tag:               "latest",
}

var BeaconExecClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/exec_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "gethHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
