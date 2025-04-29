package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromYAML(data []byte) (Config, error) {
	// Unmarshal the YAML data into a Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML: %w", err)
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

	return LoadConfigFromYAML(data)
}
