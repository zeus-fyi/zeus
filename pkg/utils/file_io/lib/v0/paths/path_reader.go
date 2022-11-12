package filepaths

import (
	"os"
)

func (p *Path) ReadFileInPath() []byte {
	byteArray, err := os.ReadFile(p.FileInPath())
	if err != nil {
		panic(err)
	}
	return byteArray
}
