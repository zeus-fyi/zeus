package web3signer_cookbooks

import (
	"fmt"
	"time"

	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

var (
	Web3SignerComponentBaseName = "web3Signer"
	Web3SignerSkeletonBaseName  = "web3Signer"
)

var Web3SignerClusterDefinition = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName: validator_cookbooks.EphemeryValidatorClusterClassName,
	SkeletonBaseOptions: []string{
		validator_cookbooks.ExecSkeletonBase,
		validator_cookbooks.ConsensusSkeletonBase,
		validator_cookbooks.ValidatorSkeletonBaseName,
		validator_cookbooks.ChoreographySkeletonBase,
		Web3SignerSkeletonBaseName,
	},
	CloudCtxNs: validator_cookbooks.ValidatorCloudCtxNs,
}

var DeployWeb3SignerKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: validator_cookbooks.ValidatorCloudCtxNs,
}

var Web3SignerChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      Web3SignerSkeletonBaseName,
	ChartName:         Web3SignerSkeletonBaseName,
	ChartDescription:  Web3SignerSkeletonBaseName,
	Version:           fmt.Sprintf("web3SignerBase-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  Web3SignerSkeletonBaseName,
	ComponentBaseName: Web3SignerComponentBaseName,
	ClusterClassName:  validator_cookbooks.EphemeryValidatorClusterClassName,
	Tag:               "latest",
}

var Web3SignerChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/web3signers/infra/consensys_web3signer",
	DirOut:      "./ethereum/web3signers/infra/processed_consensys_web3signer",
	FnIn:        Web3SignerSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
