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

package codec_test

import (
	"bytes"
	"testing"

	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
)

func TestComp3_Decode(t *testing.T) {
	t.Parallel()

	data := bytes.NewReader([]byte{0x12, 0x34, 0x56, 0x7C})
	buffer := buffer.NewBufferReader(data)

	comp3 := codec.NewComp3(5, 2)

	value, err := comp3.Decode(buffer, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "+12345.67"
	if value != expected {
		t.Errorf("expected %s, got %s", expected, value)
	}
}
