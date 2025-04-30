package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromYAML(data []byte, rootpath string) (Config, error) {
	// Unmarshal the YAML data into a Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	if err := resolveIncludes(config.Schema, rootpath); err != nil {
		return Config{}, fmt.Errorf("failed to resolve includes: %w", err)
	}

	// Validate the schema
	if err := config.Validate(); err != nil {
		return Config{}, fmt.Errorf("schema validation failed: %w", err)
	}

	// Return the loaded schema
	return config, nil
}

func LoadConfigFromFile(filename string) (Config, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}

	return LoadConfigFromYAML(data, filepath.Dir(filename))
}

// resolveIncludes parcourt l’arbre et remplace tout Schema: "file.yaml"
// par le contenu de file.yaml (parsé dans le champ schema privé).
func resolveIncludes(schema Schema, rootpath string) error {
	for idx, field := range schema {
		if filename, ok := field.Schema.(string); ok && filename != "" {
			// Load the included schema
			includedSchema, err := LoadConfigFromFile(filepath.Join(rootpath, filename))
			if err != nil {
				return fmt.Errorf("failed to load included schema %s: %w", filename, err)
			}
			// Replace the field schema with the included schema
			schema[idx].schema = includedSchema.Schema
			schema[idx].Schema = ""
		}

		if field.schema != nil {
			// Recursively resolve includes in the nested schema
			if err := resolveIncludes(field.schema, rootpath); err != nil {
				return fmt.Errorf("failed to resolve includes in nested schema: %w", err)
			}
		}
	}

	return nil
}
