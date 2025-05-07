package config

type Config struct {
	Schema Schema `yaml:"schema"`
}

type Schema []Field

type Field struct {
	Name     string `yaml:"name"`
	Occurs   int    `yaml:"occurs,omitempty"`
	Redefine string `yaml:"redefine,omitempty"`
	When     string `yaml:"when,omitempty"`

	Length  int    `yaml:"length"`
	Trim    bool   `yaml:"trim,omitempty"`
	Charset string `yaml:"charset,omitempty"`

	Schema Either[string, Schema] `yaml:"schema"` // either filename for external schema or embedded schema
}
