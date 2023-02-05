package memfs

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/memoryfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"io/fs"
)

type MemFS struct {
	*memoryfs.FS
}

func NewMemFs() MemFS {
	memfs := memoryfs.New()
	m := MemFS{memfs}
	return m
}

func (m *MemFS) ReadFileOutPath(p *filepaths.Path) ([]byte, error) {
	var b []byte
	if p == nil {
		return b, errors.New("need to include a path")
	}
	b, err := fs.ReadFile(m, p.FileOutPath())
	if err != nil {
		log.Err(err).Msgf("ReadFileOutPath %s", p.FileOutPath())
		return b, err
	}
	return b, nil
}

func (m *MemFS) ReadFileInPath(p *filepaths.Path) ([]byte, error) {
	var b []byte
	if p == nil {
		return b, errors.New("need to include a path")
	}
	b, err := fs.ReadFile(m, p.FileInPath())
	if err != nil {
		log.Err(err).Msgf("ReadFileInPath %s", p.FileInPath())
		return b, err
	}
	return b, nil
}

func (m *MemFS) MakeFileDirOutFnInPath(p *filepaths.Path, content []byte) error {
	merr := m.MkPathDirAll(p)
	if merr != nil {
		return merr
	}
	if err := m.WriteFile(p.FileDirOutFnInPath(), content, 0644); err != nil {
		return err
	}
	return nil
}

func (m *MemFS) MakeFileIn(p *filepaths.Path, content []byte) error {
	merr := m.MkPathDirAll(p)
	if merr != nil {
		log.Err(merr).Msgf("MemFS, MakeFile fileIn path %s, fileOut path %s", p.FileInPath(), p.FileOutPath())
		return merr
	}
	if err := m.WriteFile(p.FileInPath(), content, 0644); err != nil {
		log.Err(err).Msgf("MemFS, WriteFile, fileOut path %s", p.FileInPath())
		return err
	}
	return nil
}

func (m *MemFS) MakeFileOut(p *filepaths.Path, content []byte) error {
	merr := m.MkPathDirAll(p)
	if merr != nil {
		log.Err(merr).Msgf("MemFS, MakeFile fileIn path %s, fileOut path %s", p.FileInPath(), p.FileOutPath())
		return merr
	}
	if err := m.WriteFile(p.FileOutPath(), content, 0644); err != nil {
		log.Err(err).Msgf("MemFS, WriteFile, fileOut path %s", p.FileOutPath())
		return err
	}
	return nil
}

func (m *MemFS) MkPathDirAll(p *filepaths.Path) error {
	if err := m.MkdirAll(p.DirOut, 0700); err != nil {
		return err
	}
	if err := m.MkdirAll(p.DirIn, 0700); err != nil {
		return err
	}
	return nil
}
