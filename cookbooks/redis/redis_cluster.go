package zeus_redis

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	redisClusterConfigDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "redis-cluster",
		CloudCtxNs:       redisClusterCloudCtxNs,
		ComponentBases:   redisClusterConfigComponentBases,
	}
	redisClusterConfigComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"redis-cluster": redisClusterComponentBase,
	}
	redisClusterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"redis-cluster": redisClusterSkeletonBaseConfig,
		},
	}
	redisClusterSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: redisClusterChartPath,
	}
)
var (
	redisClusterCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "redis-cluster", // set with your own namespace
		Env:           "production",
	}
	redisClusterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./redis/cluster",
		DirOut:      "./redis/outputs",
		FnIn:        "redisCluster", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)
