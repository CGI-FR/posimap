//nolint:dupl
package command_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/cgi-fr/posimap/internal/appli/command"
	"github.com/rs/zerolog"
)

func BenchmarkUnfold(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// File is loaded once before the loop to avoid benchmarking I/O operations
	datafile, err := os.ReadFile("testdata/unfold/07-complete/stdin.jsonl")
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

func RunUnfoldTestFromFile(t *testing.T, filename string) {
	t.Helper()

	command := command.NewUnfoldCommand("posimap", "testgroup")

	RunCommandTestFromFile(t, command, "unfold/"+filename)
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
	t.Run("08-missing-filler", func(t *testing.T) { t.Parallel(); RunUnfoldTestFromFile(t, "08-missing-filler.yaml") })
}
