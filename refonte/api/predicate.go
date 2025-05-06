package api

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig/v3"
)

type Predicate func(root Record) bool

func If(value bool) Predicate {
	return func(_ Record) bool { return value }
}

func Always() Predicate {
	return func(_ Record) bool { return true }
}

func Never() Predicate {
	return func(_ Record) bool { return false }
}

func When(tmpl string) Predicate {
	if tmpl == "" {
		return nil
	}

	template, err := template.New("predicate").Funcs(sprig.TxtFuncMap()).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	return func(root Record) bool {
		var result bytes.Buffer
		_ = template.Execute(&result, root.AsPrimitive())

		return result.String() != "false"
	}
}
