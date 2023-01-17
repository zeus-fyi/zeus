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

func (c *ClusterDefinition) GenerateSkeletonBaseCharts() []ClusterSkeletonBaseDefinition {
	sbDefinitons := []ClusterSkeletonBaseDefinition{}
	for cbName, cb := range c.ComponentBases {
		for sbName, sb := range cb.SkeletonBases {
			sbDef := ClusterSkeletonBaseDefinition{
				SkeletonBaseChart: zeus_req_types.TopologyCreateRequest{
					TopologyName:      c.ClusterClassName,
					ClusterClassName:  c.ClusterClassName,
					ComponentBaseName: cbName,
					ChartName:         cbName,
					ChartDescription:  fmt.Sprintf("%s-%s-%s", c.ClusterClassName, cbName, sbName),
					SkeletonBaseName:  sbName,
					Tag:               "latest",
					Version:           fmt.Sprintf("v0.0.%d", time.Now().Unix()),
				}, SkeletonBaseNameChartPath: sb.SkeletonBaseNameChartPath,
			}
			if sb.TopologyConfigDriver != nil {
				// Customize Config
				// TODO, parse files & apply custom config, and generate new outputs
			}
			sbDefinitons = append(sbDefinitons, sbDef)
		}
	}
	return sbDefinitons
}
