package ephemery_reset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ResetTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *ResetTestSuite) SetupTest() {
}

func (t *ResetTestSuite) TestReset() {
	dd := filepaths.Path{}
	ExtractAndDecEphemeralTestnetConfig(dd, "test")
	kt := ExtractResetTime("./data/testnet/retention.vars")
	t.Assert().NotEmpty(kt)
	fmt.Println(kt)
}

func (t *ResetTestSuite) TestResetExtract() {
	rl, err := getLatestTestnetDataReleaseNumber()
	t.Assert().NoError(err)
	t.Assert().NotEmpty(rl)

	urlPath := GetLatestReleaseConfigDownloadURL()
	fmt.Println(urlPath)
}

func TestResetTestSuite(t *testing.T) {
	suite.Run(t, new(ResetTestSuite))
}
