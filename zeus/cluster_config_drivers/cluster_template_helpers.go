package zeus_cluster_config_drivers

import (
	"context"

	"github.com/rs/zerolog/log"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	zk8s_templates "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/templates"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
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
	Containers   zk8s_templates.Containers
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
	ingressPaths := map[string]zk8s_templates.IngressPath{}
	c := &Cluster{
		ClusterName:    wd.WorkloadName,
		IngressPaths:   ingressPaths,
		ComponentBases: componentBases,
	}
	return c, nil
}

type Cluster struct {
	ClusterName     string                                `json:"clusterName"`
	ComponentBases  ComponentBases                        `json:"componentBases"`
	IngressSettings zk8s_templates.Ingress                `json:"ingressSettings"`
	IngressPaths    map[string]zk8s_templates.IngressPath `json:"ingressPaths"`
}

type ComponentBases map[string]SkeletonBases

type SkeletonBases map[string]SkeletonBase

type SkeletonBase struct {
	TopologyID        string                     `json:"topologyID,omitempty"`
	AddStatefulSet    bool                       `json:"addStatefulSet"`
	AddDeployment     bool                       `json:"addDeployment"`
	AddConfigMap      bool                       `json:"addConfigMap"`
	AddService        bool                       `json:"addService"`
	AddIngress        bool                       `json:"addIngress"`
	AddServiceMonitor bool                       `json:"addServiceMonitor"`
	ConfigMap         zk8s_templates.ConfigMap   `json:"configMap,omitempty"`
	Deployment        zk8s_templates.Deployment  `json:"deployment,omitempty"`
	StatefulSet       zk8s_templates.StatefulSet `json:"statefulSet,omitempty"`
	Containers        zk8s_templates.Containers  `json:"containers,"`
}

func CreateGeneratedClusterClassCreationRequest(c *Cluster) GeneratedClusterCreationRequests {
	var cbn []string
	var sbns []zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest
	for componentBaseName, sb := range c.ComponentBases {
		if componentBaseName == "" {
			continue
		}
		cbn = append(cbn, componentBaseName)
		sbComp := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
			ClusterClassName:  c.ClusterName,
			ComponentBaseName: componentBaseName,
			SkeletonBaseNames: []string{},
		}
		for sbName, _ := range sb {
			if sbName == "" {
				continue
			}
			sbComp.SkeletonBaseNames = append(sbComp.SkeletonBaseNames, sbName)
		}
		sbns = append(sbns, sbComp)
	}
	gcd := GeneratedClusterCreationRequests{
		ClusterClassRequest: zeus_req_types.TopologyCreateClusterClassRequest{
			ClusterClassName: c.ClusterName,
		},
		ComponentBasesRequests: zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
			ClusterClassName:   c.ClusterName,
			ComponentBaseNames: cbn,
		},
		SkeletonBasesRequests: sbns,
	}

	return gcd
}

/*
	gcd := DocusaurusClusterDefinition.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
*/

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

func PreviewTemplateGeneration(ctx context.Context, cluster Cluster) ClusterDefinition {
	templateClusterDefinition := ClusterDefinition{
		ClusterClassName: cluster.ClusterName,
		ComponentBases:   make(map[string]ComponentBaseDefinition),
	}
	for cbName, componentBase := range cluster.ComponentBases {
		cbDef := ComponentBaseDefinition{
			SkeletonBases: make(map[string]ClusterSkeletonBaseDefinition),
		}
		for sbName, skeletonBase := range componentBase {
			sbDef := ClusterSkeletonBaseDefinition{
				SkeletonBaseChart:    zeus_req_types.TopologyCreateRequest{},
				Workload:             topology_workloads.TopologyBaseInfraWorkload{},
				TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{},
			}
			if skeletonBase.AddStatefulSet {
				sbDef.Workload.StatefulSet = zk8s_templates.GetStatefulSetTemplate(ctx, cbName)
				stsDriver, err := zk8s_templates.BuildStatefulSetDriver(ctx, skeletonBase.Containers, skeletonBase.StatefulSet)
				if err != nil {
					log.Err(err).Msg("error building statefulset driver")
				}
				sbDef.TopologyConfigDriver.StatefulSetDriver = &stsDriver
			} else if skeletonBase.AddDeployment {
				sbDef.Workload.Deployment = zk8s_templates.GetDeploymentTemplate(ctx, cbName)
				depDriver, err := zk8s_templates.BuildDeploymentDriver(ctx, skeletonBase.Containers, skeletonBase.Deployment)
				if err != nil {
					log.Err(err).Msg("error building deployment driver")
				}
				sbDef.TopologyConfigDriver.DeploymentDriver = &depDriver
			}
			if skeletonBase.AddIngress {
				sbDef.Workload.Ingress = zk8s_templates.GetIngressTemplate(ctx, cbName)
				ingDriver, err := zk8s_templates.BuildIngressDriver(ctx, cbName, skeletonBase.Containers, cluster.IngressSettings, cluster.IngressPaths)
				if err != nil {
					log.Err(err).Msg("error building ingress driver")
				}
				sbDef.TopologyConfigDriver.IngressDriver = &ingDriver
			}
			if skeletonBase.AddService {
				sbDef.Workload.Service = zk8s_templates.GetServiceTemplate(ctx, cbName)
				svcDriver, err := zk8s_templates.BuildServiceDriver(ctx, skeletonBase.Containers)
				if err != nil {
					log.Err(err).Msg("error building service driver")
				}
				sbDef.TopologyConfigDriver.ServiceDriver = &svcDriver
			}
			if skeletonBase.AddConfigMap {
				sbDef.Workload.ConfigMap = zk8s_templates.GetConfigMapTemplate(ctx, cbName)
				cmDriver, err := zk8s_templates.BuildConfigMapDriver(ctx, skeletonBase.ConfigMap)
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
