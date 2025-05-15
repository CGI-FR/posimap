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

package command_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type CommandTest struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Stdin       string         `yaml:"stdin"`
	Flags       map[string]any `yaml:"flags"`
	LogLevel    string         `yaml:"loglevel"`
	Expected    struct {
		Stdout string `yaml:"stdout"`
		Stderr string `yaml:"stderr"`
		Exit   int    `yaml:"exit"`
	} `yaml:"expected"`
}

func LoadCommandTestFromFile(filename string) (CommandTest, error) {
	var test CommandTest

	bytes, err := os.ReadFile("testdata/" + filename)
	if err != nil {
		return test, fmt.Errorf("%w", err)
	}

	if err := yaml.Unmarshal(bytes, &test); err != nil {
		return test, fmt.Errorf("%w", err)
	}

	return test, nil
}

func RunCommandTestFromFile(t *testing.T, command *cobra.Command, filename string) {
	t.Helper()

	test, err := LoadCommandTestFromFile(filename)
	require.NoError(t, err)

	stdin, err := os.ReadFile(test.Stdin)
	require.NoError(t, err)

	args := []string{}

	for flag, value := range test.Flags {
		switch param := value.(type) {
		case string:
			args = append(args, flag, param)
		case bool:
			if param {
				args = append(args, flag)
			}
		}
	}

	command.SetArgs(args)

	actualStdout := &bytes.Buffer{}
	actualStderr := &bytes.Buffer{}

	command.SetIn(bytes.NewReader(stdin))
	command.SetOut(actualStdout)
	command.SetErr(actualStderr)

	logLvl, err := zerolog.ParseLevel(test.LogLevel)
	require.NoError(t, err)

	log.Logger = zerolog.New(actualStderr)

	zerolog.SetGlobalLevel(logLvl)

	_ = command.Execute()

	expectedStdout, err := os.ReadFile(test.Expected.Stdout)
	require.NoError(t, err)

	assert.Equal(t, string(expectedStdout), actualStdout.String(), "stdout mismatch")

	expectedStderr, err := os.ReadFile(test.Expected.Stderr)
	require.NoError(t, err)

	assert.Equal(t, string(expectedStderr), actualStderr.String(), "stderr mismatch")
}
