package age_encryption

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	"io"
	"os"

	"filippo.io/age"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (a *Age) EncryptFromInMemFS(inMemFs memfs.MemFS, p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	recipient, err := age.ParseX25519Recipient(a.agePublicKey)
	if err != nil {
		log.Err(err)
		return err
	}
	p.FnOut = p.FnIn + ".age"
	outFile, err := os.Create(p.FileOutPath())
	if err != nil {
		log.Err(err)
		return err
	}
	defer outFile.Close()
	bytesToEncrypt, err := inMemFs.ReadFile(p.FileInPath())
	if err != nil {
		log.Err(err)
		return err
	}
	w, err := age.Encrypt(outFile, recipient)
	if err != nil {
		log.Err(err)
		return err
	}
	if _, werr := w.Write(bytesToEncrypt); werr != nil {
		log.Err(werr)
		return werr
	}
	_, err = io.Copy(w, outFile)
	if err != nil {
		log.Err(err)
		return err
	}
	if cerr := w.Close(); cerr != nil {
		log.Err(cerr)
		return cerr
	}
	p.FnIn = p.FnOut
	return err
}

func (a *Age) Encrypt(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	recipient, err := age.ParseX25519Recipient(a.agePublicKey)
	if err != nil {
		log.Err(err)
		return err
	}
	p.FnOut = p.FnIn + ".age"
	outFile, err := os.Create(p.FileOutPath())
	if err != nil {
		log.Err(err)
		return err
	}
	defer outFile.Close()
	bytesToEncrypt, err := os.ReadFile(p.FileDirOutFnInPath())
	if err != nil {
		log.Err(err)
		return err
	}
	w, err := age.Encrypt(outFile, recipient)
	if err != nil {
		log.Err(err)
		return err
	}
	if _, werr := w.Write(bytesToEncrypt); werr != nil {
		log.Err(werr)
		return werr
	}
	_, err = io.Copy(w, outFile)
	if err != nil {
		log.Err(err)
		return err
	}
	if cerr := w.Close(); cerr != nil {
		log.Err(cerr)
		return cerr
	}
	p.FnIn = p.FnOut
	return err
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
