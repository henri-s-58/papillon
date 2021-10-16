package errorx

import (
	"fmt"

	"github.com/henri-s-58/jalmot/v1"
)

type IllegalStateError interface {
	error
}

func NewIllegalStateError(message string) IllegalStateError {
	return IllegalStateError(jalmot.New(message))
}

func NewIllegalStateErrorf(format string, args ...interface{}) IllegalStateError {
	return IllegalStateError(jalmot.New(fmt.Sprintf(format, args...)))
}
