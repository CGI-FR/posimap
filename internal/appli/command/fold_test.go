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

func BenchmarkFold(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// File is loaded once before the loop to avoid benchmarking I/O operations
	datafile, err := os.ReadFile("testdata/fold/07-complete/stdin.fixed-width")
	if err != nil {
		b.Fatalf("Failed to open test data file: %v", err)
	}

	for b.Loop() {
		command := command.NewFoldCommand("posimap", "testgroup")
		command.SetArgs([]string{"-s", "testdata/fold/07-complete/schema.yaml"})
		command.SetIn(bytes.NewReader(datafile))
		command.SetOut(io.Discard)

		if err := command.Execute(); err != nil {
			b.Fatalf("Failed to execute command: %v", err)
		}
	}
}

func RunFoldTestFromFile(t *testing.T, filename string) {
	t.Helper()

	command := command.NewFoldCommand("posimap", "testgroup")

	RunCommandTestFromFile(t, command, "fold/"+filename)
}

func TestFold(t *testing.T) {
	t.Parallel()

	t.Run("01-simple", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "01-simple.yaml") })
	t.Run("02-simple-separator", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "02-simple-separator.yaml") })
	t.Run("03-multiple", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "03-multiple.yaml") })
	t.Run("04-nested", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "04-nested.yaml") })
	t.Run("05-occurs", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "05-occurs.yaml") })
	t.Run("06-redefines", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "06-redefines.yaml") })
	t.Run("07-complete", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "07-complete.yaml") })
	t.Run("08-missing-filler", func(t *testing.T) { t.Parallel(); RunFoldTestFromFile(t, "08-missing-filler.yaml") })
}
