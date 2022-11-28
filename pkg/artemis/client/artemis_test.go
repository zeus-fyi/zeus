package artemis_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type ArtemisClientTestSuite struct {
	test_suites.BaseTestSuite
	ArtemisTestClient ArtemisClient

	TestAccount1 ecdsa.Account
	TestAccount2 ecdsa.Account
}

func (t *ArtemisClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.ArtemisTestClient = NewDefaultArtemisClient(tc.Bearer)
	//t.ArtemisTestClient = NewLocalArtemisClient(tc.Bearer)
	// points working dir to inside /test

	pkHexString := tc.LocalEcsdaTestPkey
	t.TestAccount1 = ecdsa.NewAccount(pkHexString)
	t.ArtemisTestClient.Account = t.TestAccount1

	pkHexString2 := tc.LocalEcsdaTestPkey2
	t.TestAccount2 = ecdsa.NewAccount(pkHexString2)
}

func TestArtemisClientTestSuite(t *testing.T) {
	suite.Run(t, new(ArtemisClientTestSuite))
}
