package types

import (
	"bytes"
	"encoding/binary"
)

type FieldInt64 struct {
	Field
}

func NewFieldInt64(name string, docString string) (*FieldInt64, SchemaError) {
	f, err := NewField(name, INT64, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldInt64{*f}, nil
}

func NewFieldInt64WithDefault(name string, docString string, defaultValue int64) (*FieldInt64, SchemaError) {
	f, err := NewField(name, INT64, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldInt64{*f}, nil
}

var INT64 DocumentedTyp = TypINT64{}

type TypINT64 struct {
}

func (t TypINT64) String() string {
	return t.TypName()
}

func (t TypINT64) Documentation() string {
	return "Represents an integer between -2<sup>63</sup> and 2<sup>63</sup>-1 inclusive.\nThe values are encoded using four bytes in network byte order (big-endian)."
}

func (t TypINT64) TypName() string {
	return "INT64"
}

func (t TypINT64) Validate(i interface{}) (interface{}, SchemaError) {
	b, ok := i.(int64)
	if ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a int64.", i)
}

func (t TypINT64) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v.(int64)))
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypINT64) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != t.SizeOf(bs) {
		return nil, NewSchemaErrorf("%v is not a int64.", bs)
	}
	return int64(binary.BigEndian.Uint64(bs)), nil
}

func (t TypINT64) SizeOf(i interface{}) int {
	return 8
}

func (t TypINT64) IsNullable() bool {
	return false
}

func (t TypINT64) ArrayElementType() Typ {
	return nil
}

func (t TypINT64) IsArray() bool {
	return false
}
