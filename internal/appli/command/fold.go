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

package command

import (
	"errors"
	"io"
	"os"

	"github.com/cgi-fr/posimap/internal/infra/config"
	"github.com/cgi-fr/posimap/internal/infra/jsonline"
	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Fold struct {
	cmd *cobra.Command

	configfile string
	trim       bool
}

func NewFoldCommand(rootname string, groupid string) *cobra.Command {
	fold := &Fold{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "fold",
			Short:   "Transform fixed-length records into JSON objects",
			Long:    "Transform fixed-length records into JSON objects",
			Example: "  " + rootname + " fold -c schema.yaml < input.fixed-width > output.json",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile: "schema.yaml",
		trim:       false,
	}

	fold.cmd.Flags().StringVarP(&fold.configfile, "config", "c", fold.configfile, "set the config file")
	fold.cmd.Flags().BoolVarP(&fold.trim, "trim", "t", fold.trim, "trim the input records")

	fold.cmd.Run = fold.execute

	return fold.cmd
}

func (f *Fold) execute(_ *cobra.Command, _ []string) {
	buffer := buffer.NewBufferReader(os.Stdin)
	writer := jsonline.NewWriter(os.Stdout)

	config, err := config.LoadConfigFromFile(f.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load schema")
	}

	if f.trim {
		// TODO: Implement trim functionality
		// config.Trim = true
	}

	schema := config.Compile()

	record, err := schema.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to build record")
	}

	for {
		if err := record.Unmarshal(buffer); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal().Err(err).Msg("Failed to unmarshal record")
		}

		if err := record.Export(writer); err != nil {
			log.Fatal().Err(err).Msg("Failed to export record")
		}

		if err := buffer.Reset(); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal().Err(err).Msg("Failed to reset buffer")
		}
	}

	if err := writer.WriteEOF(); err != nil {
		log.Fatal().Err(err).Msg("Failed to close stream")
	}

	log.Info().Msg("Fold command completed successfully")
}
