package types

type DocumentedTyp interface {
	Typ
	TypName() string
	Documentation() string
}
