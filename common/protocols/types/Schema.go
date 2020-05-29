package types

type Schema struct {
	fields                            []*BoundField
	fieldsByName                      map[string]*BoundField
	tolerateMissingFieldsWithDefaults bool
}
