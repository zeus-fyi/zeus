package ethereum_beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/zeus/client/zeus_common_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var (
	BeaconGoerliClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereumGoerliBeacons",
		CloudCtxNs:       BeaconGoerliCloudCtxNs,
		ComponentBases:   BeaconGoerliComponentBases,
	}
	BeaconGoerliCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "goerli-beacon", // set with your own namespace
		Env:           "production",
	}
	BeaconGoerliComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"consensusClients": ConsensusClientGoerliComponentBase,
		"execClients":      ExecClientGoerliComponentBase,
	}
	ConsensusClientGoerliComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"lodestarHercules": ConsensusClientGoerliSkeletonBaseConfig,
		},
	}
	ExecClientGoerliComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"gethHercules": ExecClientGoerliSkeletonBaseConfig,
		},
	}
)
