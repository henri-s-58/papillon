package types

import (
	"bytes"
	"math"
	"papillon/support"
)

type FieldCompactNilableString struct {
	Field
}

func NewFieldCompactNilableString(name string, docString string) (*FieldCompactNilableString, SchemaError) {
	f, err := NewField(name, COMPACT_NILABLE_STRING, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldCompactNilableString{*f}, nil
}

func NewFieldCompactNilableStringWithDefault(name string, docString string, defaultValue string) (*FieldCompactNilableString, SchemaError) {
	f, err := NewField(name, COMPACT_NILABLE_STRING, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldCompactNilableString{*f}, nil
}

var COMPACT_NILABLE_STRING DocumentedTyp = TypCOMPACT_NILABLE_STRING{}

type TypCOMPACT_NILABLE_STRING struct {
}

func (t TypCOMPACT_NILABLE_STRING) String() string {
	return t.TypName()
}

func (t TypCOMPACT_NILABLE_STRING) Documentation() string {
	return `Represents a sequence of characters.
First the length N + 1 is given as an UNSIGNED_VARINT.
Then N bytes follow which are the UTF-8 encoding of the character sequence.
A null string is represented with a length of 0.`
}

func (t TypCOMPACT_NILABLE_STRING) TypName() string {
	return "COMPACT_NILABLE_STRING"
}

func (t TypCOMPACT_NILABLE_STRING) Validate(i interface{}) (interface{}, SchemaError) {
	if i == nil {
		return nil, nil
	}
	if b, ok := i.(string); ok {
		return b, nil
	}
	if b, ok := i.([]byte); ok {
		return string(b), nil
	}
	return nil, NewSchemaErrorf("%v is not a string.", i)
}

func (t TypCOMPACT_NILABLE_STRING) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	if i == nil {
		support.WriteUnsignedVarint(uint(0), buf)
		return nil
	}
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

func (t TypCOMPACT_NILABLE_STRING) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	ru, err := support.ReadUnsignedVarint(buf)
	if err != nil {
		return nil, NewSchemaErrorf("Invalid Compact String length. %v", err)
	}
	if ru == 0 {
		return nil, nil
	}
	length := ru - 1
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

func (t TypCOMPACT_NILABLE_STRING) SizeOf(i interface{}) int {
	if i == nil {
		return 1
	}
	l := len(i.(string))
	return support.SizeOfUnsignedVarint(uint(l+1)) + l
}

func (t TypCOMPACT_NILABLE_STRING) IsNilable() bool {
	return true
}

func (t TypCOMPACT_NILABLE_STRING) ArrayElementType() Typ {
	return nil
}

func (t TypCOMPACT_NILABLE_STRING) IsArray() bool {
	return false
}
