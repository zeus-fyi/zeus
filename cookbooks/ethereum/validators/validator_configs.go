package validator_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// TODO, will use matrix class later. Cluster class is a good near-term substitute
// extends on the beacon cluster test cases for now
var (
	EphemeryValidatorClusterClassName         = "ethereumEphemeralValidatorCluster"
	consensusValidatorClientComponentBaseName = "consensusValidatorClients"
	ValidatorSkeletonBaseName                 = "lighthouseHerculesValidatorClient"

	ExecSkeletonBase         = "gethHercules"
	ConsensusSkeletonBase    = "lighthouseHercules"
	ChoreographySkeletonBase = "choreography"
)

var cd = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName:    EphemeryValidatorClusterClassName,
	SkeletonBaseOptions: []string{ExecSkeletonBase, ConsensusSkeletonBase, ValidatorSkeletonBaseName, ChoreographySkeletonBase},
	CloudCtxNs:          ValidatorCloudCtxNs,
}

var DeployConsensusValidatorClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: ValidatorCloudCtxNs,
}

var ValidatorCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "ephemeral-staking", // set with your own namespace
	Env:           "production",
}

// chart workload metadata
var validatorsChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      ValidatorSkeletonBaseName,
	ChartName:         ValidatorSkeletonBaseName,
	ChartDescription:  ValidatorSkeletonBaseName,
	Version:           fmt.Sprintf("validatorBase-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  ValidatorSkeletonBaseName,
	ComponentBaseName: consensusValidatorClientComponentBaseName,
	ClusterClassName:  EphemeryValidatorClusterClassName,
	Tag:               "latest",
}

var validatorsChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/validators/infra/validators",
	DirOut:      "./ethereum/validators/infra/processed_validators",
	FnIn:        ValidatorSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
