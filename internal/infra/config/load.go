package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadTemplateFromYAML(data []byte) (Template, error) {
	// Unmarshal the YAML data into a Template struct
	var template Template
	if err := yaml.Unmarshal(data, &template); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	// Validate the template
	if err := template.Validate(); err != nil {
		return nil, fmt.Errorf("template validation failed: %w", err)
	}

	// Return the loaded template
	return template, nil
}

func LoadTemplateFromFile(filename string) (Template, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return LoadTemplateFromYAML(data)
}
