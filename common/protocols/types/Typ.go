package types

import "bytes"

type Typ interface {
	String() string
	Validate(i interface{}) (interface{}, SchemaError)
	Write(buf *bytes.Buffer, i interface{}) SchemaError
	Read(buf *bytes.Buffer) (interface{}, SchemaError)
	SizeOf(i interface{}) int // bytes len
	IsNullable() bool
	ArrayElementType() Typ
	IsArray() bool
}
