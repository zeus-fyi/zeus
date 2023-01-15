package age_encryption

import (
	"errors"
	"io"
	"os"

	"filippo.io/age"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (a *Age) Encrypt(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	recipient, err := age.ParseX25519Recipient(a.agePublicKey)
	if err != nil {
		return err
	}
	p.FnOut = p.FnIn + ".age"
	outFile, err := os.Create(p.FnOut)
	if err != nil {
		return err
	}
	defer outFile.Close()
	bytesToEncrypt, err := os.ReadFile(p.FileDirOutFnInPath())
	if err != nil {
		return err
	}
	w, err := age.Encrypt(outFile, recipient)
	if err != nil {
		return err
	}
	if _, werr := w.Write(bytesToEncrypt); werr != nil {
		return werr
	}
	_, err = io.Copy(w, outFile)
	if err != nil {
		return err
	}
	if cerr := w.Close(); cerr != nil {
		return cerr
	}
	p.FnIn = p.FnOut
	return err
}
