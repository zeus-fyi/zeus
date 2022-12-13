package config_fetching

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type ArtemisConfigFetchTestSuite struct {
	test_suites.BaseTestSuite
	TestClient resty_base.Resty
}

func (t *ArtemisConfigFetchTestSuite) TestDownloadExtract() {
	dataDir.DirIn = "."
	dataDir.DirOut = "."
	ExtractAndDecEphemeralTestnetConfig()
}

func (t *ArtemisConfigFetchTestSuite) TestGetConfig() {
	r, err := getLatestTestnetDataReleaseNumber()
	t.Require().Nil(err)
	t.Assert().NotEmpty(r)
	fmt.Println(r)

	url := GetLatestReleaseConfigDownloadURL()
	t.Require().NotEmpty(url)
	fmt.Println(url)

}

func (t *ArtemisConfigFetchTestSuite) SetupTest() {
	// points dir to test/configs
	//tc := configs.InitLocalTestConfigs()
	t.TestClient = resty_base.GetBaseRestyTestClient("", "")
}

func TestArtemisConfigFetchTestSuite(t *testing.T) {
	suite.Run(t, new(ArtemisConfigFetchTestSuite))
}
