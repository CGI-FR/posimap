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

package config_test

import (
	"testing"

	"github.com/cgi-fr/posimap/internal/appli/charsets"
	"github.com/cgi-fr/posimap/internal/appli/config"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/text/encoding/charmap"
	"gotest.tools/assert"
)

func TestCompile(t *testing.T) {
	t.Parallel()

	//nolint:exhaustruct
	tests := []struct {
		name     string
		schema   config.Schema
		expected *schema.Record
	}{
		{
			name: "Simple schema",
			schema: config.Schema{
				config.Field{
					Name:   "field1",
					Length: 10,
				},
			},
			expected: schema.NewRecord("ROOT").WithField("field1", codec.NewString(charmap.ISO8859_1, 10, true)),
		},
		{
			name: "Nested schema",
			schema: config.Schema{
				config.Field{
					Name: "lvel1",
					Schema: config.Either[string, config.Schema]{
						T2: &config.Schema{
							{
								Name:   "lvel2",
								Length: 10,
							},
						},
					},
				},
			},
			expected: schema.NewRecord("ROOT").WithRecord("lvel1",
				schema.NewRecord("ROOT").WithField("lvel2", codec.NewString(charmap.ISO8859_1, 10, true)),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := test.schema.Compile(config.Trim(true), config.Charset(charsets.ISO88591))
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			// Compare the loaded schema with the expected schema
			assert.DeepEqual(t, result, test.expected,
				cmp.AllowUnexported(schema.Field{}, schema.Record{}, codec.String{}, charmap.Charmap{}),
				cmpopts.IgnoreFields(charmap.Charmap{}, "decode"),
				cmpopts.IgnoreUnexported(schema.Record{}),
			)
		})
	}
}
