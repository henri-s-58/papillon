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

var INT8 DocumentedTyp = typINT8{}

type typINT8 struct {
}

func (t typINT8) String() string {
	return t.TypName()
}

func (t typINT8) Documentation() string {
	return "Represents an integer between -2<sup>7</sup> and 2<sup>7</sup>-1 inclusive."
}

func (t typINT8) TypName() string {
	return "INT8"
}

func (t typINT8) Validate(i interface{}) (interface{}, SchemaError) {
	b, ok := i.(int8)
	if ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a int8.", i)
}

func (t typINT8) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	err = buf.WriteByte(v.(uint8))
	if err != nil {
		return SchemaError(err)
	}
	return nil
}

func (t typINT8) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != 1 {
		return nil, NewSchemaErrorf("%v is not a int8.", bs)
	}
	if _, err := t.Validate(bs[0]); err != nil {
		return nil, err
	}
	return bs[0], nil
}

func (t typINT8) SizeOf(i interface{}) int {
	return 1
}

func (t typINT8) IsNullable() bool {
	return false
}

func (t typINT8) ArrayElementType() Typ {
	return nil
}

func (t typINT8) IsArray() bool {
	return false
}
