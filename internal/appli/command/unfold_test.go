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

func BenchmarkUnfold(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// File is loaded once before the loop to avoid benchmarking I/O operations
	datafile, err := os.ReadFile("testdata/unfold/07-complete/stdin.fixed-width")
	if err != nil {
		b.Fatalf("Failed to open test data file: %v", err)
	}

	for b.Loop() {
		command := command.NewUnfoldCommand("posimap", "testgroup")
		command.SetArgs([]string{"-s", "testdata/unfold/07-complete/schema.yaml"})
		command.SetIn(bytes.NewReader(datafile))
		command.SetOut(io.Discard)

		if err := command.Execute(); err != nil {
			b.Fatalf("Failed to execute command: %v", err)
		}
	}
}

func LoadUnfoldTestFromFile(filename string) (FoldTest, error) {
	var test FoldTest

	bytes, err := os.ReadFile("testdata/unfold/" + filename)
	if err != nil {
		return test, fmt.Errorf("%w", err)
	}

	if err := yaml.Unmarshal(bytes, &test); err != nil {
		return test, fmt.Errorf("%w", err)
	}

	return test, nil
}

func RunUnfoldTestFromFile(t *testing.T, filename string) {
	t.Helper()

	test, err := LoadUnfoldTestFromFile(filename)
	require.NoError(t, err)

	stdin, err := os.ReadFile(test.Stdin)
	require.NoError(t, err)

	logLvl, err := zerolog.ParseLevel(test.LogLevel)
	require.NoError(t, err)

	log.Logger = zerolog.New(os.Stderr)

	zerolog.SetGlobalLevel(logLvl)

	command := command.NewUnfoldCommand("posimap", "testgroup")

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

func TestUnfold(t *testing.T) {
	t.Parallel()

	t.Run("01-simple", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "01-simple.yaml") })
	t.Run("02-simple-separator", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "02-simple-separator.yaml") })
	t.Run("03-multiple", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "03-multiple.yaml") })
	t.Run("04-nested", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "04-nested.yaml") })
	t.Run("05-occurs", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "05-occurs.yaml") })
	t.Run("06-redefines", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "06-redefines.yaml") })
	t.Run("07-complete", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "07-complete.yaml") })
}
