package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	load_simulator "github.com/zeus-fyi/zeus/load-simulator/server"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := load_simulator.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
