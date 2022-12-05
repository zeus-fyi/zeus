package apollo_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type ApolloClientTestSuite struct {
	test_suites.BaseTestSuite
	ApolloTestClient Apollo
}

func (t *ApolloClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.ApolloTestClient = NewDefaultApolloClient(tc.Bearer)
	// t.ApolloTestClient = NewLocalApolloClient(tc.Bearer)
	// points working dir to inside /test
}

func TestApolloClientTestSuite(t *testing.T) {
	suite.Run(t, new(ApolloClientTestSuite))
}
