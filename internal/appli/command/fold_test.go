package command_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/cgi-fr/posimap/internal/appli/command"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func BenchmarkFold(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// File is loaded once before the loop to avoid benchmarking I/O operations
	datafile, err := os.ReadFile("testdata/data.fixed-width")
	if err != nil {
		b.Fatalf("Failed to open test data file: %v", err)
	}

	for b.Loop() {
		command := command.NewFoldCommand("posimap", "testgroup")
		command.SetArgs([]string{"-s", "testdata/schema.yaml"})
		command.SetIn(bytes.NewReader(datafile))
		command.SetOut(io.Discard)

		if err := command.Execute(); err != nil {
			b.Fatalf("Failed to execute command: %v", err)
		}
	}
}

type FoldTest struct {
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

func LoadFoldTestFromFile(filename string) (FoldTest, error) {
	var test FoldTest

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return test, fmt.Errorf("%w", err)
	}

	if err := yaml.Unmarshal(bytes, &test); err != nil {
		return test, fmt.Errorf("%w", err)
	}

	return test, nil
}

func RunFoldTestFromFile(t *testing.T, filename string) {
	t.Helper()

	test, err := LoadFoldTestFromFile(filename)
	require.NoError(t, err)

	stdin, err := os.ReadFile(test.Stdin)
	require.NoError(t, err)

	logLvl, err := zerolog.ParseLevel(test.LogLevel)
	require.NoError(t, err)

	log.Logger = zerolog.New(os.Stderr)

	zerolog.SetGlobalLevel(logLvl)

	command := command.NewFoldCommand("posimap", "testgroup")

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

func TestFold(t *testing.T) {
	t.Parallel()

	t.Run("01.yaml", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "testdata/01.yaml") })
	t.Run("02.yaml", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "testdata/02.yaml") })
	t.Run("03.yaml", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "testdata/03.yaml") })
}
