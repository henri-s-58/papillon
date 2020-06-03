package types

import (
	"bytes"
	"encoding/binary"
	"math"
)

type FieldFloat64 struct {
	Field
}

func NewFieldFloat64(name string, docString string) (*FieldFloat64, SchemaError) {
	f, err := NewField(name, FLOAT64, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldFloat64{*f}, nil
}

func NewFieldFloat64WithDefault(name string, docString string, defaultValue float64) (*FieldFloat64, SchemaError) {
	f, err := NewField(name, FLOAT64, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldFloat64{*f}, nil
}

var FLOAT64 DocumentedTyp = TypFLOAT64{}

type TypFLOAT64 struct {
}

func (t TypFLOAT64) String() string {
	return t.TypName()
}

func (t TypFLOAT64) Documentation() string {
	return "Represents a float64.\nThe values are encoded using eight bytes in network byte order (big-endian)."
}

func (t TypFLOAT64) TypName() string {
	return "FLOAT64"
}

func (t TypFLOAT64) Validate(i interface{}) (interface{}, SchemaError) {
	b, ok := i.(float64)
	if ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a float64.", i)
}

func (t TypFLOAT64) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(v.(float64)))
	_, err = buf.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t TypFLOAT64) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != t.SizeOf(bs) {
		return nil, NewSchemaErrorf("%v is not a float64.", bs)
	}
	bits := binary.BigEndian.Uint64(bs)
	return math.Float64frombits(bits), nil
}

func (t TypFLOAT64) SizeOf(i interface{}) int {
	return 8
}

func (t TypFLOAT64) IsNullable() bool {
	return false
}

func (t TypFLOAT64) ArrayElementType() Typ {
	return nil
}

func (t TypFLOAT64) IsArray() bool {
	return false
}
