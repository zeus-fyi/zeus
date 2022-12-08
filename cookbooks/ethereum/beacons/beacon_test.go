package beacon_cookbook

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type BeaconCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *BeaconCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewLocalZeusClient(tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func (t *BeaconCookbookTestSuite) TestDestroyDeployBeacon() {
	resp, err := t.ZeusTestClient.DestroyDeploy(context.Background(), deployExecClientKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func TestBeaconCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(BeaconCookbookTestSuite))
}
