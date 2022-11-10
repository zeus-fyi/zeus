package compression

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

// UnGzip takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func (c *Compression) UnGzip(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}

	r, err := os.Open(p.FileInPath())
	if err != nil {
		return err
	}
	defer r.Close()
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, herr := tr.Next()

		switch {

		// if no more files are found return
		case herr == io.EOF:
			return nil

		// return any other error
		case herr != nil:
			return herr

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// check the file type
		switch header.Typeflag {

		// if it's a dir do nothing
		case tar.TypeDir:

		// if it's a file create it
		case tar.TypeReg:
			p.FnOut = header.Name

			fo := p.FileOutPath()
			dir := path.Dir(fo)
			if _, zerr := os.Stat(dir); os.IsNotExist(zerr) {
				_ = os.MkdirAll(dir, 0700) // Create your dir
			}
			outFile, perr := os.Create(fo)
			if perr != nil {
				return perr
			}
			if _, cerr := io.Copy(outFile, tr); cerr != nil {
				return cerr
			}
			outFile.Close()
		}
	}
}
