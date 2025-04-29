package data

import (
	"bytes"
	"text/template"
)

func If(value bool) ExportPredicate {
	return func(_ View, _ *Buffer) bool { return value }
}

func Always() ExportPredicate {
	return func(_ View, _ *Buffer) bool { return true }
}

func Never() ExportPredicate {
	return func(_ View, _ *Buffer) bool { return false }
}

func When(tmpl string) ExportPredicate {
	template, err := template.New("predicate").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	return func(root View, buffer *Buffer) bool {
		var result bytes.Buffer
		_ = template.Execute(&result, root.Materialize(buffer))

		return result.String() != "false"
	}
}
