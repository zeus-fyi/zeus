package foundry_anvil

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

var (
	anvilClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "anvil",
		CloudCtxNs:       anvilCtxNs,
		ComponentBases:   anvilComponentBases,
	}
	anvilCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "anvil", // set with your own namespace
		Env:           "production",
	}
	anvilComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"anvil": anvilComponentBase,
	}
	anvilComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"anvil": anvilSkeletonBaseConfig,
		},
	}
	anvilSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: anvilChartPath,
	}
	anvilChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/foundry/anvil/infra",
		DirOut:      "./ethereum/outputs",
		FnIn:        "anvil", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
		FilterFiles: &strings_filter.FilterOpts{},
	}
)
