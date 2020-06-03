package types

import "bytes"

type FieldInt8 struct {
	Field
}

func NewFieldInt8(name string, docString string) (*FieldInt8, SchemaError) {
	f, err := NewField(name, INT8, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldInt8{*f}, nil
}

func NewFieldInt8WithDefault(name string, docString string, defaultValue int8) (*FieldInt8, SchemaError) {
	f, err := NewField(name, INT8, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldInt8{*f}, nil
}

var INT8 DocumentedTyp = TypINT8{}

type TypINT8 struct {
}

func (t TypINT8) String() string {
	return t.TypName()
}

func (t TypINT8) Documentation() string {
	return "Represents an integer between -2<sup>7</sup> and 2<sup>7</sup>-1 inclusive."
}

func (t TypINT8) TypName() string {
	return "INT8"
}

func (t TypINT8) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.(int8); ok {
		return b, nil
	}
	if c, ok := i.(uint8); ok {
		return c, nil
	}
	return nil, NewSchemaErrorf("%v is not a int8.", i)
}

func (t TypINT8) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	val, _ := v.(int8)
	err = buf.WriteByte(uint8(val))
	if err != nil {
		return err
	}
	return nil
}

func (t TypINT8) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != 1 {
		return nil, NewSchemaErrorf("%v is not a int8.", bs)
	}
	if _, err := t.Validate(bs[0]); err != nil {
		return nil, err
	}
	return int8(bs[0]), nil
}

func (t TypINT8) SizeOf(i interface{}) int {
	return 1
}

func (t TypINT8) IsNullable() bool {
	return false
}

func (t TypINT8) ArrayElementType() Typ {
	return nil
}

func (t TypINT8) IsArray() bool {
	return false
}
