package age_encryption

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	"filippo.io/age"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (a *Age) DecryptToFile(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	outFile, err := os.Create(p.FnOut)
	if err != nil {
		return err
	}
	defer outFile.Close()
	out, err := a.decrypt(p)
	if err != nil {
		return err
	}
	if _, cerr := io.Copy(outFile, out); cerr != nil {
		return cerr
	}
	return err
}

func (a *Age) decrypt(p *filepaths.Path) (*bytes.Buffer, error) {
	out := &bytes.Buffer{}

	if p == nil {
		return out, errors.New("need to include a path")
	}
	identity, err := age.ParseX25519Identity(a.agePrivateKey)
	if err != nil {
		return out, err
	}
	f, err := os.Open(p.FnIn)
	if err != nil {
		return out, err
	}
	defer f.Close()
	r, err := age.Decrypt(f, identity)
	if err != nil {
		return out, err
	}

	p.FnOut, _, _ = strings.Cut(p.FnIn, ".age")
	if _, cerr := io.Copy(out, r); cerr != nil {
		return out, cerr
	}
	p.FnIn = p.FnOut
	return out, err
}
