package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	echo_server_template "github.com/zeus-fyi/zeus/echo-server-template/server"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := echo_server_template.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
