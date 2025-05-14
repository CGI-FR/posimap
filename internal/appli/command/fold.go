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

	fold.cmd.Run = fold.execute

	return fold.cmd
}

//nolint:cyclop
func (f *Fold) execute(cmd *cobra.Command, _ []string) {
	reader := buffer.NewBufferReader(cmd.InOrStdin())
	writer := jsonline.NewWriter(cmd.OutOrStdout())

	cfg, err := config.LoadConfigFromFile(f.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration file")
	}

	schema, err := cfg.Compile(config.Trim(f.trim), config.Charset(f.charset))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to compile configuration file")
	}

	record, err := schema.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to build record unmarshaler")
	}

	if err := f.execUntilEOF(cfg, reader, writer, record); err != nil {
		log.Fatal().Err(err).Msg("Fold command ended with error")
	}

	log.Info().Msg("Fold command completed successfully")
}

func (f *Fold) execUntilEOF(cfg config.Config, rdr *buffer.Buffer, wrtr *jsonline.Writer, rcrd api.Record) error {
	defer func() {
		if err := wrtr.WriteEOF(); err != nil {
			log.Fatal().Err(err).Msg("Failed to close stream")
		}
	}()

	for {
		if err := rdr.Reset(cfg.Length); errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := rcrd.Unmarshal(rdr); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := rcrd.Export(wrtr); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
}
