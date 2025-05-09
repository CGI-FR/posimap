package predicate

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/cgi-fr/posimap/pkg/posimap/api"
)

func If(value bool) api.Predicate {
	return func(_ api.Record) (bool, error) { return value, nil }
}

func Always() api.Predicate {
	return func(_ api.Record) (bool, error) { return true, nil }
}

func Never() api.Predicate {
	return func(_ api.Record) (bool, error) { return false, nil }
}

func When(tmpl string) api.Predicate {
	if tmpl == "" {
		return nil
	}

	template, err := template.New("predicate").Funcs(sprig.TxtFuncMap()).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	return func(root api.Record) (bool, error) {
		var result bytes.Buffer

		if err := template.Execute(&result, root.AsPrimitive()); err != nil {
			return false, fmt.Errorf("%w", err)
		}

		return result.String() != "false", nil
	}
}
