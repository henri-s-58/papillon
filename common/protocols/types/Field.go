package types

type Field struct {
	name            string
	docString       string
	typ             Typ
	hasDefaultValue bool
	defaultValue    interface{}
}

func NewField(
	name string,
	typ Typ,
	docString string,
	hasDefaultValue bool,
	defaultValue interface{},
) (*Field, SchemaError) {
	if hasDefaultValue {
		if _, err := typ.Validate(defaultValue); err != nil {
			return nil, err
		}
	}
	return &Field{
		name:            name,
		docString:       docString,
		typ:             typ,
		hasDefaultValue: hasDefaultValue,
		defaultValue:    defaultValue,
	}, nil
}

func NewField1(name string, typ Typ, docString string) (*Field, SchemaError) {
	return NewField(name, typ, docString, false, nil)
}

func NewField2(name string, typ Typ, docString string, defaultValue interface{}) (*Field, SchemaError) {
	return NewField(name, typ, docString, true, defaultValue)
}

func NewField3(name string, typ Typ) (*Field, SchemaError) {
	return NewField(name, typ, "", false, nil)
}
