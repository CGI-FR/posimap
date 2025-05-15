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

//nolint:funlen
package config_test

import (
	"testing"

	"github.com/cgi-fr/posimap/internal/appli/config"
	"gotest.tools/assert"
)

func TestLoadSchemaFromYAML(t *testing.T) {
	t.Parallel()

	//nolint:exhaustruct
	testCases := []struct {
		name     string
		filename string
		expected config.Config
	}{
		{
			name:     "empty schema",
			filename: "schema-empty.yaml",
			expected: config.Config{
				Schema: config.Schema{},
			},
		},
		{
			name:     "simple schema",
			filename: "schema-simple.yaml",
			expected: config.Config{
				Schema: config.Schema{
					{
						Name:   "FIRSTNAME",
						Length: 25,
					},
					{
						Name:   "LASTNAME",
						Length: 25,
					},
				},
			},
		},
		{
			name:     "nested schema",
			filename: "schema-nested.yaml",
			expected: config.Config{
				Schema: config.Schema{
					{
						Name:   "BIRTHDATE",
						Length: 8,
						Schema: config.Either[string, config.Schema]{
							T2: &config.Schema{
								{
									Name:   "YEAR",
									Length: 4,
								},
								{
									Name:   "MONTH",
									Length: 2,
								},
								{
									Name:   "DAY",
									Length: 2,
								},
							},
						},
					},
				},
			},
		},
		{
			name:     "embedded schema",
			filename: "schema-embedded.yaml",
			expected: config.Config{
				Schema: config.Schema{
					{
						Name:     "PERSON",
						Length:   50,
						Feedback: true,
						Schema: config.Either[string, config.Schema]{
							T2: &config.Schema{
								{
									Name:   "FIRSTNAME",
									Length: 25,
								},
								{
									Name:   "LASTNAME",
									Length: 25,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			schema, err := config.LoadConfigFromFile("testdata/" + testCase.filename)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			// Compare the loaded schema with the expected schema
			assert.DeepEqual(t, schema, testCase.expected)
		})
	}
}
