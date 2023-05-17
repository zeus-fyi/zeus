package poseidon

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

var ctx = context.Background()

type ChainDownloadsClientTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *ChainDownloadsClientTestSuite) SetupTest() {
}

func (t *ChainDownloadsClientTestSuite) TestDownloader() {
	cfgs := configs.InitLocalTestConfigs()
	err := DownloadFile(ctx, ".", cfgs.PresignedBucketURL)
	t.Assert().Nil(err)
}

func TestChainDownloadsClientTestSuite(t *testing.T) {
	suite.Run(t, new(ChainDownloadsClientTestSuite))
}
