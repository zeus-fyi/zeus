package filepaths

import (
	"testing"

	"github.com/stretchr/testify/suite"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type FilePathsTestSuite struct {
	suite.Suite
}

func (s *FilePathsTestSuite) TestFilePathExists() {
	p := Path{
		PackageName: "",
		DirIn:       "./",
		DirOut:      "./",
		FnIn:        "filepath_check.go",
		FnOut:       "",
		Env:         "",
		FilterFiles: strings_filter.FilterOpts{},
	}
	s.Assert().True(p.FileInPathExists())
	p.FnIn = "nofile.go"
	s.Assert().False(p.FileInPathExists())
}

func TestFilePathsTestSuite(t *testing.T) {
	suite.Run(t, new(FilePathsTestSuite))
}
