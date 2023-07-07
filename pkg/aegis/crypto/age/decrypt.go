package age_encryption

import (
	"bytes"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	"io"
	"os"
	"strings"

	"filippo.io/age"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (a *Age) Decrypt(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	outFile, err := os.Create(p.FileOutPath())
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
		log.Err(err)
		return out, err
	}
	f, err := os.Open(p.FileInPath())
	if err != nil {
		log.Err(err)
		return out, err
	}
	defer f.Close()
	r, err := age.Decrypt(f, identity)
	if err != nil {
		return out, err
	}

	p.FnOut, _, _ = strings.Cut(p.FnIn, ".age")
	if _, cerr := io.Copy(out, r); cerr != nil {
		log.Err(cerr)
		return out, cerr
	}
	p.FnIn = p.FnOut
	return out, err
}

func (a *Age) DecryptToMemFsFile(p *filepaths.Path, fs memfs.MemFS) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	out, err := a.decryptFromInMemFS(p, fs)
	if err != nil {
		log.Err(err).Msgf("DecryptToMemFsFile, decryptFromInMemFS %s", p.FileInPath())
		return err
	}

	err = fs.MakeFileOut(p, out.Bytes())
	if err != nil {
		log.Err(err).Msgf("DecryptToMemFsFile, MakeFileOut %s", p.FileOutPath())
		return err
	}
	return err
}

func (a *Age) decryptFromInMemFS(p *filepaths.Path, fs memfs.MemFS) (*bytes.Buffer, error) {
	out := &bytes.Buffer{}

	if p == nil {
		return out, errors.New("need to include a path")
	}
	identity, err := age.ParseX25519Identity(a.agePrivateKey)
	if err != nil {
		log.Err(err).Msg("Age, decryptFromInMemFS")
		return out, err
	}
	f, err := fs.Open(p.FileInPath())
	if err != nil {
		log.Err(err).Msgf("Age, decryptFromInMemFS, fs.Open(p.FileInPath()) %s", p.FileInPath())
		return out, err
	}
	defer f.Close()
	r, err := age.Decrypt(f, identity)
	if err != nil {
		log.Err(err).Msg("Age, decryptFromInMemFS, age.Decrypt")
		return out, err
	}
	p.FnOut, _, _ = strings.Cut(p.FnIn, ".age")
	if _, cerr := io.Copy(out, r); cerr != nil {
		log.Err(cerr).Msg("Age, decryptFromInMemFS, io.Copy")
		return out, cerr
	}
	return out, err
}

func (a *Age) UnGzipAndDecrypt(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	c := compression.NewCompression()
	err := a.Decrypt(p)
	if err != nil {
		log.Err(err)
		return err
	}

	err = c.UnGzip(p)
	if err != nil {
		log.Err(err)
		return err
	}

	return err
}

func (a *Age) DecryptAndUnGzipToInMemFs(p *filepaths.Path, fs memfs.MemFS, unzipDir string) error {
	if p == nil {
		return errors.New("need to include a path")
	}

	b := p.ReadFileInPath()

	p.DirIn = "."
	err := fs.MakeFileIn(p, b)
	if err != nil {
		log.Err(err)
		return err
	}
	c := compression.NewCompression()
	err = a.DecryptToMemFsFile(p, fs)
	if err != nil {
		log.Err(err)
		return err
	}
	// fn in is now the unencrypted version, so fn.out -> fn.in
	p.FnIn = p.FnOut
	p.DirOut = unzipDir
	err = c.UnGzipFromInMemFsOutToInMemFS(p, fs)
	if err != nil {
		log.Err(err)
		return err
	}
	return err
}
