package host_info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type HostInfoTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *HostInfoTestSuite) SetupTest() {
}

func (t *HostInfoTestSuite) TestDiskStats() {
	ctx := context.Background()
	diskUsage, err := GetDiskUsageStats(ctx, ".")
	t.Require().Nil(err)
	t.Assert().NotNil(diskUsage)
}

func TestHostInfoTestSuite(t *testing.T) {
	suite.Run(t, new(HostInfoTestSuite))
}
