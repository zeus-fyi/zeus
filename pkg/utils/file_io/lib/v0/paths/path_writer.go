package filepaths

import (
	"os"

	"github.com/rs/zerolog/log"
)

func (p *Path) WriteFileInPath(data []byte) error {
	// make path if it doesn't exist
	if _, err := os.Stat(p.FileOutPath()); os.IsNotExist(err) {
		_ = os.MkdirAll(p.DirOut, 0700) // Create your dir
	}
	err := os.WriteFile(p.FileOutPath(), data, 0644)
	return err
}

func (p *Path) RemoveFileInPath() error {
	err := os.Remove(p.FileInPath())
	if err != nil {
		log.Err(err).Msgf("RemoveFileInPath %s", p.FileInPath())
	}
	return err
}
