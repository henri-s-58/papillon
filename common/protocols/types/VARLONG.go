package types

import (
	"bytes"
	"papillon/support"
)

type FieldVarlong struct {
	Field
}

func NewFieldVarlong(name string, docString string) (*FieldVarlong, SchemaError) {
	f, err := NewField(name, VARLONG, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldVarlong{*f}, nil
}

func NewFieldVarlongWithDefault(name string, docString string, defaultValue int64) (*FieldVarlong, SchemaError) {
	f, err := NewField(name, VARLONG, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldVarlong{*f}, nil
}

var VARLONG DocumentedTyp = TypVARLONG{}

type TypVARLONG struct {
}

func (t TypVARLONG) String() string {
	return t.TypName()
}

func (t TypVARLONG) Documentation() string {
	return `Represents an integer between -2<sup>63</sup> and 2<sup>63</sup>-1 inclusive.
Encoding follows the variable-length zig-zag encoding from
<a href=\"http://code.google.com/apis/protocolbuffers/docs/encoding.html\"> Google Protocol Buffers</a>.`
}

func (t TypVARLONG) TypName() string {
	return "VARLONG"
}

func (t TypVARLONG) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.(int64); ok {
		return b, nil
	}
	if b, ok := i.(uint64); ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a varlong.", i)
}

func (t TypVARLONG) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	if v, err := t.Validate(i); err != nil {
		return err
	} else {
		support.WriteVarlong(v.(int64), buf)
	}
	return nil
}

func (t TypVARLONG) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	i, err := support.ReadVarlong(buf)
	if err != nil {
		return nil, NewSchemaError(err.Error())
	}
	return i, nil
}

func (t TypVARLONG) SizeOf(i interface{}) int {
	return support.SizeOfVarlong(i.(int64))
}

func (t TypVARLONG) IsNilable() bool {
	return false
}

func (t TypVARLONG) ArrayElementType() Typ {
	return nil
}

func (t TypVARLONG) IsArray() bool {
	return false
}
