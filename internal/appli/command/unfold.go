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

	unfold.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := unfold.execute(cmd, args); err != nil {
			log.Error().Err(err).Msg("Unfold command failed")

			return err
		}

		return nil
	}

	return unfold.cmd
}

func (u *Unfold) execute(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	reader := jsonline.NewReader(cmd.InOrStdin())
	writer := buffer.NewBufferWriter(cmd.OutOrStdout())

	cfg, err := config.LoadConfigFromFile(u.configfile)
	if err != nil {
		return fmt.Errorf("failed to load configuration file : %w", err)
	}

	schema, err := cfg.Compile(config.Trim(true), config.Charset(u.charset))
	if err != nil {
		return fmt.Errorf("failed to compile configuration file : %w", err)
	}

	record, err := schema.Build()
	if err != nil {
		return fmt.Errorf("failed to build record marshaler : %w", err)
	}

	if err := u.execUntilEOF(cfg, writer, reader, record); err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Info().Msg("Unfold command completed successfully")

	return nil
}

func (u *Unfold) execUntilEOF(cfg config.Config, buffer api.Buffer, reader document.Reader, record api.Record) error {
	space, err := charsets.GetByteInCharset(u.charset, ' ')
	if err != nil {
		return fmt.Errorf("failed to get space in charset : %w", err)
	}

	sep, err := charsets.GetBytesInCharset(u.charset, cfg.Separator)
	if err != nil {
		return fmt.Errorf("failed to get separator in charset : %w", err)
	}

	for {
		if err := buffer.Reset(space, cfg.Length, sep...); err != nil {
			return fmt.Errorf("failed to dump buffer : %w", err)
		}

		record.Reset()

		document, err := reader.Read()
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to read next document : %w", err)
		}

		if err := record.Import(document); err != nil {
			return fmt.Errorf("failed to import record : %w", err)
		}

		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("failed to marshal buffer : %w", err)
		}
	}
}
