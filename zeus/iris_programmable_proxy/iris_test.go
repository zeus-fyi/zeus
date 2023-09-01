package iris_programmable_proxy

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

type IrisConfigTestSuite struct {
	test_suites.BaseTestSuite
	IrisClient     Iris
	IrisClientProd Iris
	BearerToken    string
}

func (t *IrisConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.BearerToken = tc.Bearer
	t.IrisClient = Iris{
		resty_base.GetBaseRestyClient("http://localhost:8080", tc.Bearer),
	}
	t.IrisClientProd = Iris{
		resty_base.GetBaseRestyClient("https://iris.zeus.fyi", tc.Bearer),
	}
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestIrisConfigTestSuite(t *testing.T) {
	suite.Run(t, new(IrisConfigTestSuite))
}
