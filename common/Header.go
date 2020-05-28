package common

type Header interface {
	Key() string
	Value() []byte
}
