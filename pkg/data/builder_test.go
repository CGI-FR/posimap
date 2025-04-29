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
