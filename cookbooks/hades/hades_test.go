package cookbooks_hades

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/client"
)

type HadesCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *HadesCookbookTestSuite) TestClusterSetup() {
	gcd := HadesClusterDefinition.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := HadesClusterDefinition.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := HadesClusterDefinition.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
}

func (t *HadesCookbookTestSuite) SetupTest() {
	cookbooks.ChangeToCookbookDir()
}

func TestHadesCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(HadesCookbookTestSuite))
}
