package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	name      string
	version   string
	commit    string
	buildDate string
	builtBy   string
)

func main() {
	// nolint: exhaustivestruct
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cmd, err := NewRootCommand()
	if err != nil {
		log.Error().Err(err).Msg("End of process")
		os.Exit(1)
	}

	err = cmd.Execute()
	if err != nil {
		log.Error().Err(err).Msg("End of process")
		os.Exit(1)
	}
}
