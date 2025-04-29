package command

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Root struct {
	cmd *cobra.Command

	name      string
	version   string
	commit    string
	buildDate string
	builtBy   string

	loglevel  string
	jsonlog   bool
	debug     bool
	colormode string
}

//nolint:lll
func NewRoot(name, version, commit, buildDate, builtBy string) *Root {
	root := &Root{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:   name,
			Short: "Transform fixed-length records into JSON objects",
			Long:  "Transform fixed-length records into JSON objects",
			Example: "  " + name + " fold -c config.yaml < input.txt > output.json" + "\n" +
				"  " + name + " unfold -c config.yaml < input.json > output.txt",
			Args: cobra.NoArgs,
		},
		name:      name,
		version:   version,
		commit:    commit,
		buildDate: buildDate,
		builtBy:   builtBy,
		loglevel:  "warn",
		jsonlog:   false,
		debug:     false,
		colormode: "auto",
	}

	root.cmd.PersistentFlags().StringVarP(&root.loglevel, "verbosity", "v", root.loglevel, "set the log level (debug, info, warn, error)")
	root.cmd.PersistentFlags().BoolVar(&root.jsonlog, "log-json", root.jsonlog, "enable JSON log format")
	root.cmd.PersistentFlags().BoolVar(&root.debug, "debug", root.debug, "add debug information to logs (very slow)")
	root.cmd.PersistentFlags().StringVar(&root.colormode, "colormode", root.colormode, "set the color mode (auto, yes, no)")

	root.cmd.AddCommand(NewFoldCommand())
	root.cmd.AddCommand(NewUnfoldCommand())

	root.cmd.PersistentPreRun = root.execute

	return root
}

func (r *Root) Execute() error {
	if err := r.cmd.Execute(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r *Root) execute(_ *cobra.Command, _ []string) {
	initLog(r.colormode, r.jsonlog, r.debug, r.loglevel)

	log.Info().Msgf("Starting %v %v (commit=%v date=%v by=%v)", r.name, r.version, r.commit, r.buildDate, r.builtBy)
}

//nolint:cyclop
func initLog(colormode string, jsonlog bool, debug bool, verbosity string) {
	color := false

	switch strings.ToLower(colormode) {
	case "auto":
		if isatty.IsTerminal(os.Stdout.Fd()) && runtime.GOOS != "windows" {
			color = true
		}
	case "yes", "true", "1", "on", "enable":
		color = true
	}

	var logger zerolog.Logger
	if jsonlog {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger() // .With().Caller().Logger()
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: !color}) //nolint:exhaustruct
	}

	if debug {
		logger = logger.With().Caller().Logger()
	}

	log.Logger = logger

	switch verbosity {
	case "trace", "5":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Info().Msg("Logger level set to trace")
	case "debug", "4":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("Logger level set to debug")
	case "info", "3":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Logger level set to info")
	case "warn", "2":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error", "1":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
}
