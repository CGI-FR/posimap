package config_test

import (
	"testing"

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
		expected schema.Record
	}{
		{
			name: "Simple schema",
			schema: config.Schema{
				config.Field{
					Name:   "field1",
					Length: 10,
				},
			},
			expected: schema.NewSchema().WithField("field1", codec.NewString(charmap.ISO8859_1, 10, true)),
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
			expected: schema.NewSchema().WithRecord("lvel1",
				schema.NewSchema().WithField("lvel2", codec.NewString(charmap.ISO8859_1, 10, true)),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := test.schema.Compile(config.Trim(true))

			// Compare the loaded schema with the expected schema
			assert.DeepEqual(t, result, test.expected,
				cmp.AllowUnexported(schema.Field{}, schema.Definition{}, codec.String{}, charmap.Charmap{}),
				cmpopts.IgnoreFields(charmap.Charmap{}, "decode"),
			)
		})
	}
}
