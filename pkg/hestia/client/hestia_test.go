package hestia_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	zeus_ecdsa "github.com/zeus-fyi/zeus/pkg/aegis/crypto/ecdsa"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type HestiaClientTestSuite struct {
	test_suites.BaseTestSuite
	HestiaTestClient Hestia
	TestAccount1     zeus_ecdsa.Account
}

func (t *HestiaClientTestSuite) SetupTest() {
	// points dir to test/configs
	t.Tc = configs.InitLocalTestConfigs()
	t.HestiaTestClient = NewDefaultHestiaClient(t.Tc.Bearer)
	//t.HestiaTestClient = NewLocalHestiaClient(t.Tc.Bearer)
	// points working dir to inside /test
	pkHexString := t.Tc.LocalEcsdaTestPkey
	t.TestAccount1 = zeus_ecdsa.NewAccount(pkHexString)
}

func TestHestiaClientTestSuite(t *testing.T) {
	suite.Run(t, new(HestiaClientTestSuite))
}
