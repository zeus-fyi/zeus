package beacon_actions

import (
	"context"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

var ctx = context.Background()

// set your own topologyID here after uploading a chart workload
var BeaconKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: beaconCloudCtxNs,
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
