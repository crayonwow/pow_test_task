package errors

import (
	"errors"
)

var (
	ErrInvalid    = errors.New("invalid")
	ErrUnexpected = errors.New("unexpected")
)
