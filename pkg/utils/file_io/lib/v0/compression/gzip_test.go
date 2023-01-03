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
	c.ChangeToTestDir()
}

func (c *CompressionTestSuite) TestTarGzip() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "./mocks/kubernetes_apps/demo",
		DirOut:      "./outputs/compression/gzip",
		FnIn:        "demo",
		Env:         "",
		FilterFiles: &strings_filter.FilterOpts{},
	}

	err := c.Comp.GzipCompressDir(&p)
	c.Require().Nil(err)
}

func (c *CompressionTestSuite) TestUnGzip() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "./outputs/compression/gzip",
		DirOut:      "./outputs/compression/ungzip",
		FnIn:        "demo.tar.gz",
		FnOut:       "",
		Env:         "",
		FilterFiles: &strings_filter.FilterOpts{},
	}
	err := c.Comp.UnGzip(&p)
	c.Require().Nil(err)
}

func (c *CompressionTestSuite) TestLz4Dec() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "/Users/alex/go/Olympus/Zeus/pkg/utils/file_io/lib/v0/compression/",
		DirOut:      "/Users/alex/go/Olympus/Zeus/pkg/utils/file_io/lib/v0/compression",
		FnIn:        "geth.tar.lz4",
		FnOut:       "",
		Env:         "",
		FilterFiles: &strings_filter.FilterOpts{},
	}
	err := c.Comp.Lz4Decompress(&p)
	c.Require().Nil(err)
}

func TestCompressionTestSuite(t *testing.T) {
	suite.Run(t, new(CompressionTestSuite))
}
