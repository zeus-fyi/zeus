package snapshot_init

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiStartupTestSuite struct {
	SnapshotStartupTestSuite
}

func (t *SuiStartupTestSuite) TestGenesisBlobDownloader() {
	// TODO
}

func (t *SuiStartupTestSuite) TestS3SnapshotDownloader() {
	// TODO
}

func TestSuiStartupTestSuite(t *testing.T) {
	suite.Run(t, new(SuiStartupTestSuite))
}
