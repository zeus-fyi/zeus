package beacon_actions

import (
	"context"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

var basePar = zeus_pods_reqs.PodActionRequest{
	TopologyDeployRequest: BeaconKnsReq,
	PodName:               "",
	FilterOpts:            nil,
	ClientReq:             nil,
	DeleteOpts:            nil,
}

// set your own topologyID here after uploading a chart workload
var BeaconKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 1669159384971627008,
	CloudCtxNs: beaconCloudCtxNs,
}

var beaconCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "beacon", // set with your own namespace
	Env:           "production",
}

type BeaconActionsTestSuite struct {
	test_suites.BaseTestSuite
	BeaconActionsClient
}

func (t *BeaconActionsTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.BeaconActionsClient = NewDefaultBeaconActionsClient(tc.Bearer, BeaconKnsReq)
	dir := cookbooks.ChangeToCookbookDir()

	t.BeaconActionsClient.PrintPath.DirIn = path.Join(dir, "/ethereum/beacons/logs")
	t.BeaconActionsClient.PrintPath.DirOut = path.Join(dir, "/ethereum/outputs")
	t.BeaconActionsClient.ConfigPaths.DirIn = "./ethereum/beacons/infra"
	t.BeaconActionsClient.ConfigPaths.DirOut = "./ethereum/outputs"
}

func TestBeaconActionsTestSuite(t *testing.T) {
	suite.Run(t, new(BeaconActionsTestSuite))
}
