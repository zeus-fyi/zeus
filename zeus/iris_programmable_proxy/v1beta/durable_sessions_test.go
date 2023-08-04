package iris_programmable_proxy_v1_beta

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"github.com/zeus-fyi/zeus/zeus/iris_programmable_proxy"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

func (t *IrisConfigTestSuite) TestDurableSessions() {

}

type IrisConfigTestSuite struct {
	test_suites.BaseTestSuite
	IrisClient     iris_programmable_proxy.Iris
	IrisClientProd iris_programmable_proxy.Iris
}

func (t *IrisConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.IrisClient = iris_programmable_proxy.Iris{
		Resty: resty_base.GetBaseRestyClient("http://localhost:8080", tc.Bearer),
	}
	t.IrisClientProd = iris_programmable_proxy.Iris{
		Resty: resty_base.GetBaseRestyClient("https://iris.zeus.fyi", tc.Bearer),
	}
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestIrisConfigTestSuite(t *testing.T) {
	suite.Run(t, new(IrisConfigTestSuite))
}
