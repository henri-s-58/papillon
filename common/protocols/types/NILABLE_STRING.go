package types

import (
	"bytes"
	"encoding/binary"
	"math"
)

type FieldNilableString struct {
	Field
}

func NewFieldNilableString(name string, docString string) (*FieldNilableString, SchemaError) {
	f, err := NewField(name, NILABLE_STRING, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldNilableString{*f}, nil
}

func NewFieldNilableStringWithDefault(name string, docString string, defaultValue string) (*FieldNilableString, SchemaError) {
	f, err := NewField(name, STRING, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldNilableString{*f}, nil
}

var NILABLE_STRING DocumentedTyp = TypNILABLE_STRING{}

type TypNILABLE_STRING struct {
}

func (t TypNILABLE_STRING) String() string {
	return t.TypName()
}

func (t TypNILABLE_STRING) Documentation() string {
	return `Represents a sequence of characters or null. For non-null strings,
first the length N is given as an INT16.
Then N bytes follow which are the UTF-8 encoding of the character sequence.
A null value is encoded with length of -1 and there are no following bytes.`
}

func (t TypNILABLE_STRING) TypName() string {
	return "NILABLE_STRING"
}

func (t TypNILABLE_STRING) Validate(i interface{}) (interface{}, SchemaError) {
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

func (t TypNILABLE_STRING) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	if i == nil {
		b := make([]byte, 2)
		i16 := int16(-1)
		u16 := uint16(i16)
		binary.BigEndian.PutUint16(b, u16)
		_, err := buf.Write(b)
		if err != nil {
			return NewSchemaError(err.Error())
		}
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
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(len(s)))
	b = append(b, s...)
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypNILABLE_STRING) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	lengthBytes := buf.Next(2)
	length := int16(binary.BigEndian.Uint16(lengthBytes))
	if length < 0 {
		return nil, nil
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

func (t TypNILABLE_STRING) SizeOf(i interface{}) int {
	if i == nil {
		return 2
	}
	return 2 + len(i.(string))
}

func (t TypNILABLE_STRING) IsNilable() bool {
	return true
}

func (t TypNILABLE_STRING) ArrayElementType() Typ {
	return nil
}

func (t TypNILABLE_STRING) IsArray() bool {
	return false
}
