package ethereum_beacon_cookbooks

import (
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/zk8s_templates"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	v1Core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const cmExecClient = "cm-exec-client"

var ExecClientGoerliSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: BeaconExecClientChartPath,
	TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
		ConfigMapDriver: &config_overrides.ConfigMapDriver{
			ConfigMap: v1Core.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: cmExecClient},
			},
			SwapKeys: map[string]string{
				"start.sh": GethGoerli + ".sh",
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
					execClientDiskName: zk8s_templates.GetPvcTemplate(pvcTempGoeExecClient),
				}},
		},
	}}

var pvcTempGoeExecClient = zk8s_templates.PVCTemplate{
	Name:               execClientDiskName,
	StorageSizeRequest: execClientDiskSizeGoerli,
}
