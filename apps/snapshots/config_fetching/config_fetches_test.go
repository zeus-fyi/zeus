package config_fetching

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ConfigFetchTestSuite struct {
	test_suites.BaseTestSuite
	TestClient resty_base.Resty
}

func (t *ConfigFetchTestSuite) TestDownloadExtract() {
	dd := filepaths.Path{}
	ExtractAndDecEphemeralTestnetConfig(dd, "test")

	cmd := exec.Command("geth", "--datadir", "./configs", "init", "./configs/genesis.json")
	err := cmd.Run()
	t.Require().Nil(err)
}

func (t *ConfigFetchTestSuite) TestGetConfig() {
	r, err := getLatestTestnetDataReleaseNumber()
	t.Require().Nil(err)
	t.Assert().NotEmpty(r)
	fmt.Println(r)

	url := GetLatestReleaseConfigDownloadURL()
	t.Require().NotEmpty(url)
	fmt.Println(url)
}

func (t *ConfigFetchTestSuite) SetupTest() {
	// points dir to test/configs
	//tc := configs.InitLocalTestConfigs()
	t.TestClient = resty_base.GetBaseRestyTestClient("", "")
}

func TestConfigFetchTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigFetchTestSuite))
}
