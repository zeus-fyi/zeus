package filepaths

import (
	"os"
)

func (p *Path) WriteFileInPath(data []byte) error {
	// make path if it doesn't exist
	if _, err := os.Stat(p.FileOutPath()); os.IsNotExist(err) {
		_ = os.MkdirAll(p.DirOut, 0700) // Create your dir
	}
	err := os.WriteFile(p.FileOutPath(), data, 0644)
	return err
}
