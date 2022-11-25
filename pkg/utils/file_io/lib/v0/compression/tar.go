package compression

import (
	"archive/tar"
	"errors"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (c *Compression) TarCompress(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	p.FnOut = p.FnIn + ".tar"
	out, err := os.Create(p.FileOutPath())
	if err != nil {
		log.Err(err).Msg("Compression: TarCompress, os.Create(p.FileOutPath()")
		return err
	}
	defer out.Close()

	tw := tar.NewWriter(out)
	defer tw.Close()
	fileSystem := os.DirFS(p.DirIn)
	err = fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && path != p.FnOut {
			aerr := addToArchive(p, tw, path)
			if aerr != nil {
				return aerr
			}
		}
		return nil
	})

	p.DirIn = p.DirOut
	p.FnIn = p.FnOut
	return err
}

func tarReader(p *filepaths.Path, r io.Reader) error {
	tr := tar.NewReader(r)
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
			dir := path.Dir(fo)
			if _, zerr := os.Stat(dir); os.IsNotExist(zerr) {
				_ = os.MkdirAll(dir, 0700) // Create your dir
			}
			outFile, perr := os.Create(fo)
			if perr != nil {
				log.Err(perr).Msg("Compression: tarReader, os.Create(fo)")
				return perr
			}
			if _, cerr := io.Copy(outFile, tr); cerr != nil {
				log.Err(cerr).Msg("Compression: tarReader, io.Copy(outFile, tr)")
				outFile.Close()
				return cerr
			}
			outFile.Close()
		}
	}
}
