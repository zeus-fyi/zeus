package ethereum_beacon_cookbooks

import (
	"context"
	"fmt"
)

func (t *BeaconCookbookTestSuite) TestGoerliClusterDeployV2() {
	cd := BeaconGoerliClusterDefinition
	_, err := cd.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestGoerliClusterDestroy() {
	knsReq := DeployConsensusClientKnsReq
	knsReq.CloudCtxNs = BeaconGoerliClusterDefinition.CloudCtxNs
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *BeaconCookbookTestSuite) TestGoerliClusterSetupV2() {
	cd := BeaconGoerliClusterDefinition
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

func (t *BeaconCookbookTestSuite) TestGoerliClusterDefinitionCreationV2() {
	cd := BeaconGoerliClusterDefinition
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}
