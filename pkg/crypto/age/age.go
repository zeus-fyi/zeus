package age_encryption

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type Age struct {
	agePrivateKey string
	agePublicKey  string
}

func NewAge(privKey, pubKey string) Age {
	a := Age{
		agePrivateKey: privKey,
		agePublicKey:  pubKey,
	}
	return a
}

/* GzipAndEncrypt, use this format
p := filepaths.Path{
		DirIn:       "./secrets",
		FnIn:        "secrets",
	}
*/

func (a *Age) GzipAndEncrypt(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	c := compression.NewCompression()

	err := c.GzipCompressDir(p)
	if err != nil {
		log.Err(err)
		return err
	}

	err = a.Encrypt(p)
	if err != nil {
		log.Err(err)
		return err
	}
	return err
}
