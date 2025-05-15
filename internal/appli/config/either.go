// Copyright (C) 2025 CGI France
//
// This file is part of posimap.
//
// posimap is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// posimap is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with posimap.  If not, see <http://www.gnu.org/licenses/>.

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
