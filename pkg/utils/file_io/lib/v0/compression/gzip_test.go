package compression

import (
	"testing"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/test/test_suites"

	"github.com/stretchr/testify/suite"
)

type CompressionTestSuite struct {
	test_suites.BaseTestSuite
	Comp Compression
}

func (c *CompressionTestSuite) SetupTest() {
	c.Comp = NewCompression()
}

func (c *CompressionTestSuite) TestTarGzip() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "./.kube",
		DirOut:      "./",
		Fn:          "kube",
		Env:         "",
		FilterFiles: strings_filter.FilterOpts{},
	}

	err := c.Comp.CreateTarGzipArchiveDir(&p)
	c.Require().Nil(err)
}

func (c *CompressionTestSuite) TestUnGzip() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "./",
		DirOut:      "./kube",
		Fn:          "./kube.tar.gz",
		Env:         "",
		FilterFiles: strings_filter.FilterOpts{},
	}

	err := c.Comp.UnGzip(&p)
	c.Require().Nil(err)
}

func (c *CompressionTestSuite) TestTar() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "./.kube",
		DirOut:      "./",
		Fn:          "kube.tar",
		Env:         "",
		FilterFiles: strings_filter.FilterOpts{},
	}

	err := c.Comp.TarFolder(&p)
	c.Require().Nil(err)
}

func TestCompressionTestSuite(t *testing.T) {
	suite.Run(t, new(CompressionTestSuite))
}
