package data_test

import (
	"testing"

	"github.com/cgi-fr/posch/pkg/data"
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
	template := data.RecordTemplate{
		data.FieldTemplate{Name: "IDENTIFICATION", Length: 16, When: data.Never()},
		data.FieldTemplate{Name: "PERSON", Redefine: "IDENTIFICATION", When: data.When(`{{ .ISCOMPANY | ne "1" }}`), Template: data.RecordTemplate{ //nolint:lll
			data.FieldTemplate{Name: "FIRSTNAME", Length: 8},
			data.FieldTemplate{Name: "LASTNAME", Length: 8},
		}},
		data.FieldTemplate{Name: "COMPANY", Redefine: "IDENTIFICATION", When: data.When(`{{ .ISCOMPANY | eq "1" }}`), Template: data.RecordTemplate{ //nolint:lll
			data.FieldTemplate{Name: "NAME", Length: 16},
		}},
		data.FieldTemplate{Name: "ADDRESSES", Occurs: 2, Template: data.RecordTemplate{
			data.FieldTemplate{Name: "LINE-1", Length: 25},
			data.FieldTemplate{Name: "LINE-2", Length: 25},
		}},
		data.FieldTemplate{Name: "ISCOMPANY", Length: 1},
		data.FieldTemplate{Name: "TITLES", Occurs: 4, Length: 2},
	}

	builder := data.NewBuilder()
	view := builder.Build(template)

	buffer := data.Buffer("JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704    56 MAPLE AVENUE          RIVERSIDE, CA 92501      0DRPR    ") //nolint:lll
	record := data.NewRecord(buffer, view)

	record.Export(&MockRecordSink{}) //nolint:errcheck
}
