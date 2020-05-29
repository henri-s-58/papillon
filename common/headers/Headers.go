package headers

import (
	"papillon/common/errorx"
)

type Headers interface {
	AddHeader(header Header) (Headers, errorx.IllegalStateError)
	AddKeyValue(key string, value []byte) (Headers, errorx.IllegalStateError)
	Remove(key string) (Headers, errorx.IllegalStateError)
	LastHeader(key string) Header
	Headers(key string) []Header
	Slice() []Header
	String() string
}
