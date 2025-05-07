package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadSchemaFromYAML(data []byte, _ string) (Config, error) {
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)

	var schema Config
	if err := dec.Decode(&schema); err != nil {
		return Config{}, fmt.Errorf("%w", err)
	}

	// Return the loaded schema
	return schema, nil
}

func LoadSchemaFromFile(filename string) (Config, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("%w", err)
	}

	return LoadSchemaFromYAML(data, filepath.Dir(filename))
}
