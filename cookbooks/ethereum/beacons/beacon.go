package ethereum_beacon_cookbooks

import (
	choreography_cookbooks "github.com/zeus-fyi/zeus/cookbooks/microservices/choreography"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
)

const (
	start    = "start"
	download = "download"
)

var (
	BeaconClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereumEphemeralBeacons",
		CloudCtxNs:       BeaconCloudCtxNs,
		ComponentBases:   BeaconComponentBases,
	}
	BeaconCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "ephemeral", // set with your own namespace
		Env:           "production",
	}
	BeaconComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"beaconIngress":                 IngressComponentBase,
		"consensusClients":              ConsensusClientComponentBase,
		"execClients":                   ExecClientComponentBase,
		"serviceMonitorConsensusClient": ConsensusClientMonitoringComponentBase,
		"serviceMonitorExecClient":      ExecClientMonitoringComponentBase,
		"choreography":                  choreography_cookbooks.ChoreographyComponentBase,
	}
	ConsensusClientComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"lighthouseHercules": ConsensusClientSkeletonBaseConfig,
		},
	}
	ConsensusClientMonitoringComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"serviceMonitorConsensusClient": ConsensusClientSkeletonBaseMonitoringConfig,
		},
	}
	ExecClientComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"gethHercules": ExecClientSkeletonBaseConfig,
		},
	}
	ExecClientMonitoringComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"serviceMonitorExecClient": ExecClientSkeletonBaseMonitoringConfig,
		},
	}
	IngressComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"beaconIngress": BeaconIngressSkeletonBaseConfig,
		},
	}
)
