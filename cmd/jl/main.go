// Copyright (C) 2021 CGI France
//
// This file is part of JL.
//
// JL is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// JL is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with JL.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Provisioned by ldflags
//
//nolint:gochecknoglobals
var (
	name      string
	version   string
	commit    string
	buildDate string
	builtBy   string
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) //nolint:exhaustruct,exhaustivestruct

	cmd, err := NewRootCommand()
	if err != nil {
		log.Error().Err(err).Msg("cannot initialize command")
		os.Exit(1)
	}

	err = cmd.Execute()
	if err != nil {
		log.Error().Err(err).Msg("end of process")
		os.Exit(1)
	}
}
