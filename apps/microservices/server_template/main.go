package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := db_api.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
