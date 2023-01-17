package zeus_cluster_config_drivers

import (
	"context"

	"github.com/rs/zerolog/log"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

type GeneratedClusterCreationRequests struct {
	zeus_client.ZeusClient
	ClusterClassRequest    zeus_req_types.TopologyCreateClusterClassRequest
	ComponentBasesRequests zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest
	SkeletonBasesRequests  []zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest
}

func (c *ClusterDefinition) BuildClusterDefinitions() GeneratedClusterCreationRequests {
	var gcd GeneratedClusterCreationRequests
	gcd.ClusterClassRequest = zeus_req_types.TopologyCreateClusterClassRequest{ClusterClassName: c.ClusterClassName}
	gcd.ComponentBasesRequests = zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   c.ClusterClassName,
		ComponentBaseNames: make([]string, len(c.ComponentBases)),
	}
	gcd.SkeletonBasesRequests = []zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{}
	i := 0
	for componentBaseName, componentBase := range c.ComponentBases {
		gcd.ComponentBasesRequests.ComponentBaseNames[i] = componentBaseName
		sbDefinition := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
			ClusterClassName:  c.ClusterClassName,
			ComponentBaseName: componentBaseName,
			SkeletonBaseNames: make([]string, len(componentBase.SkeletonBases)),
		}
		j := 0
		for skeletonBaseName, _ := range componentBase.SkeletonBases {
			sbDefinition.SkeletonBaseNames[j] = skeletonBaseName
			j++
		}
		gcd.SkeletonBasesRequests = append(gcd.SkeletonBasesRequests, sbDefinition)
		i++
	}
	return gcd
}

func (gcd *GeneratedClusterCreationRequests) CreateClusterClass(ctx context.Context) error {
	_, err := gcd.CreateClass(ctx, gcd.ClusterClassRequest)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return err
	}
	_, err = gcd.AddComponentBasesToClass(ctx, gcd.ComponentBasesRequests)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return err
	}
	for _, sb := range gcd.SkeletonBasesRequests {
		_, err = gcd.AddSkeletonBasesToClass(ctx, sb)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return err
		}
	}
	return err
}
