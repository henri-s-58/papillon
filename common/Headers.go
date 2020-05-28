package common

type Headers interface {
	AddHeader(header Header) (Headers, IllegalStateError)
	AddKeyValue(key string, value []byte) (Headers, IllegalStateError)
	Remove(key string) (Headers, IllegalStateError)
	LastHeader(key string) Header
	Headers(key string) []Header
	ToArray() []Header
	String() string
}
