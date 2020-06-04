package types

import (
	"bytes"
	"papillon/support"
)

type FieldVarint struct {
	Field
}

func NewFieldVarint(name string, docString string) (*FieldVarint, SchemaError) {
	f, err := NewField(name, VARINT, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldVarint{*f}, nil
}

func NewFieldVarintWithDefault(name string, docString string, defaultValue int) (*FieldVarint, SchemaError) {
	f, err := NewField(name, VARINT, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldVarint{*f}, nil
}

var VARINT DocumentedTyp = TypVARINT{}

type TypVARINT struct {
}

func (t TypVARINT) String() string {
	return t.TypName()
}

func (t TypVARINT) Documentation() string {
	return `Represents an integer between -2<sup>31</sup> and 2<sup>31</sup>-1 inclusive.
Encoding follows the variable-length zig-zag encoding from
<a href=\"http://code.google.com/apis/protocolbuffers/docs/encoding.html\"> Google Protocol Buffers</a>.`
}

func (t TypVARINT) TypName() string {
	return "VARINT"
}

func (t TypVARINT) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.(int32); ok {
		return b, nil
	}
	if b, ok := i.(int); ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a varint.", i)
}

func (t TypVARINT) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	if v, err := t.Validate(i); err != nil {
		return err
	} else {
		support.WriteVarint(v.(int), buf)
	}
	return nil
}

func (t TypVARINT) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	i, err := support.ReadVarint(buf)
	if err != nil {
		return nil, NewSchemaError(err.Error())
	}
	return i, nil
}

func (t TypVARINT) SizeOf(i interface{}) int {
	return support.SizeOfVarint(i.(int))
}

func (t TypVARINT) IsNilable() bool {
	return false
}

func (t TypVARINT) ArrayElementType() Typ {
	return nil
}

func (t TypVARINT) IsArray() bool {
	return false
}
