package headers

type Header interface {
	Key() string
	Value() []byte
}
