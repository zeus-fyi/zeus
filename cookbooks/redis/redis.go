package zeus_redis

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	redisClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "redis",
		CloudCtxNs:       redisCloudCtxNs,
		ComponentBases:   redisComponentBases,
	}
	redisComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"master":   redisMasterComponentBase,
		"replicas": redisReplicasComponentBase,
	}
	redisMasterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"redis-master": redisMasterSkeletonBaseConfig,
		},
	}
	redisReplicasComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"redis-replicas": redisReplicasSkeletonBaseConfig,
		},
	}
	redisMasterSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: redisMasterChartPath,
	}
	redisReplicasSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: redisReplicasChartPath,
	}
)
var (
	redisCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "redis", // set with your own namespace
		Env:           "production",
	}
	redisMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./redis/master",
		DirOut:      "./redis/outputs",
		FnIn:        "redisMaster", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	redisReplicasChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./redis/replicas",
		DirOut:      "./redis/outputs",
		FnIn:        "redisReplicas", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)
