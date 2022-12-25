package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	choreography "github.com/zeus-fyi/zeus/choreography/server"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := choreography.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
