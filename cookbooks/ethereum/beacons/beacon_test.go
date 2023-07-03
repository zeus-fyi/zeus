package ethereum_beacon_cookbooks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
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
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	const poseidonEndpoint = "https://poseidon.zeus.fyi"

	//var preSignedURL string
	//rc := resty_base.GetBaseRestyClient(poseidonEndpoint, tc.Bearer)
	//resp, err := rc.R().SetResult(&preSignedURL).Get("/v1/ethereum/mainnet/geth")
	//t.Require().Nil(err)
	//fmt.Println(resp)
	cookbooks.ChangeToCookbookDir()
}

func (t *BeaconCookbookTestSuite) TestDestroyDeployBeacon() {
	resp, err := t.ZeusTestClient.DestroyDeploy(context.Background(), DeployExecClientKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func TestBeaconCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(BeaconCookbookTestSuite))
}
