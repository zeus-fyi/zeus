package sui_cookbooks

import (
	"fmt"
)

func (t *SuiCookbookTestSuite) TestConfigDriverBuilder() {
	cfg := SuiConfigOpts{
		DownloadSnapshot: false,
		WithIngress:      false,
		CloudProvider:    "do",
		Network:          testnet,
	}
	cd := GetSuiClientClusterDef(cfg)
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
}
