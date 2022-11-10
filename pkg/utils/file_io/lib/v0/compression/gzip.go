package compression

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func (c *Compression) CreateTarGzipArchiveDir(p *filepaths.Path) error {
	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	if p == nil {
		return errors.New("need to include a path")
	}
	p.FnOut = p.Fn + ".tar.gz"
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
		if !d.IsDir() {
			aerr := addToArchive(p, tw, path)
			if aerr != nil {
				return aerr
			}
		}
		return nil
	})

	p.DirIn = p.DirOut
	p.Fn = p.FnOut
	return err
}

func addToArchive(p *filepaths.Path, tw *tar.Writer, filename string) error {
	// Open the file which will be written into the archive
	file, err := os.Open(p.DirIn + string(filepath.Separator) + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get FileInfo about our file providing file size, mode, etc.
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create a tar Header from the FileInfo data
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	// Use the full path as name (FileInfoHeader only takes the basename)
	// If we don't do this the directory strucuture would
	// not be preserved
	// https://golang.org/src/archive/tar/common.go?#L626
	header.Name = filename

	// Write file header to the tar archive
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// Copy file content to tar archive
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
