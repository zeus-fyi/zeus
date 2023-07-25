package iris_proxy_rules_configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type IrisConfigTestSuite struct {
	test_suites.BaseTestSuite
	IrisClient Iris
}

func (t *IrisConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.IrisClient = NewIrisClient(tc.Bearer)
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestIrisConfigTestSuite(t *testing.T) {
	suite.Run(t, new(IrisConfigTestSuite))
}
