package types

import "bytes"

type FieldBoolean struct {
	Field
}

func NewFieldBoolean(name string, docString string) (*FieldBoolean, SchemaError) {
	f, err := NewField(name, BOOLEAN, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldBoolean{*f}, nil
}

func NewFieldBooleanWithDefault(name string, docString string, defaultValue bool) (*FieldBoolean, SchemaError) {
	f, err := NewField(name, BOOLEAN, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldBoolean{*f}, nil
}

var BOOLEAN DocumentedTyp = TypBOOLEAN{}

type TypBOOLEAN struct {
}

func (t TypBOOLEAN) String() string {
	return t.TypName()
}

func (t TypBOOLEAN) Documentation() string {
	return "Represents a boolean value in a byte.\nValues 0 and 1 are used to represent false and true respectively.\nWhen reading a boolean value, any non-zero value is considered true."
}

func (t TypBOOLEAN) TypName() string {
	return "BOOLEAN"
}

func (t TypBOOLEAN) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.(bool); ok {
		return b, nil
	}
	if c, ok := i.(uint8); ok {
		if c != 0 && c != 1 {
			return nil, NewSchemaErrorf("%v is not a bool.", i)
		}
		return c, nil
	}
	return nil, NewSchemaErrorf("%v is not a bool.", i)
}

func (t TypBOOLEAN) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	val, _ := v.(bool)
	if val {
		err = buf.WriteByte(uint8(1))
	} else {
		err = buf.WriteByte(uint8(0))
	}
	if err != nil {
		return err
	}
	return nil
}

func (t TypBOOLEAN) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != 1 {
		return nil, NewSchemaErrorf("%v is not a bool.", bs)
	}
	if _, err := t.Validate(bs[0]); err != nil {
		return nil, err
	}
	return bs[0] == 1, nil
}

func (t TypBOOLEAN) SizeOf(i interface{}) int {
	return 1
}

func (t TypBOOLEAN) IsNullable() bool {
	return false
}

func (t TypBOOLEAN) ArrayElementType() Typ {
	return nil
}

func (t TypBOOLEAN) IsArray() bool {
	return false
}
