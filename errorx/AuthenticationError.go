package errorx

import (
	"errors"
	"fmt"
)

type AuthenticationError interface {
	error
}

func NewAuthenticationError(message string) AuthenticationError {
	return AuthenticationError(errors.New(message))
}

func NewAuthenticationErrorf(format string, args ...interface{}) AuthenticationError {
	return AuthenticationError(errors.New(fmt.Sprintf(format, args...)))
}
