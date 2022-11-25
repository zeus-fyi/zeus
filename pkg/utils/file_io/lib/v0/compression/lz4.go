package compression

import (
	"errors"
	"os"

	"github.com/pierrec/lz4"
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (c *Compression) Lz4Decompress(p *filepaths.Path) error {
	if p == nil {
		return errors.New("need to include a path")
	}
	r, err := os.Open(p.FileInPath())
	if err != nil {
		log.Err(err).Msg("Compression: Lz4Decompress, os.Open(p.FileInPath())")
		return err
	}
	defer r.Close()
	lz4Reader := lz4.NewReader(r)
	if err != nil {
		log.Err(err).Msg("Compression: Lz4Decompress, lz4.NewReader(r)")
		return err
	}
	return tarReader(p, lz4Reader)
}
