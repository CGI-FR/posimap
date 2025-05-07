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

	"github.com/cgi-fr/posimap/internal/infra/config"
	"github.com/cgi-fr/posimap/internal/infra/jsonline"
	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Unfold struct {
	cmd *cobra.Command

	configfile string
}

func NewUnfoldCommand(rootname string, groupid string) *cobra.Command {
	unfold := &Unfold{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "unfold",
			Short:   "Transform JSON objects into fixed-length records",
			Long:    "Transform JSON objects into fixed-length records",
			Example: "  " + rootname + "unfold -c schema.yaml < input.json > output.fixed-width",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile: "schema.yaml",
	}

	unfold.cmd.Flags().StringVarP(&unfold.configfile, "config", "c", unfold.configfile, "set the config file")

	unfold.cmd.Run = unfold.execute

	return unfold.cmd
}

func (u *Unfold) execute(cmd *cobra.Command, _ []string) {
	reader := jsonline.NewReader(cmd.InOrStdin())
	writer := buffer.NewBufferWriter(cmd.OutOrStdout())

	cfg, err := config.LoadConfigFromFile(u.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load schema")
	}

	schema := cfg.Compile()

	record, err := schema.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to build record")
	}

	for {
		if err := record.Import(reader); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal().Err(err).Msg("Failed to import document")
		}

		if err := record.Marshal(writer); err != nil {
			log.Fatal().Err(err).Msg("Failed to marshal record")
		}

		if err := writer.Reset(); err != nil {
			log.Fatal().Err(err).Msg("Failed to reset buffer")
		}
	}

	log.Info().Msg("Unfold command completed successfully")
}
