package compression

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/fs"

	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func GzipDirectoryToMemoryFS(p filepaths.Path, inMemfs memfs.MemFS) ([]byte, error) {
	// Create in-memory file system and populate it with files from directoryPath
	// Create gzip writer on top of output file in memory
	gzipBytes := new(bytes.Buffer)
	gzipWriter := gzip.NewWriter(gzipBytes)

	// Create tar writer on top of gzip writer
	tarWriter := tar.NewWriter(gzipWriter)

	dir, err := inMemfs.Sub(p.DirIn)
	if err != nil {
		return nil, err
	}
	// Iterate over files in directory and add to tar archive
	err = fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Create header for current file
		if d.IsDir() || path == p.FnOut {
			return nil
		}
		p.FnIn = path
		file, err := inMemfs.FS.Open(p.FileInPath())
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		header := new(tar.Header)
		header.Name = fileInfo.Name()
		header.Mode = int64(fileInfo.Mode())
		header.Size = fileInfo.Size()
		header.ModTime = fileInfo.ModTime()

		// Write header to tar archive
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		// Copy file contents to tar archive
		_, err = io.Copy(tarWriter, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Flush and close tar writer, then gzip writer
	err = tarWriter.Close()
	if err != nil {
		return nil, err
	}
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	// Return compressed bytes
	return gzipBytes.Bytes(), nil
}
