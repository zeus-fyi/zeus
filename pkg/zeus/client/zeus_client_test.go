package zeus_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/req_types"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ZeusClientTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient ZeusClient
}

var ctx = context.Background()

// set your own topologyID here
var deployKnsReq = req_types.TopologyDeployRequest{
	TopologyID:    0,
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "",
	Env:           "dev",
}

// DirOut is where it will write a copy of the chart you uploaded, which helps verify the workload is correct
var demoChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/kubernetes_apps/demo",
	DirOut:      "./outputs/demo_read_chart",
	FnIn:        "",
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

func (t *ZeusClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = NewDefaultZeusClient(tc.Bearer)

	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()

	// generates outputs to /test/outputs dir
	// uses inputs from /test/mocks dir
}

func TestZeusClientTestSuite(t *testing.T) {
	suite.Run(t, new(ZeusClientTestSuite))
}
