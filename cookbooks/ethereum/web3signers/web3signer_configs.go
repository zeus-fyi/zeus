package web3signer_cookbooks

import (
	"fmt"
	"time"

	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

var (
	web3SignerComponentBaseName = "web3Signer"
	web3SignerSkeletonBaseName  = "web3Signer"
)

var cd = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName: validator_cookbooks.EphemeryValidatorClusterClassName,
	SkeletonBaseOptions: []string{
		validator_cookbooks.ExecSkeletonBase,
		validator_cookbooks.ConsensusSkeletonBase,
		validator_cookbooks.ValidatorSkeletonBaseName,
		validator_cookbooks.ChoreographySkeletonBase,
		web3SignerSkeletonBaseName,
	},
	CloudCtxNs: validator_cookbooks.ValidatorCloudCtxNs,
}

var DeployWeb3SignerKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: validator_cookbooks.ValidatorCloudCtxNs,
}

// chart workload metadata
var web3SignerChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      web3SignerSkeletonBaseName,
	ChartName:         web3SignerSkeletonBaseName,
	ChartDescription:  web3SignerSkeletonBaseName,
	Version:           fmt.Sprintf("web3SignerBase-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  web3SignerSkeletonBaseName,
	ComponentBaseName: web3SignerComponentBaseName,
	ClusterClassName:  validator_cookbooks.EphemeryValidatorClusterClassName,
	Tag:               "latest",
}

var web3SignerChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/web3signers/infra/consensys_web3signer",
	DirOut:      "./ethereum/web3signers/infra/processed_consensys_web3signer",
	FnIn:        web3SignerSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
