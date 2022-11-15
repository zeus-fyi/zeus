package main

import (
	"github.com/rs/zerolog/log"
	hercules_server "github.com/zeus-fyi/hercules/server"
)

func main() {

	if err := hercules_server.Cmd.Execute(); err != nil {
		log.Err(err)
	}
}
