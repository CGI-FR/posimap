package command

import (
	"os"

	"github.com/cgi-fr/posimap/internal/infra/config"
	"github.com/cgi-fr/posimap/internal/infra/object"
	"github.com/cgi-fr/posimap/internal/infra/record"
	"github.com/cgi-fr/posimap/pkg/data"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/unicode"
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

func (u *Unfold) execute(_ *cobra.Command, _ []string) {
	source := object.NewJSONLineSource(os.Stdin, unicode.UTF8)
	sink := record.NewRecordSink(os.Stdout)

	config, err := config.LoadConfigFromFile(u.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load schema")
	}

	root := data.NewBuilder().Build(config.Schema.Compile())

	if err := data.TransformObjectsToRecords(root, source, sink); err != nil {
		log.Fatal().Err(err).Msg("Failed to process records")
	}

	log.Info().Msg("Unfold command completed successfully")
}
