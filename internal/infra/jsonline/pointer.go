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

package jsonline

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoObjectToClose  = errors.New("no object to close")
	ErrNoArrayToClose   = errors.New("no array to close")
	ErrIncompleteObject = errors.New("incomplete object")
)

type Pointer struct {
	indexes []uint
	levels  string
}

func NewPointer() *Pointer {
	return &Pointer{
		indexes: make([]uint, 1), // start with a root level
		levels:  "",
	}
}

func (p *Pointer) OpenObject() {
	p.levels += "{"
	p.indexes = append(p.indexes, 0)
}

func (p *Pointer) CloseObject() error {
	if !strings.HasSuffix(p.levels, "{") {
		return fmt.Errorf("%w", ErrNoObjectToClose)
	}

	index := p.indexes[len(p.indexes)-1]
	if index%2 == 1 {
		return fmt.Errorf("%w", ErrIncompleteObject)
	}

	p.levels = p.levels[:len(p.levels)-1]
	p.indexes = p.indexes[:len(p.indexes)-1]

	return nil
}

func (p *Pointer) OpenArray() {
	p.levels += "["
	p.indexes = append(p.indexes, 0)
}

func (p *Pointer) CloseArray() error {
	if !strings.HasSuffix(p.levels, "[") {
		return fmt.Errorf("%w", ErrNoArrayToClose)
	}

	p.levels = p.levels[:len(p.levels)-1]
	p.indexes = p.indexes[:len(p.indexes)-1]

	return nil
}

func (p *Pointer) Index() uint {
	return p.indexes[len(p.indexes)-1]
}

func (p *Pointer) Level() int {
	return len(p.levels)
}

func (p *Pointer) Shift() rune {
	index := p.indexes[len(p.indexes)-1]
	index++
	p.indexes[len(p.indexes)-1] = index

	if len(p.indexes) == 1 {
		if index > 1 {
			return '\n' // first document requires a new line
		}

		return 0
	}

	typ := p.levels[len(p.levels)-1]

	if index == 1 {
		return 0 // first element requires no separator
	}

	// objects need a colon as separator between key and value
	if typ == '{' && index%2 == 0 {
		return ':'
	}

	return ','
}
