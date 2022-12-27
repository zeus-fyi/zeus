package hera_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type HeraClientTestSuite struct {
	test_suites.BaseTestSuite
	HeraTestClient HeraClient
}

var ctx = context.Background()

var demoChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/hera",
	DirOut:      "./mocks/outputs",
	FnIn:        "demo", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

func (t *HeraClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.HeraTestClient = NewLocalHeraClient(tc.Bearer)

	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()

	// generates outputs to /test/outputs dir
	// uses inputs from /test/mocks dir
}

func TestHeraClientTestSuite(t *testing.T) {
	suite.Run(t, new(HeraClientTestSuite))
}
