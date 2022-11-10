package compression

import (
	"archive/tar"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (c *Compression) TarFolder(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}

	files, err := ioutil.ReadDir(p.DirIn)
	if err != nil {
		return err
	}
	tarfile, err := os.Create(p.Fn)
	if err != nil {
		return err
	}
	defer tarfile.Close()
	var fileW io.WriteCloser = tarfile
	//Tar file writer
	tarfileW := tar.NewWriter(fileW)
	defer tarfileW.Close()

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}
		file, ferr := os.Open(p.DirIn + string(filepath.Separator) + fileInfo.Name())
		if ferr != nil {
			return ferr
		}
		defer file.Close()
		header := new(tar.Header)
		header.Name = file.Name()
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()
		err = tarfileW.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tarfileW, file)
		if err != nil {
			return err
		}
	}
	return err
}
