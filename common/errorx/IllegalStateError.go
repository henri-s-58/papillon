package errorx

import (
	"errors"
	"fmt"
)

type IllegalStateError interface {
	error
}

func NewIllegalStateError(message string) IllegalStateError {
	return IllegalStateError(errors.New(message))
}

func NewIllegalStateErrorf(format string, args ...interface{}) IllegalStateError {
	return IllegalStateError(errors.New(fmt.Sprintf(format, args...)))
}
