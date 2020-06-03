package types

import (
	"bytes"
	"encoding/binary"
	"math"
)

type FieldString struct {
	Field
}

func NewFieldString(name string, docString string) (*FieldString, SchemaError) {
	f, err := NewField(name, STRING, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldString{*f}, nil
}

func NewFieldStringWithDefault(name string, docString string, defaultValue string) (*FieldString, SchemaError) {
	f, err := NewField(name, STRING, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldString{*f}, nil
}

var STRING DocumentedTyp = TypSTRING{}

type TypSTRING struct {
}

func (t TypSTRING) String() string {
	return t.TypName()
}

func (t TypSTRING) Documentation() string {
	return `Represents a sequence of characters. First the length N is given as an TypINT16.
Then N bytes follow which are the UTF-8 encoding of the character sequence.
Length must not be negative.`
}

func (t TypSTRING) TypName() string {
	return "STRING"
}

func (t TypSTRING) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.(string); ok {
		return b, nil
	}
	if b, ok := i.([]byte); ok {
		return string(b), nil
	}
	return nil, NewSchemaErrorf("%v is not a string.", i)
}

func (t TypSTRING) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	s := []byte(v.(string))
	if len(s) > math.MaxInt16 {
		return NewSchemaErrorf("String length %d is larger than the maximum string length.", len(s))
	}
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(len(s)))
	b = append(b, s...)
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypSTRING) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	lengthBytes := buf.Next(2)
	length := int16(binary.BigEndian.Uint16(lengthBytes))
	if length < 0 {
		return nil, NewSchemaErrorf("String length %d cannot be negative", length)
	}
	remainLen := buf.Len()
	if int(length) > remainLen {
		return nil, NewSchemaErrorf("Error reading string of length %d, only %d bytes available", length, remainLen)
	}
	bs := buf.Bytes()
	if _, err := t.Validate(bs); err != nil {
		return nil, err
	}
	return string(bs), nil
}

func (t TypSTRING) SizeOf(i interface{}) int {
	return 2 + len(i.(string))
}

func (t TypSTRING) IsNullable() bool {
	return false
}

func (t TypSTRING) ArrayElementType() Typ {
	return nil
}

func (t TypSTRING) IsArray() bool {
	return false
}
