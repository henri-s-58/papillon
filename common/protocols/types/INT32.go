package types

import (
	"bytes"
	"encoding/binary"
)

type FieldInt32 struct {
	Field
}

func NewFieldInt32(name string, docString string) (*FieldInt32, SchemaError) {
	f, err := NewField(name, INT32, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldInt32{*f}, nil
}

func NewFieldInt32WithDefault(name string, docString string, defaultValue int32) (*FieldInt32, SchemaError) {
	f, err := NewField(name, INT32, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldInt32{*f}, nil
}

var INT32 DocumentedTyp = TypINT32{}

type TypINT32 struct {
}

func (t TypINT32) String() string {
	return t.TypName()
}

func (t TypINT32) Documentation() string {
	return "Represents an integer between -2<sup>31</sup> and 2<sup>31</sup>-1 inclusive.\nThe values are encoded using four bytes in network byte order (big-endian)."
}

func (t TypINT32) TypName() string {
	return "INT32"
}

func (t TypINT32) Validate(i interface{}) (interface{}, SchemaError) {
	b, ok := i.(int32)
	if ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a int32.", i)
}

func (t TypINT32) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v.(int32)))
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypINT32) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != t.SizeOf(bs) {
		return nil, NewSchemaErrorf("%v is not a int32.", bs)
	}
	return int32(binary.BigEndian.Uint32(bs)), nil
}

func (t TypINT32) SizeOf(i interface{}) int {
	return 4
}

func (t TypINT32) IsNilable() bool {
	return false
}

func (t TypINT32) ArrayElementType() Typ {
	return nil
}

func (t TypINT32) IsArray() bool {
	return false
}
