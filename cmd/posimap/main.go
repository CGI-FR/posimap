// Copyright (C) 2025 CGI France
//
// This file is part of posimap.
//
// posimap is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// posimap is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with posimap.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	"github.com/cgi-fr/posimap/internal/infra/config"
	"github.com/cgi-fr/posimap/internal/infra/object"
	"github.com/cgi-fr/posimap/internal/infra/record"
	"github.com/cgi-fr/posimap/pkg/data"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/encoding/unicode"
)

// Provisioned by ldflags.
var (
	name      string //nolint: gochecknoglobals
	version   string //nolint: gochecknoglobals
	commit    string //nolint: gochecknoglobals
	buildDate string //nolint: gochecknoglobals
	builtBy   string //nolint: gochecknoglobals
)

func main() {
	//nolint: exhaustruct
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msgf("%v %v (commit=%v date=%v by=%v)", name, version, commit, buildDate, builtBy)

	source := record.NewRecordSource(os.Stdin, unicode.UTF8)
	sink := object.NewJSON(os.Stdout)

	config, err := config.LoadConfigFromFile("schema.yaml")
	if err != nil {
		log.Error().Err(err).Msg("failed to load schema")

		return
	}

	root := data.NewBuilder().Build(config.Schema.Compile())

	if err := data.TransformRecordsToObjects(root, source, sink); err != nil {
		log.Error().Err(err).Msg("failed to process records")
	}

	fmt.Println()
}
