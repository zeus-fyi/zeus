package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	hades_server "github.com/zeus-fyi/hades/server"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if err := hades_server.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
