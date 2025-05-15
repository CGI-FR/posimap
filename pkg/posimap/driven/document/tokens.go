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

package document

type Token rune

const (
	TokenObjStart Token = '{'
	TokenObjEnd   Token = '}'
	TokenArrStart Token = '['
	TokenArrEnd   Token = ']'
	TokenString   Token = '"' // with string value
	TokenNumber   Token = '0' // with float64 value
	TokenTrue     Token = 't' // with bool value
	TokenFalse    Token = 'f' // with bool value
	TokenNull     Token = 'n'
)
