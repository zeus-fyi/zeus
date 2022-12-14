package beacon_cookbooks

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io/fs"
	"os"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func SetCustomConfigOpts(p *filepaths.Path) error {
	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	if p == nil {
		return errors.New("need to include a path")
	}
	p.FnOut = p.FnIn
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
			// override
			//aerr := addToArchive(p, tw, path)
			//if aerr != nil {
			//	return aerr
			//}
		}
		return nil
	})

	p.DirIn = p.DirOut
	p.FnIn = p.FnOut
	return err
}
