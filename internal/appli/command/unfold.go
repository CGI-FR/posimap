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

	"github.com/cgi-fr/posimap/internal/appli/config"
	"github.com/cgi-fr/posimap/internal/infra/jsonline"
	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/cgi-fr/posimap/pkg/posimap/core/record"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/charmap"
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
			Example: "  " + rootname + "unfold -c schema.yaml < input.json > output.fixed-width",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile: "schema.yaml",
		charset:    charmap.ISO8859_1.String(),
	}

	unfold.cmd.Flags().StringVarP(&unfold.configfile, "config", "c", unfold.configfile, "set the config file")
	unfold.cmd.Flags().StringVarP(&unfold.charset, "charset", "C", unfold.charset, "set the charset for the output records") //nolint:lll

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

	record := u.buildRecord(cfg)

	space, err := getSpaceInCharset(u.charset)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get space character in charset")
	}

	for {
		if err := writer.Reset(cfg.Length); err != nil {
			log.Fatal().Err(err).Msg("Failed to prepare next byte buffer")
		}

		writer.Fill(space)
		record.Reset()

		document, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal().Err(err).Msg("Failed to read document")
		}

		if err := record.Import(document); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal().Err(err).Msg("Failed to convert document to record")
		}

		if err := record.Marshal(writer); err != nil {
			log.Fatal().Err(err).Msg("Failed to marshal record")
		}
	}

	log.Info().Msg("Unfold command completed successfully")
}

func (u *Unfold) buildRecord(cfg config.Config) *record.Object {
	schema, err := cfg.Compile(config.Trim(true), config.Charset(u.charset))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to compile configuration file")
	}

	record, err := schema.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to build record marshaler")
	}

	return record
}

var ErrUnsupportedCharset = errors.New("unsupported charset")

func getCharmap(charset string) (*charmap.Charmap, error) {
	for _, encoding := range charmap.All {
		if charmap, ok := encoding.(*charmap.Charmap); ok && charmap.String() == charset {
			return charmap, nil
		}
	}

	return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
}

func getSpaceInCharset(charset string) (byte, error) {
	charmap, err := getCharmap(charset)
	if err != nil {
		return 0, err
	}

	space, _ := charmap.EncodeRune(' ')

	return space, nil
}
