package errorx

import (
	"fmt"

	"github.com/henri-s-58/jalmot/v1"
)

type AuthenticationError interface {
	error
}

func NewAuthenticationError(message string) AuthenticationError {
	return AuthenticationError(jalmot.New(message))
}

func NewAuthenticationErrorf(format string, args ...interface{}) AuthenticationError {
	return AuthenticationError(jalmot.New(fmt.Sprintf(format, args...)))
}
