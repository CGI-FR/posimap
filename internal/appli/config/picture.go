package config

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Picture string

type PictureType byte

const (
	PictureTypeAny     PictureType = 'X'
	PictureTypeAlpha   PictureType = 'A'
	PictureTypeNumeric PictureType = '9'
)

type PictureFormat struct {
	Type    PictureType // Type indicates the type of the picture (X, A, 9)
	Length  int         // Length indicates the length of the picture (excluding decimal places for numeric types)
	Decimal int         // Decimal indicates the number of decimal places (only applicable for numeric types)
	Signed  bool        // Signed indicates if the picture is signed (S) or not (only applicable for numeric types)
}

func (pf PictureFormat) IsValid() bool {
	// A valid picture format must have a positive length.
	if pf.Length < 0 {
		return false
	}

	// Numeric types can have decimal places, but they must be non-negative.
	if pf.Type == PictureTypeNumeric && pf.Decimal < 0 {
		return false
	}

	// Non-numeric types should not have decimal places.
	if pf.Type != PictureTypeNumeric && pf.Decimal > 0 {
		return false
	}

	// Non-numeric types should not be signed.
	if pf.Type != PictureTypeNumeric && pf.Signed {
		return false
	}

	return true
}

func (pf PictureFormat) String() string {
	signed := ""
	if pf.Signed {
		signed = "S"
	}

	decimal := ""
	if pf.Decimal > 0 {
		decimal = fmt.Sprintf("V(%d)", pf.Decimal)
	}

	return fmt.Sprintf("%s%s(%d)%s", signed, string(pf.Type), pf.Length, decimal)
}

// picturePattern is a regex pattern to match the picture format.
// It captures the signed indicator (S), type (X, A, 9), optional length, and optional decimal places.
// Captures:
// 1. Signed indicator (S)
// 2. Type (X, A, 9)
// 3. Optional length
// 4. Optional length (9+ format)
// 5. Optional decimal length
// 6. Optional decimal length (9+ format)
// The pattern allows for whitespace around the components.
const picturePattern = `^\s*(S?)\s*([X9A])\s*(?:\(\s*(\d+)\s*\)|(9+))?\s*(?:V\s*(?:\(\s*(\d+)\s*\)|(9+)))?\s*$`

var pictureRegex = regexp.MustCompile(picturePattern)

var (
	ErrInvalidPictureFormat = errors.New("invalid picture format")
	ErrInvalidPictureLength = errors.New("invalid picture length")
)

func (p Picture) Compile() (PictureFormat, error) {
	picture := PictureFormat{
		Type:    PictureTypeAny,
		Length:  0,
		Decimal: 0,
		Signed:  false,
	}

	matches := pictureRegex.FindStringSubmatch(string(p))
	if matches == nil {
		return PictureFormat{}, fmt.Errorf("%w: %s", ErrInvalidPictureFormat, p)
	}

	picture.Signed = matches[1] == "S"

	switch matches[2] {
	case "X":
		picture.Type = PictureTypeAny
	case "A":
		picture.Type = PictureTypeAlpha
	case "9":
		picture.Type = PictureTypeNumeric
	default:
		return PictureFormat{}, fmt.Errorf("%w: %s", ErrInvalidPictureFormat, p)
	}

	length, err := extractLength(matches[3], matches[4])
	if err != nil {
		return PictureFormat{}, fmt.Errorf("%w: %s (%w)", ErrInvalidPictureFormat, p, err)
	}

	picture.Length = length

	length, err = extractLength(matches[5], matches[6])
	if err != nil {
		return PictureFormat{}, fmt.Errorf("%w: %s (%w)", ErrInvalidPictureFormat, p, err)
	}

	picture.Decimal = length

	return picture, nil
}

func (p Picture) Type() PictureType {
	if len(p) == 0 {
		return PictureTypeAny
	}

	switch p[0] {
	case 'X':
		return PictureTypeAny
	case 'A':
		return PictureTypeAlpha
	case '9':
		return PictureTypeNumeric
	default:
		return PictureTypeAny
	}
}

func extractLength(directDefinition, indirectDefinition string) (int, error) {
	if directDefinition != "" {
		length, err := strconv.Atoi(directDefinition)
		if err != nil {
			return 0, fmt.Errorf("%w: %s", ErrInvalidPictureLength, directDefinition)
		}

		return length, nil
	} else {
		return len(indirectDefinition), nil
	}
}
