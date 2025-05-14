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
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/internal/appli/charsets"
	"github.com/cgi-fr/posimap/internal/appli/config"
	"github.com/cgi-fr/posimap/internal/infra/jsonline"
	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Unfold struct {
	cmd *cobra.Command

	configfile string
	charset    string
}

func NewUnfoldCommand(rootname string, groupid string) *cobra.Command {
	unfold := &Unfold{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "unfold",
			Short:   "Transform JSON objects into fixed-length records",
			Long:    "Transform JSON objects into fixed-length records",
			Example: "  " + rootname + "unfold -s schema.yaml < input.jsonl > output.fixed-width",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile: "schema.yaml",
		charset:    charsets.ISO88591,
	}

	unfold.cmd.Flags().StringVarP(&unfold.configfile, "schema", "s", unfold.configfile, "set the schema file")
	unfold.cmd.Flags().StringVarP(&unfold.charset, "charset", "c", unfold.charset, "set the charset for output records") //nolint:lll

	unfold.cmd.Run = unfold.execute

	return unfold.cmd
}

func (u *Unfold) execute(cmd *cobra.Command, _ []string) {
	reader := jsonline.NewReader(cmd.InOrStdin())
	writer := buffer.NewBufferWriter(cmd.OutOrStdout())

	cfg, err := config.LoadConfigFromFile(u.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration file")
	}

	schema, err := cfg.Compile(config.Trim(true), config.Charset(u.charset))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to compile configuration file")
	}

	record, err := schema.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to build record marshaler")
	}

	if err := u.execUntilEOF(cfg, writer, reader, record); err != nil {
		log.Fatal().Err(err).Msg("Unfold command failed")
	}

	log.Info().Msg("Unfold command completed successfully")
}

func (u *Unfold) execUntilEOF(cfg config.Config, buffer api.Buffer, reader document.Reader, record api.Record) error {
	space, err := charsets.GetByteInCharset(u.charset, ' ')
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	sep, err := charsets.GetBytesInCharset(u.charset, cfg.Separator)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for {
		if err := buffer.Reset(cfg.Length, sep...); err != nil {
			return fmt.Errorf("%w", err)
		}

		buffer.Fill(space)
		record.Reset()

		document, err := reader.Read()
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := record.Import(document); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
}
