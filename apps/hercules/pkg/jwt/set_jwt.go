package hercules_jwt

import (
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func SetToken(p filepaths.Path, token string) error {
	err := p.WriteFileInPath([]byte(token))
	return err
}

func CheckIfJwtTokenExistsElseWriteDefault(p filepaths.Path, tokenFileName, defaultToken string) error {
	p.FnIn = tokenFileName
	p.FnOut = tokenFileName

	if !p.FileInPathExists() {
		p.FnOut = "jwt.hex"
		err := p.WriteFileInPath([]byte(defaultToken))
		if err != nil {
			log.Err(err).Msg("error setting jwt token")
		}
		return err
	}
	return nil
}
