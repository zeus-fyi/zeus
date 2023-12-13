package ethereum_beacon_cookbooks

import (
	"fmt"
	"time"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
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
	TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
		ConfigMapDriver: &config_overrides.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "cm-exec-client"},
			},
			SwapKeys: map[string]string{
				"start.sh": GethEphemeral + ".sh",
			},
		},
		StatefulSetDriver: &config_overrides.StatefulSetDriver{
			ContainerDrivers: map[string]config_overrides.ContainerDriver{
				zeusExecClient: {
					Container: v1Core.Container{
						Name:  zeusExecClient,
						Image: gethDockerImage,
					},
				},
			},
			PVCDriver: &config_overrides.PersistentVolumeClaimsConfigDriver{
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
	TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
		ServiceMonitorDriver: &config_overrides.ServiceMonitorDriver{
			ServiceMonitor: v1.ServiceMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name: "zeus-exec-client-monitor",
					Labels: map[string]string{
						"app":     zeusExecClient,
						"release": "kube-prometheus-stack",
					}},
				Spec: v1.ServiceMonitorSpec{
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/instance": zeusExecClient,
							"app.kubernetes.io/name":     zeusExecClient,
						},
					},
				},
			}},
	},
}

var ExecClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "geth-hercules",
	ChartName:         "geth-hercules",
	ChartDescription:  "geth-hercules",
	Version:           fmt.Sprintf("gethHerculesv0.0.%d", time.Now().Unix()),
	ClusterClassName:  "ethereum-beacons",
	ComponentBaseName: "execution-client",
	SkeletonBaseName:  "geth-hercules",
	Tag:               "latest",
}

var BeaconExecClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/exec_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "geth-hercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
