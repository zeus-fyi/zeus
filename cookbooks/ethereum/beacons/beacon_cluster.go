package beacon_cookbooks

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"

var cd = zeus_req_types.ClusterTopologyDeployRequest{
	// ethereumBeacons is a reserved keyword, to make it global, you can replace the below with your own setup
	ClusterName: "ethereumBeacons",
	BaseOptions: []string{"gethHercules", "lighthouseHercules"},
	CloudCtxNs:  BeaconCloudCtxNs,
}
