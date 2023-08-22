package iris_operators

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

func (t *IrisOpsTestSuite) TestOperation() {
	val, ok := ConvertToFloat64("1.0")
	t.True(ok)
	t.Equal(1.0, val)
}

type IrisOpsTestSuite struct {
	test_suites.BaseTestSuite
	BearerToken string
}

func (t *IrisOpsTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.BearerToken = tc.QuickNodeIrisToken

	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestIrisOpsTestSuite(t *testing.T) {
	suite.Run(t, new(IrisOpsTestSuite))
}
