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

package charsets

import (
	"errors"
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

const (
	IBM00037          = "IBM_037"
	IBM00437          = "IBM_437"
	IBM00850          = "IBM_850"
	IBM00852          = "IBM_852"
	IBM00855          = "IBM_855"
	IBM00858          = "IBM_858"
	IBM00860          = "IBM_860"
	IBM00862          = "IBM_862"
	IBM00863          = "IBM_863"
	IBM00865          = "IBM_865"
	IBM00866          = "IBM_866"
	IBM01047          = "IBM_1047"
	IBM01140          = "IBM_1140"
	ISO88591          = "ISO8859_1"
	ISO88592          = "ISO8859_2"
	ISO88593          = "ISO8859_3"
	ISO88594          = "ISO8859_4"
	ISO88595          = "ISO8859_5"
	ISO88596          = "ISO8859_6"
	ISO88597          = "ISO8859_7"
	ISO88598          = "ISO8859_8"
	ISO88599          = "ISO8859_9"
	ISO885910         = "ISO8859_10"
	ISO885913         = "ISO8859_13"
	ISO885914         = "ISO8859_14"
	ISO885915         = "ISO8859_15"
	ISO885916         = "ISO8859_16"
	KOI8R             = "KOI8R"
	KOI8U             = "KOI8U"
	MACINTOSH         = "Macintosh"
	MACINTOSHCYRILLIC = "MacintoshCyrillic"
	WINDOWS874        = "Windows874"
	WINDOWS1250       = "Windows1250"
	WINDOWS1251       = "Windows1251"
	WINDOWS1252       = "Windows1252"
	WINDOWS1253       = "Windows1253"
	WINDOWS1254       = "Windows1254"
	WINDOWS1255       = "Windows1255"
	WINDOWS1256       = "Windows1256"
	WINDOWS1257       = "Windows1257"
	WINDOWS1258       = "Windows1258"
)

var ErrCharsetNotFound = errors.New("charset not found")

func List() []string {
	return []string{
		IBM00037,
		IBM00437,
		IBM00850,
		IBM00852,
		IBM00855,
		IBM00858,
		IBM00860,
		IBM00862,
		IBM00863,
		IBM00865,
		IBM00866,
		IBM01047,
		IBM01140,
		ISO88591,
		ISO88592,
		ISO88593,
		ISO88594,
		ISO88595,
		ISO88596,
		ISO88597,
		ISO88598,
		ISO88599,
		ISO885910,
		ISO885913,
		ISO885914,
		ISO885915,
		ISO885916,
		KOI8R,
		KOI8U,
		MACINTOSH,
		MACINTOSHCYRILLIC,
		WINDOWS874,
		WINDOWS1250,
		WINDOWS1251,
		WINDOWS1252,
		WINDOWS1253,
		WINDOWS1254,
		WINDOWS1255,
		WINDOWS1256,
		WINDOWS1257,
		WINDOWS1258,
	}
}

//nolint:gocyclo,cyclop,funlen
func Get(name string) (*charmap.Charmap, error) {
	switch name {
	case IBM00037:
		return charmap.CodePage037, nil
	case IBM00437:
		return charmap.CodePage437, nil
	case IBM00850:
		return charmap.CodePage850, nil
	case IBM00852:
		return charmap.CodePage852, nil
	case IBM00855:
		return charmap.CodePage855, nil
	case IBM00858:
		return charmap.CodePage858, nil
	case IBM00860:
		return charmap.CodePage860, nil
	case IBM00862:
		return charmap.CodePage862, nil
	case IBM00863:
		return charmap.CodePage863, nil
	case IBM00865:
		return charmap.CodePage865, nil
	case IBM00866:
		return charmap.CodePage866, nil
	case IBM01047:
		return charmap.CodePage1047, nil
	case IBM01140:
		return charmap.CodePage1140, nil
	case ISO88591:
		return charmap.ISO8859_1, nil
	case ISO88592:
		return charmap.ISO8859_2, nil
	case ISO88593:
		return charmap.ISO8859_3, nil
	case ISO88594:
		return charmap.ISO8859_4, nil
	case ISO88595:
		return charmap.ISO8859_5, nil
	case ISO88596:
		return charmap.ISO8859_6, nil
	case ISO88597:
		return charmap.ISO8859_7, nil
	case ISO88598:
		return charmap.ISO8859_8, nil
	case ISO88599:
		return charmap.ISO8859_9, nil
	case ISO885910:
		return charmap.ISO8859_10, nil
	case ISO885913:
		return charmap.ISO8859_13, nil
	case ISO885914:
		return charmap.ISO8859_14, nil
	case ISO885915:
		return charmap.ISO8859_15, nil
	case ISO885916:
		return charmap.ISO8859_16, nil
	case KOI8R:
		return charmap.KOI8R, nil
	case KOI8U:
		return charmap.KOI8U, nil
	case MACINTOSH:
		return charmap.Macintosh, nil
	case MACINTOSHCYRILLIC:
		return charmap.MacintoshCyrillic, nil
	case WINDOWS874:
		return charmap.Windows874, nil
	case WINDOWS1250:
		return charmap.Windows1250, nil
	case WINDOWS1251:
		return charmap.Windows1251, nil
	case WINDOWS1252:
		return charmap.Windows1252, nil
	case WINDOWS1253:
		return charmap.Windows1253, nil
	case WINDOWS1254:
		return charmap.Windows1254, nil
	case WINDOWS1255:
		return charmap.Windows1255, nil
	case WINDOWS1256:
		return charmap.Windows1256, nil
	case WINDOWS1257:
		return charmap.Windows1257, nil
	case WINDOWS1258:
		return charmap.Windows1258, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrCharsetNotFound, name)
	}
}
