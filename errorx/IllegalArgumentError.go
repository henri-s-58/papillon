package errorx

import (
	"fmt"

	"github.com/henri-s-58/jalmot/v1"
)

type IllegalArgumentError interface {
	error
}

func NewIllegalArgumentError(message string) IllegalArgumentError {
	return IllegalArgumentError(jalmot.New(message))
}

func NewIllegalArgumentErrorf(format string, args ...interface{}) IllegalArgumentError {
	return IllegalArgumentError(jalmot.New(fmt.Sprintf(format, args...)))
}
