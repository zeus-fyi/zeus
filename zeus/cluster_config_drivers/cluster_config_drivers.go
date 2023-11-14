package zeus_cluster_config_drivers

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

type ClusterDefinition struct {
	ClusterClassName          string
	CloudCtxNs                zeus_common_types.CloudCtxNs
	ComponentBases            map[string]ComponentBaseDefinition
	FilterSkeletonBaseUploads *strings_filter.FilterOpts
	DisablePrint              bool
	UseEmbeddedWorkload       bool
}

type ComponentBaseDefinition struct {
	SkeletonBases map[string]ClusterSkeletonBaseDefinition
}

type ClusterSkeletonBaseDefinition struct {
	SkeletonBaseChart         zeus_req_types.TopologyCreateRequest
	SkeletonBaseNameChartPath filepaths.Path

	TopologyID           int
	Workload             topology_workloads.TopologyBaseInfraWorkload
	TopologyConfigDriver *zeus_topology_config_drivers.TopologyConfigDriver
}

type ClusterSkeletonBaseDefinitions []ClusterSkeletonBaseDefinition

func (c *ClusterDefinition) UploadChartsFromClusterDefinition(ctx context.Context, z zeus_client.ZeusClient, print bool) ([]zeus_resp_types.TopologyCreateResponse, error) {
	sbs, err := c.GenerateSkeletonBaseCharts()
	if err != nil {
		log.Err(err)
		return nil, err
	}
	responses := make([]zeus_resp_types.TopologyCreateResponse, len(sbs))
	for i, sb := range sbs {
		if !strings_filter.FilterStringWithOpts(sb.SkeletonBaseChart.SkeletonBaseName, c.FilterSkeletonBaseUploads) {
			i -= 1
			continue
		}
		resp, rerr := z.UploadChart(ctx, sb.SkeletonBaseNameChartPath, sb.SkeletonBaseChart)
		if rerr != nil {
			log.Err(err)
			return responses, rerr
		}
		if print {
			tar := zeus_req_types.TopologyRequest{TopologyID: resp.TopologyID}
			chartResp, cerr := z.ReadChart(ctx, tar)
			if cerr != nil {
				log.Err(cerr)
			}
			cerr = chartResp.PrintWorkload(sb.SkeletonBaseNameChartPath)
			if cerr != nil {
				log.Err(cerr)
			}
		}
		responses[i] = resp
	}
	return responses, err
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

func (c *ClusterDefinition) GenerateSkeletonBaseCharts() ([]ClusterSkeletonBaseDefinition, error) {
	sbDefinitons := []ClusterSkeletonBaseDefinition{}
	for cbName, cb := range c.ComponentBases {
		for sbName, sb := range cb.SkeletonBases {
			inf := topology_workloads.NewTopologyBaseInfraWorkload()
			if c.UseEmbeddedWorkload {
				inf = sb.Workload
			} else {
				err := sb.SkeletonBaseNameChartPath.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
				if err != nil {
					log.Err(err)
					return []ClusterSkeletonBaseDefinition{}, err
				}
			}
			// This will customize your config with the supplied workload override supplied
			if sb.TopologyConfigDriver != nil {
				sb.TopologyConfigDriver.SetCustomConfig(&inf)
				if !c.DisablePrint || c.UseEmbeddedWorkload {
					tmp := sb.SkeletonBaseNameChartPath.DirOut
					dir, _ := filepath.Split(sb.SkeletonBaseNameChartPath.DirIn)
					lastDir := strings.Split(dir, "/")[len(strings.Split(dir, "/"))-1]
					newPath := fmt.Sprintf("%scustom_%s", dir[:len(dir)-len(lastDir)], sbName)
					sb.SkeletonBaseNameChartPath.DirOut = newPath
					err := inf.PrintWorkload(sb.SkeletonBaseNameChartPath)
					if err != nil {
						log.Err(err)
						return []ClusterSkeletonBaseDefinition{}, err
					}
					sb.SkeletonBaseNameChartPath.DirOut = tmp
					sb.SkeletonBaseNameChartPath.DirIn = newPath
					err = sb.SkeletonBaseNameChartPath.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
					if err != nil {
						log.Err(err)
						return []ClusterSkeletonBaseDefinition{}, err
					}
				}
			}

			err := inf.ValidateWorkloads()
			if err != nil {
				log.Err(err)
				return []ClusterSkeletonBaseDefinition{}, err
			}
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
				},
				SkeletonBaseNameChartPath: sb.SkeletonBaseNameChartPath,
				Workload:                  inf,
			}
			c.ComponentBases[cbName].SkeletonBases[sbName] = sbDef
			sbDefinitons = append(sbDefinitons, sbDef)
		}
	}
	return sbDefinitons, nil
}
