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

//nolint:paralleltest // Cobra commands are not parallel-safe
func TestUnfold(t *testing.T) {
	t.Run("01-simple", func(t *testing.T) { RunUnfoldTestFromFile(t, "01-simple.yaml") })
	t.Run("02-simple-separator", func(t *testing.T) { RunUnfoldTestFromFile(t, "02-simple-separator.yaml") })
	t.Run("03-multiple", func(t *testing.T) { RunUnfoldTestFromFile(t, "03-multiple.yaml") })
	t.Run("04-nested", func(t *testing.T) { RunUnfoldTestFromFile(t, "04-nested.yaml") })
	t.Run("05-occurs", func(t *testing.T) { RunUnfoldTestFromFile(t, "05-occurs.yaml") })
	t.Run("06-redefines", func(t *testing.T) { RunUnfoldTestFromFile(t, "06-redefines.yaml") })
	t.Run("07-complete", func(t *testing.T) { RunUnfoldTestFromFile(t, "07-complete.yaml") })
	t.Run("08-missing-filler", func(t *testing.T) { RunUnfoldTestFromFile(t, "08-missing-filler.yaml") })
	t.Run("09-trim", func(t *testing.T) { RunUnfoldTestFromFile(t, "09-trim.yaml") })
	t.Run("10-charsets", func(t *testing.T) { RunUnfoldTestFromFile(t, "10-charsets.yaml") })
	t.Run("11-json", func(t *testing.T) { RunUnfoldTestFromFile(t, "11-json.yaml") })
	t.Run("12-flags", func(t *testing.T) { RunUnfoldTestFromFile(t, "12-flags.yaml") })
	t.Run("99-import-order", func(t *testing.T) { RunUnfoldTestFromFile(t, "99-import-order.yaml") })
}
