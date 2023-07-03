package cosmos_cookbooks

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var (
	CosmosTestnetNodeClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "cosmosNode",
		ComponentBases:   CosmosNodeComponentBases,
	}
	CosmosNodeComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"cosmosClients": CosmosNodeComponentBase,
	}
	CosmosNodeComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"cosmos": CosmosClientSkeletonBaseConfig,
		},
	}
	CosmosClientSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: CosmosClientChartPath,
	}
	CosmosClientChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./cosmos/node/infra",
		DirOut:      "./cosmos/outputs",
		FnIn:        "cosmos", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)
