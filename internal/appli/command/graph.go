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

	"github.com/cgi-fr/posimap/internal/appli/charsets"
	"github.com/cgi-fr/posimap/internal/appli/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Graph struct {
	cmd *cobra.Command

	configfile       string
	showDependencies bool
}

func NewGraphCommand(rootname string, groupid string) *cobra.Command {
	graph := &Graph{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "graph",
			Short:   "Export the schema as a graphviz graph",
			Long:    "Export the schema as a graphviz graph",
			Example: "  " + rootname + "graph -s schema.yaml > schema.dot",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile:       "schema.yaml",
		showDependencies: false,
	}

	graph.cmd.Flags().StringVarP(&graph.configfile, "schema", "s", graph.configfile, "set the schema file")
	graph.cmd.Flags().BoolVarP(&graph.showDependencies, "dependencies", "d", graph.showDependencies, "show dependencies")

	graph.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := graph.execute(cmd, args); err != nil {
			log.Error().Err(err).Msg("Graph command failed")

			return err
		}

		return nil
	}

	return graph.cmd
}

func (g *Graph) execute(_ *cobra.Command, _ []string) error {
	cfg, err := config.LoadConfigFromFile(g.configfile)
	if err != nil {
		return fmt.Errorf("failed to load configuration file : %w", err)
	}

	schema, err := cfg.Compile(config.Trim(true), config.Charset(charsets.ISO88591))
	if err != nil {
		return fmt.Errorf("failed to compile configuration file : %w", err)
	}

	if err := schema.PrintGraph(g.showDependencies); err != nil {
		return fmt.Errorf("failed to load configuration file : %w", err)
	}

	return nil
}
