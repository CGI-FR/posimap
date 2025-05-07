package jsonline

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

type Reader struct {
	decoder *json.Decoder
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		decoder: json.NewDecoder(bufio.NewReader(reader)),
	}
}

func (r *Reader) ReadToken() (document.Token, error) {
	token, _, err := r.ReadValue()
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	return token, nil
}

func (r *Reader) ReadValue() (document.Token, any, error) {
	token, err := r.decoder.Token()
	if err != nil {
		return 0, nil, fmt.Errorf("%w", err)
	}

	if value, ok := token.(json.Delim); ok {
		switch value {
		case '{':
			return document.TokenObjStart, nil, nil
		case '}':
			return document.TokenObjEnd, nil, nil
		case '[':
			return document.TokenArrStart, nil, nil
		case ']':
			return document.TokenArrEnd, nil, nil
		}
	}

	return r.convert(token)
}

func (r *Reader) ReadString() (string, error) {
	token, err := r.decoder.Token()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if str, ok := token.(string); ok {
		return str, nil
	}

	return "", fmt.Errorf("%w: %v", ErrUnexpectedType, token)
}

func (r *Reader) ReadNumber() (float64, error) {
	token, err := r.decoder.Token()
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	if num, ok := token.(json.Number); ok {
		if f, err := num.Float64(); err == nil {
			return f, nil
		}
	}

	return 0, fmt.Errorf("%w: %v", ErrUnexpectedType, token)
}

func (r *Reader) ReadBool() (bool, error) {
	token, err := r.decoder.Token()
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}

	if b, ok := token.(bool); ok {
		return b, nil
	}

	return false, fmt.Errorf("%w: %v", ErrUnexpectedType, token)
}

func (r *Reader) ReadNull() error {
	token, err := r.decoder.Token()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if token == nil {
		return nil
	}

	return fmt.Errorf("%w: %v", ErrUnexpectedType, token)
}

//nolint:cyclop
func (r *Reader) convert(token json.Token) (document.Token, any, error) {
	switch value := token.(type) {
	case string:
		return document.TokenString, value, nil
	case json.Number:
		if f, err := value.Float64(); err == nil {
			return document.TokenNumber, f, nil
		}

		return 0, nil, fmt.Errorf("%w: %v", ErrInvalidNumber, value)
	case bool:
		if value {
			return document.TokenTrue, value, nil
		}

		return document.TokenFalse, value, nil
	case nil:
		return document.TokenNull, nil, nil
	case json.Delim:
		switch value {
		case '{':
			return document.TokenObjStart, nil, nil
		case '}':
			return document.TokenObjEnd, nil, nil
		case '[':
			return document.TokenArrStart, nil, nil
		case ']':
			return document.TokenArrEnd, nil, nil
		default:
			return 0, nil, fmt.Errorf("%w: %v", ErrUnknownDelim, value)
		}
	default:
		return 0, nil, fmt.Errorf("%w: %T", ErrUnknownToken, token)
	}
}
