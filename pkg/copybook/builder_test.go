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

package copybook_test

import (
	"testing"

	"github.com/cgi-fr/posimap/pkg/copybook"
	"github.com/cgi-fr/posimap/pkg/data2"
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
	copyb := copybook.RecordSchema{
		copybook.FieldSchema{Name: "IDENTIFICATION", Length: 16, When: data2.Never()},
		copybook.FieldSchema{Name: "PERSON", Redefine: "IDENTIFICATION", When: data2.When(`{{ .ISCOMPANY | ne "1" }}`), Schema: copybook.RecordSchema{ //nolint:lll
			copybook.FieldSchema{Name: "FIRSTNAME", Length: 8},
			copybook.FieldSchema{Name: "LASTNAME", Length: 8},
		}},
		copybook.FieldSchema{Name: "COMPANY", Redefine: "IDENTIFICATION", When: data2.When(`{{ .ISCOMPANY | eq "1" }}`), Schema: copybook.RecordSchema{ //nolint:lll
			copybook.FieldSchema{Name: "NAME", Length: 16},
		}},
		copybook.FieldSchema{Name: "ADDRESSES", Occurs: 2, Schema: copybook.RecordSchema{
			copybook.FieldSchema{Name: "LINE-1", Length: 25},
			copybook.FieldSchema{Name: "LINE-2", Length: 25},
		}},
		copybook.FieldSchema{Name: "ISCOMPANY", Length: 1},
		copybook.FieldSchema{Name: "TITLES", Occurs: 4, Length: 2},
	}

	builder := copybook.NewBuilder()
	schema := builder.Build(copyb)

	buffer := data2.NewBufferFrom("JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704    56 MAPLE AVENUE          RIVERSIDE, CA 92501      0DRPR    ") //nolint:lll

	record, err := schema.CreateRecord(buffer, 0)
	if err != nil {
		t.Fatalf("failed to create record: %v", err)
	}

	data2.Export(record, &MockRecordSink{}) //nolint:errcheck
}
