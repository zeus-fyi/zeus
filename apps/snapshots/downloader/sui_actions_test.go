package snapshot_init

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type SuiStartupTestSuite struct {
	SnapshotStartupTestSuite
}

func (t *SuiStartupTestSuite) TestGenesisBlobDownloader() {
	forceDirToTestDirLocation()
	blobUrl := "https://github.com/MystenLabs/sui-genesis/raw/main/devnet/genesis.blob"
	w := WorkloadInfo{
		WorkloadType: "full",
		Protocol:     "sui",
		Network:      "devnet",
		DataDir: filepaths.Path{
			DirIn: ".",
		},
	}

	err := DownloadGenesisBlob(w, blobUrl)
	t.Require().Nil(err)
}

func (t *SuiStartupTestSuite) TestS3SnapshotDownloader() {
	forceDirToTestDirLocation()
	w := WorkloadInfo{
		WorkloadType: "full",
		Protocol:     "sui",
		Network:      "testnet",
		DataDir: filepaths.Path{
			DirIn: "/tmp",
		},
	}

	err := SuiDownloadSnapshotS3(w)
	t.Require().Nil(err)
}

func TestSuiStartupTestSuite(t *testing.T) {
	suite.Run(t, new(SuiStartupTestSuite))
}

func forceDirToTestDirLocation() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}
