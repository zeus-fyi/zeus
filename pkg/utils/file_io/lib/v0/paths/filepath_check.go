package filepaths

import (
	"os"

	"github.com/rs/zerolog/log"
)

func (p *Path) FileInPathExists() bool {
	return doesFileExist(p.FileInPath())
}

func doesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)

	// check if error is "file not exists"
	if os.IsNotExist(err) {
		return false
	} else {
		if err != nil {
			log.Err(err).Msgf("doesFileExist: path %s", filePath)
			return false
		}
		return true
	}
}
