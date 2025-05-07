package jsonline

import "errors"

var (
	ErrTokenNeedValue = errors.New("token need a value")
	ErrUnknownToken   = errors.New("unknown token")
	ErrUnknownDelim   = errors.New("unknown delimiter")
	ErrUnexpectedType = errors.New("unexpected type")
	ErrInvalidNumber  = errors.New("invalid number")
)
