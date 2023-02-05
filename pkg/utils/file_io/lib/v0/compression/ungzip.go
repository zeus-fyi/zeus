package compression

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	"io"
	"os"
	"path"

	"github.com/rs/zerolog/log"
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
		log.Err(err).Interface("path", p).Msg("UnGzip: os.Open")
		return err
	}
	defer r.Close()
	gzr, err := gzip.NewReader(r)
	if err != nil {
		log.Err(err)
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
				log.Err(perr).Interface("path", p).Msg("UnGzip: os.Create(fo)")
				return perr
			}
			if _, cerr := io.Copy(outFile, tr); cerr != nil {
				log.Err(cerr).Interface("path", p).Msg("UnGzip: io.Copy")
				return cerr
			}
			outFile.Close()
		}
	}
}

func (c *Compression) UnGzipFromInMemFsOutToInMemFS(p *filepaths.Path, fs memfs.MemFS) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	r, err := fs.ReadFileInPath(p)
	if err != nil {
		log.Err(err).Msgf("Compression: UnGzipFromInMemFsOutToInMemFS, fs.ReadFileInPath(p) %s", p.FileInPath())
		return err
	}
	in := &bytes.Buffer{}
	_, err = in.Write(r)
	if err != nil {
		return err
	}
	gzr, err := gzip.NewReader(in)
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
			p.FnIn = header.Name

			fo := p.FileDirOutFnInPath()
			p.DirIn = path.Dir(fo)

			out := &bytes.Buffer{}
			if _, cerr := io.Copy(out, tr); cerr != nil {
				log.Err(err).Msg("Compression: UnGzipFromInMemFsOutToInMemFS, io.Copy(out, tr)")
				return cerr
			}
			ferr := fs.MakeFileDirOutFnInPath(p, out.Bytes())
			if ferr != nil {
				log.Err(err).Msg("Compression: UnGzipFromInMemFsOutToInMemFS, fs.MakeFile(p, out.Bytes())")
				return ferr
			}
		}
	}
}
