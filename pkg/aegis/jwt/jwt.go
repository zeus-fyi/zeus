package init_jwt

import (
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func ReplaceToken(p filepaths.Path, token string) error {
	err := p.RemoveFileInPath()
	if err != nil {
		log.Err(err).Msg("error removing jwt token")
		return err
	}
	return SetToken(p, token)
}

func SetToken(p filepaths.Path, token string) error {
	p.DirOut = p.DirIn
	err := p.WriteToFileOutPath([]byte(token))
	if err != nil {
		log.Err(err).Msg("error setting jwt token")
	}
	return err
}

func SetTokenToDefault(p filepaths.Path, tokenFileName, defaultToken string) error {
	p.FnIn = tokenFileName
	p.FnOut = tokenFileName
	if !p.FileInPathExists() {
		err := SetToken(p, defaultToken)
		return err
	} else {
		return ReplaceToken(p, defaultToken)
	}
}
