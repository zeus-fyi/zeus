package hercules_jwt

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func SetToken(p filepaths.Path, token string) error {

	err := p.WriteFileInPath([]byte(token))
	return err
}
