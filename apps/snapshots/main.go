package snapshots

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/snapshots/download"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := download.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
