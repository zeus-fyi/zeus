package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	hera_server "github.com/zeus-fyi/chatgpt/server"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := hera_server.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
