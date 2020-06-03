package types

import (
	"bytes"
	"encoding/binary"
)

type FieldInt16 struct {
	Field
}

func NewFieldInt16(name string, docString string) (*FieldInt16, SchemaError) {
	f, err := NewField(name, INT16, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldInt16{*f}, nil
}

func NewFieldInt16WithDefault(name string, docString string, defaultValue int16) (*FieldInt16, SchemaError) {
	f, err := NewField(name, INT16, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldInt16{*f}, nil
}

var INT16 DocumentedTyp = TypINT16{}

type TypINT16 struct {
}

func (t TypINT16) String() string {
	return t.TypName()
}

func (t TypINT16) Documentation() string {
	return "Represents an integer between -2<sup>15</sup> and 2<sup>15</sup>-1 inclusive.\nThe values are encoded using two bytes in network byte order (big-endian)"
}

func (t TypINT16) TypName() string {
	return "INT16"
}

func (t TypINT16) Validate(i interface{}) (interface{}, SchemaError) {
	b, ok := i.(int16)
	if ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a int16.", i)
}

func (t TypINT16) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v.(int16)))
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypINT16) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != t.SizeOf(bs) {
		return nil, NewSchemaErrorf("%v is not a int16.", bs)
	}
	return int16(binary.BigEndian.Uint16(bs)), nil
}

func (t TypINT16) SizeOf(i interface{}) int {
	return 2
}

func (t TypINT16) IsNullable() bool {
	return false
}

func (t TypINT16) ArrayElementType() Typ {
	return nil
}

func (t TypINT16) IsArray() bool {
	return false
}
