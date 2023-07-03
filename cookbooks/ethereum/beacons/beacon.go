package ethereum_beacon_cookbooks

import (
	choreography_cookbooks "github.com/zeus-fyi/zeus/cookbooks/microservices/choreography"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	start         = "start"
	download      = "download"
	initSnapshots = "init-snapshots"
)

var (
	BeaconClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereumEphemeralBeaconsDev",
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
			"lodestarHercules": ConsensusClientSkeletonBaseConfig,
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
	BearerTokenSecretFromChoreography = zeus_topology_config_drivers.MakeEnvVar("BEARER", "bearer", "choreography")
)

func GetClientClusterDef(consensusClient, execClient, network string) zeus_cluster_config_drivers.ClusterDefinition {
	return zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereumBeacon" + cases.Title(language.English).String(network) + cases.Title(language.English).String(consensusClient) + cases.Title(language.English).String(execClient),
		ComponentBases:   GetComponentBases(consensusClient, execClient, network),
	}
}

func GetComponentBases(consensusClient, execClient, network string) map[string]zeus_cluster_config_drivers.ComponentBaseDefinition {
	beaconComponentBases := map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"beaconIngress":                 IngressComponentBase,
		"consensusClients":              GetConsensusClientNetworkConfig(consensusClient, network, true),
		"execClients":                   GetExecClientNetworkConfig(execClient, network, true),
		"serviceMonitorConsensusClient": ConsensusClientMonitoringComponentBase,
		"serviceMonitorExecClient":      ExecClientMonitoringComponentBase,
		"choreography":                  choreography_cookbooks.ChoreographyComponentBase,
	}
	return beaconComponentBases
}
