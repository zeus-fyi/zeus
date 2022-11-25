package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	snapshot_init "github.com/zeus-fyi/snapshots/downloader"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if err := snapshot_init.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
