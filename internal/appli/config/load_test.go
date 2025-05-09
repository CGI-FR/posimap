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
						Name:   "PERSON",
						Length: 50,
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
