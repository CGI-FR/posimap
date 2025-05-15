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

type Fold struct {
	cmd *cobra.Command

	configfile string
	trim       bool
	charset    string
}

func NewFoldCommand(rootname string, groupid string) *cobra.Command {
	fold := &Fold{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "fold",
			Short:   "Transform fixed-length records into JSON objects",
			Long:    "Transform fixed-length records into JSON objects",
			Example: "  " + rootname + " fold -s schema.yaml < input.fixed-width > output.jsonl",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile: "schema.yaml",
		trim:       true,
		charset:    charsets.ISO88591,
	}

	fold.cmd.Flags().StringVarP(&fold.configfile, "schema", "s", fold.configfile, "set the schema file")
	fold.cmd.Flags().BoolVarP(&fold.trim, "trim", "t", fold.trim, "trim the input records")
	fold.cmd.Flags().StringVarP(&fold.charset, "charset", "c", fold.charset, "set the charset for input records")

	fold.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := fold.execute(cmd, args); err != nil {
			log.Error().Err(err).Msg("Fold command failed")

			return err
		}

		return nil
	}

	return fold.cmd
}

func (f *Fold) execute(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	reader := buffer.NewBufferReader(cmd.InOrStdin())
	writer := jsonline.NewWriter(cmd.OutOrStdout())

	cfg, err := config.LoadConfigFromFile(f.configfile)
	if err != nil {
		return fmt.Errorf("failed to load configuration file : %w", err)
	}

	schema, err := cfg.Compile(config.Trim(f.trim), config.Charset(f.charset))
	if err != nil {
		return fmt.Errorf("failed to compile configuration file : %w", err)
	}

	record, err := schema.Build()
	if err != nil {
		return fmt.Errorf("failed to build record marshaler : %w", err)
	}

	if err := f.execUntilEOF(cfg, reader, writer, record); err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Info().Msg("Fold command completed successfully")

	return nil
}

func (f *Fold) execUntilEOF(cfg config.Config, buffer api.Buffer, writer document.Writer, record api.Record) error {
	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal().Err(err).Msg("Failed to close stream")
		}
	}()

	space, err := charsets.GetByteInCharset(f.charset, ' ')
	if err != nil {
		return fmt.Errorf("failed to get space in charset : %w", err)
	}

	sep, err := charsets.GetBytesInCharset(f.charset, cfg.Separator)
	if err != nil {
		return fmt.Errorf("failed to get separator in charset : %w", err)
	}

	for {
		if err := buffer.Reset(space, cfg.Length, sep...); errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to read next buffer : %w", err)
		}

		if err := record.Unmarshal(buffer); err != nil {
			return fmt.Errorf("failed to marshal buffer : %w", err)
		}

		if err := record.Export(writer); err != nil {
			return fmt.Errorf("failed to export record : %w", err)
		}
	}
}
