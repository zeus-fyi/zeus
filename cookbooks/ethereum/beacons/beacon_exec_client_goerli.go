package ethereum_beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const cmExecClient = "cm-exec-client"

var ExecClientGoerliSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: BeaconExecClientChartPath,
	TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
		ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: cmExecClient},
			},
			SwapKeys: map[string]string{
				"start.sh": GethGoerli + ".sh",
			},
		},
		StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
			ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
				zeusExecClient: {
					Container: v1Core.Container{
						Name:  zeusExecClient,
						Image: gethDockerImage,
					},
				},
			},
			PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
				PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
					execClientDiskName: {
						ObjectMeta: metav1.ObjectMeta{Name: execClientDiskName},
						Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
							Requests: v1Core.ResourceList{"storage": resource.MustParse(execClientDiskSizeGoerli)},
						}},
					},
				}},
		},
	}}
