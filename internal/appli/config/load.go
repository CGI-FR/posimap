package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromYAML(data []byte, rootpath string) (Config, error) {
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)

	var config Config
	if err := dec.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("%w", err)
	}

	if err := resolveIncludes(config.Schema, rootpath); err != nil {
		return Config{}, fmt.Errorf("failed to resolve includes: %w", err)
	}

	return config, nil
}

func LoadConfigFromFile(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("%w", err)
	}

	return LoadConfigFromYAML(data, filepath.Dir(filename))
}

// resolveIncludes parcourt l’arbre et remplace tout Schema: "file.yaml"
// par le contenu de file.yaml (parsé dans le champ schema privé).
func resolveIncludes(schema Schema, rootpath string) error {
	for idx, field := range schema {
		if field.Schema.T1 != nil {
			includedSchema, err := LoadConfigFromFile(filepath.Join(rootpath, *field.Schema.T1))
			if err != nil {
				return fmt.Errorf("failed to load included schema %s: %w", *field.Schema.T1, err)
			}
			// Replace the field schema with the included schema
			schema[idx].Schema.T2 = &includedSchema.Schema
			schema[idx].Schema.T1 = nil
		}

		if schema[idx].Schema.T2 != nil {
			// Recursively resolve includes in the nested schema
			if err := resolveIncludes(*schema[idx].Schema.T2, rootpath); err != nil {
				return fmt.Errorf("failed to resolve includes in nested schema: %w", err)
			}
		}
	}

	return nil
}
