package filepaths

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func (p *Path) WipeDirIn() error {
	return RemoveContents(p.DirIn)
}

func (p *Path) WipeDirOut() error {
	return RemoveContents(p.DirOut)
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		log.Err(err).Interface("dir", dir)
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Err(err).Interface("dir", dir)
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			log.Err(err).Interface("dir", dir)
			return err
		}
	}
	return nil
}
