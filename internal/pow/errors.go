package pow

import (
	"errors"
)

var (
	ErrSolveFail  = errors.New("failed to solve the challenge")
	ErrInvalid    = errors.New("invalid")
	ErrUnexpected = errors.New("unexpected")
)
