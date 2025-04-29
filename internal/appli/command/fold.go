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

type Fold struct {
	cmd *cobra.Command

	configfile string
}

func NewFoldCommand() *cobra.Command {
	fold := &Fold{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:   "fold",
			Short: "Transform fixed-length records into JSON objects",
			Long:  "Transform fixed-length records into JSON objects",
			Example: "  fold -c config.yaml < input.txt > output.json" + "\n" +
				"  unfold -c config.yaml < input.json > output.txt",
			Args: cobra.NoArgs,
		},
		configfile: "config.yaml",
	}

	fold.cmd.Flags().StringVarP(&fold.configfile, "config", "c", fold.configfile, "set the config file")

	fold.cmd.Run = fold.execute

	return fold.cmd
}

func (f *Fold) execute(_ *cobra.Command, _ []string) {
	source := record.NewRecordSource(os.Stdin, unicode.UTF8)
	sink := object.NewJSON(os.Stdout)

	config, err := config.LoadConfigFromFile(f.configfile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load schema")
	}

	root := data.NewBuilder().Build(config.Schema.Compile())

	if err := data.TransformRecordsToObjects(root, source, sink); err != nil {
		log.Fatal().Err(err).Msg("Failed to process records")
	}

	log.Info().Msg("Fold command completed successfully")
}
