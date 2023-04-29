package snapshot_init

import (
	"context"
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

func (t *CosmosStartupTestSuite) TestGetTestnetBlockHeight() {
	testRPC := "https://rpc.sentry-01.theta-testnet.polypore.xyz"
	si := GetStateSyncInfo(ctx, testRPC)
	t.Assert().NotEmpty(si)
}

func TestCosmosStartupTestSuite(t *testing.T) {
	suite.Run(t, new(CosmosStartupTestSuite))
}
