package validator_cookbooks

import (
	"fmt"
	"time"

	ethereum_beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	choreography_cookbooks "github.com/zeus-fyi/zeus/cookbooks/microservices/choreography"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
)

var (
	EphemeryValidatorClusterClassName         = "ethereumEphemeralValidatorCluster"
	consensusValidatorClientComponentBaseName = "consensusValidatorClients"
	ValidatorSkeletonBaseName                 = "lighthouseHerculesValidatorClient"

	ExecSkeletonBase         = "gethHercules"
	ConsensusSkeletonBase    = "lighthouseHercules"
	ChoreographySkeletonBase = "choreography"

	EphemeryValidatorClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: EphemeryValidatorClusterClassName,
		CloudCtxNs:       ValidatorCloudCtxNs,
		ComponentBases:   ValidatorComponentBases,
	}
	ValidatorCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "ephemeral-staking", // set with your own namespace
		Env:           "production",
	}
	ValidatorComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"consensusClients":          ConsensusClientComponentBase,
		"execClients":               ExecClientComponentBase,
		"consensusValidatorClients": ValidatorClientComponentBase,
		"choreography":              choreography_cookbooks.ChoreographyComponentBase,
	}
	ConsensusClientComponentBase = ethereum_beacon_cookbooks.ConsensusClientComponentBase
	ExecClientComponentBase      = ethereum_beacon_cookbooks.ExecClientComponentBase
	ValidatorClientComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"lighthouseHerculesValidatorClient": ValidatorClientClientSkeletonBaseConfig,
		},
	}
	ValidatorClientClientSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: ValidatorsChartPath,
	}
)

var ValidatorClusterDefinition = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName:    EphemeryValidatorClusterClassName,
	SkeletonBaseOptions: []string{ExecSkeletonBase, ConsensusSkeletonBase, ValidatorSkeletonBaseName, ChoreographySkeletonBase},
	CloudCtxNs:          ValidatorCloudCtxNs,
}

var DeployConsensusValidatorClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: ValidatorCloudCtxNs,
}

var ValidatorsChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      ValidatorSkeletonBaseName,
	ChartName:         ValidatorSkeletonBaseName,
	ChartDescription:  ValidatorSkeletonBaseName,
	Version:           fmt.Sprintf("validatorBase-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  ValidatorSkeletonBaseName,
	ComponentBaseName: consensusValidatorClientComponentBaseName,
	ClusterClassName:  EphemeryValidatorClusterClassName,
	Tag:               "latest",
}

var ValidatorsChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/validators/infra/validators",
	DirOut:      "./ethereum/validators/infra/processed_validators",
	FnIn:        ValidatorSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
