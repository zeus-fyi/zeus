package compression

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (c *Compression) GzipCompressDir(p *filepaths.Path) error {
	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	if p == nil {
		return errors.New("need to include a path")
	}
	p.FnOut = p.FnIn + ".tar.gz"
	err := os.MkdirAll(filepath.Dir(p.FileOutPath()), 0755)
	if err != nil {
		return err
	}

	out, err := os.Create(p.FileOutPath())
	if err != nil {
		return err
	}
	defer out.Close()

	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
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

func addToArchive(p *filepaths.Path, tw *tar.Writer, filename string) error {
	// Open the file which will be written into the archive
	file, err := os.Open(p.DirIn + string(filepath.Separator) + filename)
	if err != nil {
		log.Err(err).Msg("Compression: addToArchive,  os.Open(p.DirIn + string(filepath.Separator) + filename)")
		return err
	}
	defer file.Close()

	// Get FileInfo about our file providing file size, mode, etc.
	info, err := file.Stat()
	if err != nil {
		log.Err(err).Msg("Compression: addToArchive, file.Stat()")
		return err
	}

	// Create a tar Header from the FileInfo data
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		log.Err(err).Msg("Compression: addToArchive, tar.FileInfoHeader(info, info.Name())")
		return err
	}

	// Use full path as name (FileInfoHeader only takes the basename)
	// If we don't do this the directory strucuture would
	// not be preserved
	// https://golang.org/src/archive/tar/common.go?#L626
	header.Name = filename
	header.Size = info.Size()
	// Write file header to the tar archive
	err = tw.WriteHeader(header)
	if err != nil {
		log.Err(err).Msg("Compression: addToArchive, WriteHeader(header)")
		return err
	}

	// Copy file content to tar archive
	_, err = io.CopyN(tw, file, info.Size())
	if err != nil {
		log.Info().Int64("filesize", info.Size())
		log.Err(err).Msg("Compression: addToArchive, io.Copy(tw, file)")
		return err
	}
	return nil
}
