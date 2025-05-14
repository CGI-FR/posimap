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
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Stdin       string            `yaml:"stdin"`
	Flags       map[string]string `yaml:"flags"`
	LogLevel    string            `yaml:"loglevel"`
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

	logLvl, err := zerolog.ParseLevel(test.LogLevel)
	require.NoError(t, err)

	log.Logger = zerolog.New(os.Stderr)

	zerolog.SetGlobalLevel(logLvl)

	for flag, value := range test.Flags {
		command.SetArgs([]string{flag, value})
	}

	actualStdout := &bytes.Buffer{}
	actualStderr := &bytes.Buffer{}

	command.SetIn(bytes.NewReader(stdin))
	command.SetOut(actualStdout)
	command.SetErr(actualStderr)

	if err := command.Execute(); assert.NoError(t, err) {
		expectedStdout, err := os.ReadFile(test.Expected.Stdout)
		require.NoError(t, err)

		assert.Equal(t, string(expectedStdout), actualStdout.String(), "stdout mismatch")

		expectedStderr, err := os.ReadFile(test.Expected.Stderr)
		require.NoError(t, err)

		assert.Equal(t, string(expectedStderr), actualStderr.String(), "stderr mismatch")
	}
}
