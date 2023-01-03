package filepaths

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func (p *Path) ReadFileInPath() []byte {
	byteArray, err := os.ReadFile(p.FileInPath())
	if err != nil {
		panic(err)
	}
	return byteArray
}

func (p *Path) ReadFirstFileInPathWithFilter() ([]byte, error) {
	var b []byte
	fileSystem := os.DirFS(p.DirIn)
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if (p.FilterFiles != nil && strings_filter.FilterStringWithOpts(path, p.FilterFiles)) || p.FnIn == path {
			filePath := pathJoin(p.DirIn, path)
			byteArray, rerr := os.ReadFile(filePath)
			if rerr != nil {
				panic(rerr)
			}
			if byteArray != nil && b == nil {
				b = byteArray
				return nil
			}
		}
		return err
	})
	return b, err
}

func (p *Path) WalkAndApplyFuncToFileType(ext string, f func(p string) error) error {
	fileSystem := os.DirFS(p.DirIn)
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (filepath.Ext(d.Name()) == ext || ext == "") {
			if strings_filter.FilterStringWithOpts(path, p.FilterFiles) {
				filePath := pathJoin(p.DirIn, path)
				return f(filePath)
			}
		}
		return nil
	})
	return err
}

func pathJoin(root, file string) string {
	return path.Join(root, file)
}
