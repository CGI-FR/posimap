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

package data_test

import (
	"testing"

	"github.com/cgi-fr/posimap/pkg/data"
)

type MockRecordSink struct{}

func (m *MockRecordSink) OpenRecord() error {
	return nil
}

func (m *MockRecordSink) CloseRecord() error {
	println()

	return nil
}

func (m *MockRecordSink) OpenObject() error {
	print("{")

	return nil
}

func (m *MockRecordSink) CloseObject() error {
	print("}")

	return nil
}

func (m *MockRecordSink) OpenArray() error {
	print("[")

	return nil
}

func (m *MockRecordSink) CloseArray() error {
	print("]")

	return nil
}

func (m *MockRecordSink) WriteString(data string) error {
	print("\"", data, "\"")

	return nil
}

func (m *MockRecordSink) WriteKey(key string) error {
	m.WriteString(key) //nolint:errcheck
	print(":")

	return nil
}

func (m *MockRecordSink) Next() error {
	print(",")

	return nil
}

func (m *MockRecordSink) Close() error {
	return nil
}

func TestBuilder(t *testing.T) {
	t.Parallel()

	//nolint:exhaustruct
	schema := data.RecordSchema{
		data.FieldSchema{Name: "IDENTIFICATION", Length: 16, When: data.Never()},
		data.FieldSchema{Name: "PERSON", Redefine: "IDENTIFICATION", When: data.When(`{{ .ISCOMPANY | ne "1" }}`), Schema: data.RecordSchema{ //nolint:lll
			data.FieldSchema{Name: "FIRSTNAME", Length: 8},
			data.FieldSchema{Name: "LASTNAME", Length: 8},
		}},
		data.FieldSchema{Name: "COMPANY", Redefine: "IDENTIFICATION", When: data.When(`{{ .ISCOMPANY | eq "1" }}`), Schema: data.RecordSchema{ //nolint:lll
			data.FieldSchema{Name: "NAME", Length: 16},
		}},
		data.FieldSchema{Name: "ADDRESSES", Occurs: 2, Schema: data.RecordSchema{
			data.FieldSchema{Name: "LINE-1", Length: 25},
			data.FieldSchema{Name: "LINE-2", Length: 25},
		}},
		data.FieldSchema{Name: "ISCOMPANY", Length: 1},
		data.FieldSchema{Name: "TITLES", Occurs: 4, Length: 2},
	}

	builder := data.NewBuilder()
	view := builder.Build(schema)

	buffer := data.NewBufferFrom("JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704    56 MAPLE AVENUE          RIVERSIDE, CA 92501      0DRPR    ") //nolint:lll
	record := data.NewRecord(buffer, view)

	record.Export(&MockRecordSink{}) //nolint:errcheck
}
