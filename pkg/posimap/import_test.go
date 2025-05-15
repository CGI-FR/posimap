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

package posimap_test

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/cgi-fr/posimap/internal/infra/jsonline"
	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/record"
	"golang.org/x/text/encoding/charmap"
)

func ExampleRecord_Import() {
	buffer := buffer.NewBufferWriter(os.Stdout)

	rec := record.NewObject()
	rec.Add("NAME", record.NewValue(0, codec.NewString(charmap.ISO8859_1, 35, true)), nil)
	rec.Add("ADDRESS-LINE1", record.NewValue(35, codec.NewString(charmap.ISO8859_1, 30, true)), nil)
	rec.Add("ADDRESS-LINE2", record.NewValue(65, codec.NewString(charmap.ISO8859_1, 30, true)), nil)

	data := strings.NewReader(`{"NAME":"JOHN DOE","ADDRESS-LINE1":"1234 ELM STREET","ADDRESS-LINE2":"SPRINGFIELD, IL 62704"}`) //nolint:lll
	reader := jsonline.NewReader(data)

	document, err := reader.Read()
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	if err := rec.Import(document); err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	if err := rec.Marshal(buffer); err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	if err := buffer.Reset(0, 0); err != nil {
		panic(err)
	}

	// Output:
	// JOHN DOE                           1234 ELM STREET               SPRINGFIELD, IL 62704
}
