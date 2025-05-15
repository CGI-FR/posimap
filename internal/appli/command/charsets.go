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
	"fmt"
	"strings"

	"github.com/cgi-fr/posimap/internal/appli/charsets"
	"github.com/spf13/cobra"
)

type Charsets struct {
	cmd *cobra.Command
}

func NewCharsetsCommand(rootname string, groupid string) *cobra.Command {
	charsets := &Charsets{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "charsets",
			Short:   "List all supported charsets",
			Long:    "List all supported charsets",
			Example: "  " + rootname + " charsets",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
	}

	charsets.cmd.Run = charsets.execute

	return charsets.cmd
}

func (c *Charsets) execute(cmd *cobra.Command, _ []string) {
	const (
		nameWidth   = 20
		descWidth   = 40
		nameHeader  = "CHARSET NAME"
		descHeader  = "DESCRIPTION"
		headerUnder = "-"
	)

	// Format pour les données et les en-têtes
	format := fmt.Sprintf("%%-%ds%%-%ds\n", nameWidth, descWidth)

	// Affichage des en-têtes
	cmd.Printf(format, nameHeader, descHeader)
	cmd.Printf(format,
		strings.Repeat(headerUnder, nameWidth-1),
		strings.Repeat(headerUnder, descWidth-1))

	// Affichage des données
	for _, charset := range charsets.List() {
		c, err := charsets.Get(charset)
		if err == nil {
			cmd.Printf(format, charset, c.String())
		}
	}
}
