package types

import (
	"bytes"
	"math"
	"papillon/support"
)

type FieldCompactString struct {
	Field
}

func NewFieldCompactString(name string, docString string) (*FieldCompactString, SchemaError) {
	f, err := NewField(name, COMPACT_STRING, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldCompactString{*f}, nil
}

func NewFieldCompactStringWithDefault(name string, docString string, defaultValue string) (*FieldCompactString, SchemaError) {
	f, err := NewField(name, COMPACT_STRING, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldCompactString{*f}, nil
}

var COMPACT_STRING DocumentedTyp = TypCOMPACT_STRING{}

type TypCOMPACT_STRING struct {
}

func (t TypCOMPACT_STRING) String() string {
	return t.TypName()
}

func (t TypCOMPACT_STRING) Documentation() string {
	return `Represents a sequence of characters. First the length N + 1 is given as an UNSIGNED_VARINT 
. Then N bytes follow which are the UTF-8 encoding of the character sequence.`
}

func (t TypCOMPACT_STRING) TypName() string {
	return "COMPACT_STRING"
}

func (t TypCOMPACT_STRING) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.(string); ok {
		return b, nil
	}
	if b, ok := i.([]byte); ok {
		return string(b), nil
	}
	return nil, NewSchemaErrorf("%v is not a string.", i)
}

func (t TypCOMPACT_STRING) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	s := []byte(v.(string))
	if len(s) > math.MaxInt16 {
		return NewSchemaErrorf("String length %d is larger than the maximum string length.", len(s))
	}
	support.WriteUnsignedVarint(uint(len(s)+1), buf)
	_, err = buf.Write(s)
	if err != nil {
		return err
	}
	return nil
}

func (t TypCOMPACT_STRING) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	ru, err := support.ReadUnsignedVarint(buf)
	if err != nil {
		return nil, NewSchemaErrorf("Invalid Compact String length. %v", err)
	}
	length := ru - 1
	if length < 0 {
		return nil, NewSchemaErrorf("String length %d cannot be negative", length)
	}
	if length > math.MaxInt16 {
		return nil, NewSchemaErrorf("String length %d, is larger than the maximum string length.", length)
	}
	remainLen := buf.Len()
	if length > uint(remainLen) {
		return nil, NewSchemaErrorf("Error reading string of length %d, only %d bytes available", length, remainLen)
	}
	bs := buf.Bytes()
	if _, err := t.Validate(bs); err != nil {
		return nil, err
	}
	return string(bs), nil
}

func (t TypCOMPACT_STRING) SizeOf(i interface{}) int {
	l := len(i.(string))
	return support.SizeOfUnsignedVarint(l+1) + l
}

func (t TypCOMPACT_STRING) IsNullable() bool {
	return false
}

func (t TypCOMPACT_STRING) ArrayElementType() Typ {
	return nil
}

func (t TypCOMPACT_STRING) IsArray() bool {
	return false
}
