package ethereum_beacon_cookbooks

import (
	choreography_cookbooks "github.com/zeus-fyi/zeus/cookbooks/microservices/choreography"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

const (
	start         = "start"
	download      = "download"
	initSnapshots = "init-snapshots"
)

var (
	BeaconClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereum-ephemeral-beacons",
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
		"beacon-ingress":                  IngressComponentBase,
		"consensus-clients":               ConsensusClientComponentBase,
		"exec-clients":                    ExecClientComponentBase,
		"servicemonitor-consensus-client": ConsensusClientMonitoringComponentBase,
		"servicemonitor-exec-client":      ExecClientMonitoringComponentBase,
		"choreography":                    choreography_cookbooks.ChoreographyComponentBase,
	}
	ConsensusClientComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"lodestar-hercules": ConsensusClientSkeletonBaseConfig,
		},
	}
	ConsensusClientMonitoringComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"servicemonitor-consensus-client": ConsensusClientSkeletonBaseMonitoringConfig,
		},
	}
	ExecClientComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"geth-hercules": ExecClientSkeletonBaseConfig,
		},
	}
	ExecClientMonitoringComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"servicemonitor-exec-client": ExecClientSkeletonBaseMonitoringConfig,
		},
	}
	IngressComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"beacon-ingress": BeaconIngressSkeletonBaseConfig,
		},
	}
	BearerTokenSecretFromChoreography = zeus_topology_config_drivers.MakeSecretEnvVar("BEARER", "bearer", "choreography")
)

type BeaconConfig struct {
	ConsensusClient    string
	ExecClient         string
	Network            string
	WithIngress        bool
	WithServiceMonitor bool
	WithChoreography   bool
}

func GetClientClusterDef(consensusClient, execClient, network string, withIngress bool) zeus_cluster_config_drivers.ClusterDefinition {
	beaconConfig := BeaconConfig{
		ConsensusClient:    consensusClient,
		ExecClient:         execClient,
		Network:            network,
		WithIngress:        withIngress,
		WithServiceMonitor: true,
	}
	return zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereum-beacon" + "-" + beaconConfig.Network + "-" + beaconConfig.ConsensusClient + "-" + beaconConfig.ExecClient,
		ComponentBases:   GetComponentBases(beaconConfig),
	}
}

func CreateClientClusterDefWithParams(beaconConfig BeaconConfig) zeus_cluster_config_drivers.ClusterDefinition {
	return zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereum-beacon" + "-" + beaconConfig.Network + "-" + beaconConfig.ConsensusClient + "-" + beaconConfig.ExecClient,
		ComponentBases:   GetComponentBases(beaconConfig),
	}
}

func GetComponentBases(beaconConfig BeaconConfig) map[string]zeus_cluster_config_drivers.ComponentBaseDefinition {
	beaconComponentBases := map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"consensus-clients": GetConsensusClientNetworkConfig(beaconConfig),
		"exec-clients":      GetExecClientNetworkConfig(beaconConfig),
	}
	if beaconConfig.WithServiceMonitor {
		beaconComponentBases["servicemonitor-consensus-client"] = ConsensusClientMonitoringComponentBase
		beaconComponentBases["servicemonitor-exec-client"] = ExecClientMonitoringComponentBase
	}
	if beaconConfig.WithChoreography {
		beaconComponentBases["choreography"] = choreography_cookbooks.ChoreographyComponentBase
	}
	if beaconConfig.WithIngress {
		beaconComponentBases["beacon-ingress"] = IngressComponentBase
	}
	return beaconComponentBases
}
