package hercules_jwt

import (
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func SetToken(p filepaths.Path, token string) error {
	p.DirOut = p.DirIn
	err := p.WriteFileInPath([]byte(token))
	if err != nil {
		log.Err(err).Msg("error setting jwt token")
	}
	return err
}

func CheckIfJwtTokenExistsElseWriteDefault(p filepaths.Path, tokenFileName, defaultToken string) error {
	p.FnIn = tokenFileName
	p.FnOut = tokenFileName

	if !p.FileInPathExists() {
		err := SetToken(p, defaultToken)
		return err
	}
	return nil
}
