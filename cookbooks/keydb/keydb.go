package zeus_keydb

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	keyDBClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "keydb",
		CloudCtxNs:       keyDBCloudCtxNs,
		ComponentBases:   keyDBComponentBases,
	}
	keyDBComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"keydb": keyDBComponentBase,
	}
	keyDBComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"keydb": keyDBSkeletonBaseConfig,
		},
	}
	keyDBSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: keyDBMasterChartPath,
	}
)
var (
	keyDBCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "keydb", // set with your own namespace
		Env:           "production",
	}
	keyDBMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./keydb/infra",
		DirOut:      "./keydb/outputs",
		FnIn:        "keydb", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)
