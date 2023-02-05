package serverless_inmemdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type ServerlessInMemDBsTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (s *ServerlessInMemDBsTestSuite) TestInMemDBsImport() {
	// change to actual path
	// KeystorePath.DirIn = "keystores"
	ImportIntoInMemDB(ctx, s.Tc.HDWalletPassword)
}
func TestInMemDBsTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessInMemDBsTestSuite))
}
