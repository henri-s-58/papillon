package types

import (
	"errors"
	"fmt"
)

type SchemaError interface {
	error
}

func NewSchemaError(message string) SchemaError {
	return SchemaError(errors.New(message))
}

func NewSchemaErrorf(format string, args ...interface{}) SchemaError {
	return SchemaError(errors.New(fmt.Sprintf(format, args...)))
}
