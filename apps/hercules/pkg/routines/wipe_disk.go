package routines

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
)

func WipeDisk(ctx context.Context, path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("WipeDisk: RemoveAll %s", path)
		return err
	}
	err = os.MkdirAll(path, 644)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("WipeDisk: MkdirAll %s", path)
		return err
	}
	return err
}
