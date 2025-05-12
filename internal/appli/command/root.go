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
			Example: "  " + name + " fold   -c schema.yaml < input.txt  > output.json" + "\n" +
				"  " + name + " unfold -c schema.yaml < input.json > output.txt",
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

	root.cmd.AddGroup(&cobra.Group{ID: "transform", Title: "Transform commands:"})
	root.cmd.AddCommand(NewFoldCommand(root.name, "transform"))
	root.cmd.AddCommand(NewUnfoldCommand(root.name, "transform"))

	root.cmd.AddGroup(&cobra.Group{ID: "helpers", Title: "Helper commands:"})
	root.cmd.AddCommand(NewGraphCommand(root.name, "helpers"))

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
