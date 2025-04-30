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

package data

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
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
	template, err := template.New("predicate").Funcs(sprig.TxtFuncMap()).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	return func(root View, buffer *Buffer) bool {
		var result bytes.Buffer
		_ = template.Execute(&result, root.Materialize(buffer))

		return result.String() != "false"
	}
}
