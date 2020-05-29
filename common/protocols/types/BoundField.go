package types

type BoundField struct {
	def    *Field
	index  int
	schema *Schema
}

func NewBoundField(def *Field, schema *Schema, index int) *BoundField {
	return &BoundField{
		def:    def,
		index:  index,
		schema: schema,
	}
}

func (b *BoundField) String() string {
	return b.def.name + ":" + b.def.typ.String()
}
