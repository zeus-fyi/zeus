package cookbooks_hades

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

type HadesCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

var ctx = context.Background()

func (t *HadesCookbookTestSuite) TestClusterSetup() {
	gcd := HadesClusterDefinition.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	// Create cluster class definitions on Zeus
	err := gcd.CreateClusterClassDefinitions(ctx, t.ZeusTestClient)
	t.Require().Nil(err)

	resp, err := HadesClusterDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *HadesCookbookTestSuite) SetupTest() {
	cookbooks.ChangeToCookbookDir()
}

func TestHadesCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(HadesCookbookTestSuite))
}
