package types

import (
	"bytes"
	"github.com/google/uuid"
)

type FieldUUID struct {
	Field
}

func NewFieldUUID(name string, docString string) (*FieldUUID, SchemaError) {
	f, err := NewField(name, UUID, docString, false, nil)
	if err != nil {
		return nil, err
	}
	return &FieldUUID{*f}, nil
}

func NewFieldUUIDWithDefault(name string, docString string, defaultValue uuid.UUID) (*FieldUUID, SchemaError) {
	f, err := NewField(name, UUID, docString, true, defaultValue)
	if err != nil {
		return nil, err
	}
	return &FieldUUID{*f}, nil
}

var UUID DocumentedTyp = TypUUID{}

type TypUUID struct {
}

func (t TypUUID) String() string {
	return t.TypName()
}

func (t TypUUID) Documentation() string {
	return "Represents a github.com/google/uuid.(RFC4122)\nThe values are encoded using sixteen bytes in network byte order (big-endian)."
}

func (t TypUUID) TypName() string {
	return "UUID"
}

func (t TypUUID) Validate(i interface{}) (interface{}, SchemaError) {
	if b, ok := i.([]uint8); ok {
		if len(b) == 16 {
			return b, nil
		}
	}
	if b, ok := i.(uuid.UUID); ok {
		return b, nil
	}
	return nil, NewSchemaErrorf("%v is not a uuid.UUID.", i)
}

func (t TypUUID) Write(buf *bytes.Buffer, i interface{}) SchemaError {
	v, err := t.Validate(i)
	if err != nil {
		return err
	}
	val, _ := v.(uuid.UUID).MarshalBinary()
	_, err = buf.Write(val)
	if err != nil {
		return err
	}
	return nil
}

func (t TypUUID) Read(buf *bytes.Buffer) (interface{}, SchemaError) {
	bs := buf.Bytes()
	if len(bs) != 16 {
		return nil, NewSchemaErrorf("%v is not a uuid.UUID.", bs)
	}
	if _, err := t.Validate(bs); err != nil {
		return nil, err
	}
	uid, err := uuid.FromBytes(bs)
	if err != nil {
		return nil, NewSchemaError(err.Error())
	}
	return uid, nil
}

func (t TypUUID) SizeOf(i interface{}) int {
	return 16
}

func (t TypUUID) IsNullable() bool {
	return false
}

func (t TypUUID) ArrayElementType() Typ {
	return nil
}

func (t TypUUID) IsArray() bool {
	return false
}
