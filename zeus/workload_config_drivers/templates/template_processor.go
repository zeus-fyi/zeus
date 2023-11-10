package zk8s_templates

import (
	"context"

	"github.com/rs/zerolog/log"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
	"k8s.io/api/core/v1"
	//v1Apps "k8s.io/api/apps/v1"
)

type ClusterPreviewWorkloads struct {
	ClusterName    string                                                             `json:"clusterName"`
	ComponentBases map[string]map[string]topology_workloads.TopologyBaseInfraWorkload `json:"componentBases"`
}

type ClusterPreviewWorkloadsOlympus struct {
	ClusterName    string                    `json:"clusterName"`
	ComponentBases map[string]map[string]any `json:"componentBases"`
}

type WorkloadDefinition struct {
	WorkloadName string
	ReplicaCount int
	Containers   Containers
}

type Containers map[string]Container

type Container struct {
	IsInitContainer bool        `json:"isInitContainer"`
	ImagePullPolicy string      `json:"imagePullPolicy,omitempty"`
	DockerImage     DockerImage `json:"dockerImage"`
}

type DockerImage struct {
	ImageName            string               `json:"imageName"`
	Cmd                  string               `json:"cmd"`
	Args                 string               `json:"args"`
	ResourceRequirements ResourceRequirements `json:"resourceRequirements,omitempty"`
	Ports                []Port               `json:"ports,omitempty"`
	VolumeMounts         []VolumeMount        `json:"volumeMounts,omitempty"`
}

func GenerateDeploymentCluster(ctx context.Context, wd WorkloadDefinition) (*Cluster, error) {
	componentBases := map[string]SkeletonBases{
		wd.WorkloadName: {
			wd.WorkloadName: SkeletonBase{
				Containers:    wd.Containers,
				AddDeployment: true,
				AddIngress:    true,
				AddService:    true,
			},
		},
	}
	ingressPaths := map[string]IngressPath{}
	c := &Cluster{
		ClusterName:    wd.WorkloadName,
		IngressPaths:   ingressPaths,
		ComponentBases: componentBases,
	}
	return c, nil
}

func GenerateSkeletonBaseChartsPreview(ctx context.Context, cluster Cluster) (ClusterPreviewWorkloads, error) {
	pcg := ClusterPreviewWorkloads{
		ClusterName:    cluster.ClusterName,
		ComponentBases: make(map[string]map[string]topology_workloads.TopologyBaseInfraWorkload),
	}
	cd := PreviewTemplateGeneration(ctx, cluster)
	cd.UseEmbeddedWorkload = true
	cd.DisablePrint = true
	_, err := cd.GenerateSkeletonBaseCharts()
	if err != nil {
		log.Err(err)
		return pcg, err
	}
	for cbName, componentBase := range cd.ComponentBases {
		pcg.ComponentBases[cbName] = make(map[string]topology_workloads.TopologyBaseInfraWorkload)
		for sbName, skeletonBase := range componentBase.SkeletonBases {
			pcg.ComponentBases[cbName][sbName] = skeletonBase.Workload
		}
	}
	return pcg, nil
}

func PreviewTemplateGeneration(ctx context.Context, cluster Cluster) zeus_cluster_config_drivers.ClusterDefinition {
	templateClusterDefinition := zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: cluster.ClusterName,
		ComponentBases:   make(map[string]zeus_cluster_config_drivers.ComponentBaseDefinition),
	}
	for cbName, componentBase := range cluster.ComponentBases {
		cbDef := zeus_cluster_config_drivers.ComponentBaseDefinition{
			SkeletonBases: make(map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition),
		}
		for sbName, skeletonBase := range componentBase {
			sbDef := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
				SkeletonBaseChart:    zeus_req_types.TopologyCreateRequest{},
				Workload:             topology_workloads.TopologyBaseInfraWorkload{},
				TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{},
			}
			if skeletonBase.AddStatefulSet {
				sbDef.Workload.StatefulSet = GetStatefulSetTemplate(ctx, cbName)
				stsDriver, err := BuildStatefulSetDriver(ctx, skeletonBase.Containers, skeletonBase.StatefulSet)
				if err != nil {
					log.Err(err).Msg("error building statefulset driver")
				}
				sbDef.TopologyConfigDriver.StatefulSetDriver = &stsDriver
			} else if skeletonBase.AddDeployment {
				sbDef.Workload.Deployment = GetDeploymentTemplate(ctx, cbName)
				depDriver, err := BuildDeploymentDriver(ctx, skeletonBase.Containers, skeletonBase.Deployment)
				if err != nil {
					log.Err(err).Msg("error building deployment driver")
				}
				sbDef.TopologyConfigDriver.DeploymentDriver = &depDriver
			}
			if skeletonBase.AddIngress {
				sbDef.Workload.Ingress = GetIngressTemplate(ctx, cbName)
				ingDriver, err := BuildIngressDriver(ctx, cbName, skeletonBase.Containers, cluster.IngressSettings, cluster.IngressPaths)
				if err != nil {
					log.Err(err).Msg("error building ingress driver")
				}
				sbDef.TopologyConfigDriver.IngressDriver = &ingDriver
			}
			if skeletonBase.AddService {
				sbDef.Workload.Service = GetServiceTemplate(ctx, cbName)
				svcDriver, err := BuildServiceDriver(ctx, skeletonBase.Containers)
				if err != nil {
					log.Err(err).Msg("error building service driver")
				}
				sbDef.TopologyConfigDriver.ServiceDriver = &svcDriver
			}
			if skeletonBase.AddConfigMap {
				sbDef.Workload.ConfigMap = GetConfigMapTemplate(ctx, cbName)
				cmDriver, err := BuildConfigMapDriver(ctx, skeletonBase.ConfigMap)
				if err != nil {
					log.Err(err).Msg("error building configmap driver")
				}
				sbDef.TopologyConfigDriver.ConfigMapDriver = &cmDriver
			}
			cbDef.SkeletonBases[sbName] = sbDef
		}
		templateClusterDefinition.ComponentBases[cbName] = cbDef
	}
	return templateClusterDefinition
}

func BuildConfigMapDriver(ctx context.Context, configMap ConfigMap) (zeus_topology_config_drivers.ConfigMapDriver, error) {
	cmDriver := zeus_topology_config_drivers.ConfigMapDriver{
		ConfigMap: v1.ConfigMap{
			Data: make(map[string]string),
		},
	}
	for key, value := range configMap {
		cmDriver.ConfigMap.Data[key] = value
	}
	return cmDriver, nil
}
