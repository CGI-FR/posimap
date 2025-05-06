package predicate

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/cgi-fr/posimap/refonte/api"
)

func If(value bool) api.Predicate {
	return func(_ api.Record) bool { return value }
}

func Always() api.Predicate {
	return func(_ api.Record) bool { return true }
}

func Never() api.Predicate {
	return func(_ api.Record) bool { return false }
}

func When(tmpl string) api.Predicate {
	if tmpl == "" {
		return nil
	}

	template, err := template.New("predicate").Funcs(sprig.TxtFuncMap()).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	return func(root api.Record) bool {
		var result bytes.Buffer
		_ = template.Execute(&result, root.AsPrimitive())

		return result.String() != "false"
	}
}
