package hestia_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type HestiaClientTestSuite struct {
	test_suites.BaseTestSuite
	HestiaTestClient Hestia

	TestAccount1 ecdsa.Account
}

func (t *HestiaClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.HestiaTestClient = NewDefaultHestiaClient(tc.Bearer)
	// t.HestiaTestClient = NewLocalHestiaClient(tc.Bearer)
	// points working dir to inside /test
	pkHexString := tc.LocalEcsdaTestPkey
	t.TestAccount1 = ecdsa.NewAccount(pkHexString)
}

func TestHestiaClientTestSuite(t *testing.T) {
	suite.Run(t, new(HestiaClientTestSuite))
}
