package types

import (
	"bytes"
	"encoding/binary"
)

type FieldUint32 struct {
	Field
}

func NewFieldUint32(name string, docString string) (*FieldUint32, SchemaError) {
	f, err := NewField(name, Uint32, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldUint32{*f}, nil
}

func NewFieldUint32WithDefault(name string, docString string, defaultValue uint32) (*FieldUint32, SchemaError) {
	f, err := NewField(name, Uint32, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldUint32{*f}, nil
}

var Uint32 DocumentedTyp = TypUINT32{}

type TypUINT32 struct {
}

func (t TypUINT32) String() string {
	return t.TypName()
}

func (t TypUINT32) Documentation() string {
	return "Represents an integer between 0 and 2<sup>32</sup>-1 inclusive.\nThe values are encoded using four bytes in network byte order (big-endian)."
}

func (t TypUINT32) TypName() string {
	return "UINT32"
}

func (t TypUINT32) Validate(i interface{}) (interface{}, SchemaError) {
	b, ok := i.(uint32)
	if ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a uint32.", i)
}

func (t TypUINT32) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v.(uint32))
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypUINT32) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != t.SizeOf(bs) {
		return nil, NewSchemaErrorf("%v is not a Uint32.", bs)
	}
	return binary.BigEndian.Uint32(bs), nil
}

func (t TypUINT32) SizeOf(i interface{}) int {
	return 4
}

func (t TypUINT32) IsNullable() bool {
	return false
}

func (t TypUINT32) ArrayElementType() Typ {
	return nil
}

func (t TypUINT32) IsArray() bool {
	return false
}
