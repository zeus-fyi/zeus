package ethereum_beacon_cookbooks

import (
	"errors"

	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	zeusExecClient              = "zeus-exec-client"
	execClientDiskName          = "exec-client-storage"
	execClientDiskSizeEphemeral = "10Gi"
	execClientDiskSizeGoerli    = "500Gi"

	hercules          = "hercules"
	herculesEphemeral = "herculesEphemeral"
	GethEphemeral     = "gethEphemeral"
	GethGoerli        = "gethGoerli"

	downloadGethEphemeral = "downloadGethEphemeral"
	gethDockerImage       = "ethereum/client-go:v1.11.6"

	gethDockerImageCapella = "ethpandaops/geth:master"
)

func GetExecClientNetworkConfig(execClient, network string) zeus_cluster_config_drivers.ComponentBaseDefinition {
	dockerImage := ""
	cmConfig := ""
	diskSize := ""
	herculesStartup := hercules + ".sh"
	switch execClient {
	case client_consts.Geth:
		switch network {
		case hestia_req_types.Ephemery, "ephemeral":
			dockerImage = gethDockerImageCapella
			diskSize = execClientDiskSizeEphemeral
			cmConfig = GethEphemeral
			herculesStartup = herculesEphemeral + ".sh"
		}
	}
	cp := filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/beacons/infra/exec_client",
		DirOut:      "./ethereum/outputs",
		FnIn:        execClient + "Hercules", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	sbCfg := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: cp,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			ConfigMapDriver: &zeus_topology_config_drivers.ConfigMapDriver{
				ConfigMap: v1Core.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{Name: "cm-exec-client"},
				},
				SwapKeys: map[string]string{
					"start.sh":       cmConfig + ".sh",
					hercules + ".sh": herculesStartup,
					download + ".sh": herculesStartup,
				},
			},
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					zeusExecClient: {
						Container: v1Core.Container{
							Name:  zeusExecClient,
							Image: dockerImage,
						},
					},
				},
				PVCDriver: &zeus_topology_config_drivers.PersistentVolumeClaimsConfigDriver{
					PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
						execClientDiskName: {
							ObjectMeta: metav1.ObjectMeta{Name: execClientDiskName},
							Spec: v1Core.PersistentVolumeClaimSpec{Resources: v1Core.ResourceRequirements{
								Requests: v1Core.ResourceList{"storage": resource.MustParse(diskSize)},
							}},
						},
					}},
			},
		}}
	execCompBase := zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			execClient + "Hercules": sbCfg,
		},
	}
	return execCompBase
}

func EphemeralExecClientGethConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.ConfigMap != nil {
		m := make(map[string]string)
		m = inf.ConfigMap.Data
		vSrc, ok := inf.ConfigMap.Data[GethEphemeral+".sh"]
		if ok {
			m[start+".sh"] = vSrc
		} else {
			err := errors.New("key not found")
			panic(err)
		}
		vSrc, ok = inf.ConfigMap.Data[downloadGethEphemeral+".sh"]
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
			if c.Name == zeusExecClient {
				inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = gethDockerImageCapella
			}
		}
		for i, v := range inf.StatefulSet.Spec.VolumeClaimTemplates {
			if v.Name == execClientDiskName {
				q, err := resource.ParseQuantity(execClientDiskSizeEphemeral)
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
