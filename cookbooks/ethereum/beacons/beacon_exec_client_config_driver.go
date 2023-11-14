package ethereum_beacon_cookbooks

import (
	"errors"

	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	zeusExecClient              = "zeus-exec-client"
	execClientDiskName          = "exec-client-storage"
	execClientDiskSizeEphemeral = "10Gi"
	execClientDiskSizeGoerli    = "500Gi"
	execClientDiskSizeMainnet   = "2Ti"

	execClientEphemeralRequestCPU    = "2.5"
	execClientEphemeralRequestMemory = "4Gi"
	execClientMainnetCPU             = "7"
	execClientMainnetCPUMemory       = "20Gi"

	hercules          = "hercules"
	herculesEphemeral = "herculesEphemeral"
	GethEphemeral     = "gethEphemeral"
	GethGoerli        = "gethGoerli"
	GethMainnet       = "geth"

	downloadGethEphemeral = "downloadGethEphemeral"
	gethDockerImage       = "ethereum/client-go:v1.12.0"

	gethDockerImageCapella = "ethpandaops/geth:master"

	rethDockerImage = "ghcr.io/paradigmxyz/reth:v0.1.0-alpha.1"
	rethMainnet     = "reth"
	rethDownload    = "downloadReth"

	execClientRethDiskSizeMainnet = "3Ti"
)

func GetExecClientNetworkConfig(execClient, network string, choreographySecretsExist bool) zeus_cluster_config_drivers.ComponentBaseDefinition {
	dockerImage := ""
	cmConfig := ""
	diskSize := execClientDiskSizeMainnet
	herculesStartup := hercules + ".sh"
	downloadStartup := download + ".sh"
	cpuSize := execClientMainnetCPU
	memSize := execClientMainnetCPUMemory

	switch execClient {
	case client_consts.Geth:
		switch network {
		case hestia_req_types.Ephemery, "ephemeral":
			cpuSize = execClientEphemeralRequestCPU
			memSize = execClientEphemeralRequestMemory
			dockerImage = gethDockerImageCapella
			diskSize = execClientDiskSizeEphemeral
			cmConfig = GethEphemeral
			herculesStartup = herculesEphemeral + ".sh"
			downloadStartup = downloadGethEphemeral + ".sh"
		case hestia_req_types.Mainnet:
			cmConfig = GethMainnet
		}
	case client_consts.Reth:
		diskSize = execClientRethDiskSizeMainnet
		cmConfig = rethMainnet
		dockerImage = rethDockerImage
		downloadStartup = rethDownload + ".sh"
	}
	cp := filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/beacons/infra/exec_client",
		DirOut:      "./ethereum/outputs",
		FnIn:        execClient + "Hercules", // filename for your gzip workload
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
	rr := v1Core.ResourceRequirements{
		Requests: v1Core.ResourceList{
			"cpu":    resource.MustParse(cpuSize),
			"memory": resource.MustParse(memSize),
		},
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
					download + ".sh": downloadStartup,
				},
			},
			StatefulSetDriver: &zeus_topology_config_drivers.StatefulSetDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					initSnapshots: initContDriver,
					zeusExecClient: {
						Container: v1Core.Container{
							Name:      zeusExecClient,
							Image:     dockerImage,
							Resources: rr,
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
