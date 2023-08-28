package sui_cookbooks

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	suiNodeDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "sui",
		CloudCtxNs:       suiCloudCtxNs,
		ComponentBases:   suiComponentBases,
	}
	suiComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"sui": suiMasterComponentBase,
	}
	suiMasterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"sui": suiSkeletonBaseConfig,
		},
	}
	suiSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: suiMasterChartPath,
	}
)
var (
	suiCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "sui", // set with your own namespace
		Env:           "production",
	}
	suiMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./sui/infra",
		DirOut:      "./sui/infra",
		FnIn:        "sui", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)
