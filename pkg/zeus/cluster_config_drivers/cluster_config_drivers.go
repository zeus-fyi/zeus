package zeus_cluster_config_drivers

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
)

type ClusterDefinition struct {
	ClusterClassName string
	CloudCtxNs       zeus_common_types.CloudCtxNs
	ComponentBases   map[string]ComponentBaseDefinition
}

type ComponentBaseDefinition struct {
	SkeletonBases map[string]ClusterSkeletonBaseDefinition
}

type ClusterSkeletonBaseDefinition struct {
	SkeletonBaseChart         zeus_req_types.TopologyCreateRequest
	SkeletonBaseNameChartPath filepaths.Path
	TopologyConfigDriver      *zeus_topology_config_drivers.TopologyConfigDriver
}

func (c *ClusterDefinition) GenerateDeploymentRequest() zeus_req_types.ClusterTopologyDeployRequest {
	sbNameSlice := []string{}
	for _, componentBase := range c.ComponentBases {
		for skeletonBaseName, _ := range componentBase.SkeletonBases {
			sbNameSlice = append(sbNameSlice, skeletonBaseName)
		}
	}
	return zeus_req_types.ClusterTopologyDeployRequest{
		ClusterClassName:    c.ClusterClassName,
		SkeletonBaseOptions: sbNameSlice,
		CloudCtxNs:          c.CloudCtxNs,
	}
}

func (c *ClusterDefinition) GenerateSkeletonBaseChart(clusterClassName, componentBaseName, skeletonBaseName string) ClusterSkeletonBaseDefinition {
	componentBase, ok := c.ComponentBases[componentBaseName]
	if !ok {
		return ClusterSkeletonBaseDefinition{}
	}
	skeletonBase, ok := componentBase.SkeletonBases[skeletonBaseName]
	if !ok {
		return ClusterSkeletonBaseDefinition{}
	}
	cp := skeletonBase.SkeletonBaseNameChartPath
	cp.FnIn = skeletonBaseName
	return ClusterSkeletonBaseDefinition{
		SkeletonBaseChart: zeus_req_types.TopologyCreateRequest{
			TopologyName:      clusterClassName,
			ClusterClassName:  clusterClassName,
			ComponentBaseName: componentBaseName,
			ChartName:         componentBaseName,
			ChartDescription:  fmt.Sprintf("%s-%s-%s", clusterClassName, componentBaseName, skeletonBaseName),
			SkeletonBaseName:  skeletonBaseName,
			Tag:               "latest",
			Version:           fmt.Sprintf("v0.0.%d", time.Now().Unix()),
		}, SkeletonBaseNameChartPath: skeletonBase.SkeletonBaseNameChartPath,
	}
}
