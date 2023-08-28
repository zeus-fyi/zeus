package snapshot_init

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type SnapshotStartupTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *SnapshotStartupTestSuite) SetupTest() {
}

func TestSnapshotStartupTestSuite(t *testing.T) {
	suite.Run(t, new(SnapshotStartupTestSuite))
}
