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

//nolint:funlen
func TestComp3_Decode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		data      []byte
		intDigits int
		decDigits int
		expected  any
		wantErr   bool
	}{
		{
			name:      "positive value",
			data:      []byte{0x12, 0x34, 0x56, 0x7C},
			intDigits: 5,
			decDigits: 2,
			expected:  "+12345.67",
			wantErr:   false,
		},
		{
			name:      "negative value",
			data:      []byte{0x12, 0x34, 0x56, 0x7D},
			intDigits: 5,
			decDigits: 2,
			expected:  "-12345.67",
			wantErr:   false,
		},
		{
			name:      "zero sign",
			data:      []byte{0x00, 0x00, 0x00, 0x0F},
			intDigits: 5,
			decDigits: 2,
			expected:  "00000.00",
			wantErr:   false,
		},
		{
			name:      "invalid sign nibble",
			data:      []byte{0x12, 0x34, 0x56, 0x7A},
			intDigits: 5,
			decDigits: 2,
			expected:  "12345.67",
			wantErr:   true,
		},
		{
			name:      "short buffer",
			data:      []byte{0x12, 0x34},
			intDigits: 3,
			decDigits: 1,
			expected:  nil,
			wantErr:   true,
		},
		{
			name:      "even number of digits",
			data:      []byte{0x12, 0x34, 0x0C},
			intDigits: 2,
			decDigits: 2,
			expected:  "+12.34",
			wantErr:   false,
		},
		{
			name:      "only decimal digits",
			data:      []byte{0x12, 0x34, 0x5C},
			intDigits: 0,
			decDigits: 5,
			expected:  "+.12345",
			wantErr:   false,
		},
		{
			name:      "decimal after 1 digit",
			data:      []byte{0x12, 0x34, 0x5C},
			intDigits: 1,
			decDigits: 4,
			expected:  "+1.2345",
			wantErr:   false,
		},
		{
			name:      "decimal after 2 digits",
			data:      []byte{0x12, 0x34, 0x5C},
			intDigits: 2,
			decDigits: 3,
			expected:  "+12.345",
			wantErr:   false,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			buf := buffer.NewBufferReader(bytes.NewReader(testcase.data))
			comp3 := codec.NewComp3(testcase.intDigits, testcase.decDigits)

			value, err := comp3.Decode(buf, 0)
			if (err != nil) != testcase.wantErr {
				t.Fatalf("[%s] expected error: %v, got: %v", testcase.name, testcase.wantErr, err)
			}

			if value != testcase.expected {
				t.Errorf("[%s] expected %s, got %s", testcase.name, testcase.expected, value)
			}
		})
	}
}
