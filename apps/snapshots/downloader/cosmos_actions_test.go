package snapshot_init

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type CosmosStartupTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *CosmosStartupTestSuite) SetupTest() {
}

func (t *CosmosStartupTestSuite) TestGetMainnetPersistentPeers() {
	peers := GetMainnetPersistentPeers(ctx)
	t.Assert().NotEmpty(peers)
	fmt.Println(peers)
}

func (t *CosmosStartupTestSuite) TestGetStateSyncInfo() {
	testRPC := "https://rpc.sentry-01.theta-testnet.polypore.xyz"
	si := GetStateSyncInfo(ctx, testRPC)
	t.Assert().NotEmpty(si)

	si = GetStateSyncInfoMainnet(ctx)
	t.Assert().NotEmpty(si)
}

func (t *CosmosStartupTestSuite) TestConfigOverride() {
	testRPC := "https://rpc.sentry-01.theta-testnet.polypore.xyz"
	si := GetStateSyncInfo(ctx, testRPC)
	t.Assert().NotEmpty(si)
	err := replaceLineIfStartsWith("./config/config.toml", "enable = false", "enable = true")
	t.Assert().NoError(err)
	err = replaceLineIfStartsWith("./config/config.toml", "rpc_servers = \"\"", fmt.Sprintf("rpc_servers = \"%s\"", cosmosTestnetStateSyncRPC))
	t.Assert().NoError(err)
	err = replaceLineIfStartsWith("./config/config.toml", "trust_height = 0", fmt.Sprintf("trust_height = %s", si.TrustHeight))
	t.Assert().NoError(err)
	err = replaceLineIfStartsWith("./config/config.toml", "trust_hash = \"\"", fmt.Sprintf("trust_hash = \"%s\"", si.TrustHash))
	t.Assert().NoError(err)
	err = replaceLineIfStartsWith("./config/config.toml", "trust_period = \"0s\"", fmt.Sprintf("trust_period = \"%s\"", "8h0m0s"))
	t.Assert().NoError(err)
}
func TestCosmosStartupTestSuite(t *testing.T) {
	suite.Run(t, new(CosmosStartupTestSuite))
}
