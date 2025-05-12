package command

import (
	"github.com/cgi-fr/posimap/internal/appli/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/charmap"
)

type Graph struct {
	cmd *cobra.Command

	configfile        string
	charset           string
	showDenpendencies bool
}

func NewGraphCommand(rootname string, groupid string) *cobra.Command {
	graph := &Graph{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "graph",
			Short:   "Export the schema as a graphviz graph",
			Long:    "Export the schema as a graphviz graph",
			Example: "  " + rootname + "graph -c schema.yaml > schema.dot",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
		configfile:        "schema.yaml",
		charset:           charmap.ISO8859_1.String(),
		showDenpendencies: false,
	}

	graph.cmd.Flags().StringVarP(&graph.configfile, "config", "c", graph.configfile, "set the config file")
	graph.cmd.Flags().StringVarP(&graph.charset, "charset", "C", graph.charset, "set the charset for the output records") //nolint:lll
	graph.cmd.Flags().BoolVarP(&graph.showDenpendencies, "dependencies", "d", graph.showDenpendencies, "show dependencies")

	graph.cmd.Run = graph.execute

	return graph.cmd
}

func (g *Graph) execute(_ *cobra.Command, _ []string) {
	cfg, err := config.LoadConfigFromFile(g.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration file")
	}

	schema, err := cfg.Compile(config.Trim(true), config.Charset(g.charset))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to compile configuration file")
	}

	schema.PrintGraph(g.showDenpendencies)
}
