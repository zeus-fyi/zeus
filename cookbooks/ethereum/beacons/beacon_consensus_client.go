package ethereum_beacon_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

// DeployConsensusClientKnsReq set your own topologyID here after uploading a chart workload
var DeployConsensusClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
}

var ConsensusClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "lighthouse-hercules",
	ChartName:         "lighthouse-hercules",
	ChartDescription:  "lighthouse-hercules",
	Version:           fmt.Sprintf("lighthouse-herculesv0.0.%d", time.Now().Unix()),
	ClusterClassName:  "ethereum-beacons",
	ComponentBaseName: "zeus-consensus-client",
	SkeletonBaseName:  "lighthouse-hercules",
	Tag:               "latest",
}

var BeaconConsensusClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/consensus_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "lighthouse-hercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
