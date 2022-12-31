package web3signer_cookbooks

import (
	"fmt"
	"time"

	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

var (
	web3SignerComponentBaseName = "web3Signer"
	web3SignerSkeletonBaseName  = "web3Signer"
)

var cd = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName:    validator_cookbooks.EphemeryValidatorClusterClassName,
	SkeletonBaseOptions: []string{web3SignerSkeletonBaseName},
	CloudCtxNs:          Web3SignerCloudCtxNs,
}

var DeployWeb3SignerKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: Web3SignerCloudCtxNs,
}

var Web3SignerCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "ephemery-staking", // set with your own namespace
	Env:           "production",
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
	DirIn:       "./ethereum/web3signer/infra/consensys_web3signer",
	DirOut:      "./ethereum/web3signer/infra/processed_consensys_web3signer",
	FnIn:        web3SignerSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}
