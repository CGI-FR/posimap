package config

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Either[T1 any, T2 any] struct {
	T1 *T1
	T2 *T2
}

func (e *Either[T1, T2]) UnmarshalYAML(value *yaml.Node) error {
	out, err := yaml.Marshal(value)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	dec := yaml.NewDecoder(bytes.NewReader(out))
	dec.KnownFields(true)

	t1 := new(T1)

	err = dec.Decode(&t1)
	if err == nil {
		e.T1 = t1

		return nil
	}

	dec = yaml.NewDecoder(bytes.NewReader(out))
	dec.KnownFields(true)

	t2 := new(T2)

	err = dec.Decode(&t2)
	if err == nil {
		e.T2 = t2

		return nil
	}

	return fmt.Errorf("%w", err)
}
