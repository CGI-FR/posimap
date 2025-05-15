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
