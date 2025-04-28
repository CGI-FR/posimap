package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadSchemaFromYAML(data []byte) (Schema, error) {
	// Unmarshal the YAML data into a Schema struct
	var schema Schema
	if err := yaml.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	// Validate the schema
	if err := schema.Validate(); err != nil {
		return nil, fmt.Errorf("schema validation failed: %w", err)
	}

	// Return the loaded schema
	return schema, nil
}

func LoadSchemaFromFile(filename string) (Schema, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return LoadSchemaFromYAML(data)
}
