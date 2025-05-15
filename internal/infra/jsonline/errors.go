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

package jsonline

import "errors"

var (
	ErrTokenNeedValue = errors.New("token need a value")
	ErrUnknownToken   = errors.New("unknown token")
	ErrUnknownDelim   = errors.New("unknown delimiter")
	ErrUnexpectedType = errors.New("unexpected type")
	ErrInvalidNumber  = errors.New("invalid number")
)
