package zeus_keydb

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	keyDBClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "keyDB",

		CloudCtxNs:     keyDBCloudCtxNs,
		ComponentBases: keyDBComponentBases,
	}
	keyDBComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"master":   keyDBMasterComponentBase,
		"replicas": keyDBReplicasComponentBase,
	}
	keyDBMasterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"keyDB-master": keyDBMasterSkeletonBaseConfig,
		},
	}
	keyDBReplicasComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"keyDB-replicas": keyDBReplicasSkeletonBaseConfig,
		},
	}
	keyDBMasterSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: keyDBMasterChartPath,
	}
	keyDBReplicasSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: keyDBReplicasChartPath,
	}
)
var (
	keyDBCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "keyDB", // set with your own namespace
		Env:           "production",
	}
	keyDBMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./keydb/master",
		DirOut:      "./keydb/outputs",
		FnIn:        "keydb", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	keyDBReplicasChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./keydb/replicas",
		DirOut:      "./keydb/outputs",
		FnIn:        "keydb", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)
