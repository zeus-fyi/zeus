package beacon_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// DeployConsensusClientKnsReq set your own topologyID here after uploading a chart workload
var DeployConsensusClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
}

var ConsensusClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "lighthouseHercules",
	ChartName:         "lighthouseHercules",
	ChartDescription:  "lighthouseHercules",
	Version:           fmt.Sprintf("lighthouseHerculesv0.0.%d", time.Now().Unix()),
	ClusterClassName:  "ethereumBeacons",
	ComponentBaseName: "consensusClient",
	SkeletonBaseName:  "lighthouseHercules",
	Tag:               "latest",
}

var BeaconConsensusClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/consensus_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "lighthouseHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
