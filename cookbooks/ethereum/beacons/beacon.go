package beacon_cookbook

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
)

var beaconCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "beacon", // set with your own namespace
	Env:           "production",
}
