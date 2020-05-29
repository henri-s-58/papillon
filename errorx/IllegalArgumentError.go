package errorx

import (
	"errors"
	"fmt"
)

type IllegalArgumentError interface {
	error
}

func NewIllegalArgumentError(message string) IllegalArgumentError {
	return IllegalArgumentError(errors.New(message))
}

func NewIllegalArgumentErrorf(format string, args ...interface{}) IllegalArgumentError {
	return IllegalArgumentError(errors.New(fmt.Sprintf(format, args...)))
}
