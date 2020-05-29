package headers

import (
	"fmt"
	"go.uber.org/atomic"
	"papillon/errorx"
)

type RecordHeaders struct {
	headers    []Header
	isReadOnly *atomic.Bool
}

func NewEmptyRecordHeaders() *RecordHeaders {
	return &RecordHeaders{
		headers:    nil,
		isReadOnly: atomic.NewBool(false),
	}
}

func NewRecordHeaders(headers []Header) *RecordHeaders {
	return &RecordHeaders{
		headers:    headers,
		isReadOnly: atomic.NewBool(false),
	}
}

func (h *RecordHeaders) canWrite() errorx.IllegalStateError {
	if h.isReadOnly.Load() {
		return errorx.NewIllegalStateError("RecordHeaders has been closed.")
	}
	return nil
}

func (h *RecordHeaders) AddHeader(header Header) (Headers, errorx.IllegalStateError) {
	if header == nil {
		return h, errorx.NewIllegalStateError("Header cannot be nil.")
	}
	if err := h.canWrite(); err != nil {
		return h, err
	}
	h.headers = append(h.headers, header)
	return h, nil
}

func (h *RecordHeaders) AddKeyValue(key string, value []byte) (Headers, errorx.IllegalStateError) {
	return h.AddHeader(NewRecordHeader(key, value))
}

func (h *RecordHeaders) checkKey(key string) errorx.IllegalStateError {
	if key == "" {
		return errorx.NewIllegalStateError("key cannot be empty.")
	}
	return nil
}

func (h *RecordHeaders) Remove(key string) (Headers, errorx.IllegalStateError) {
	if err := h.canWrite(); err != nil {
		return h, err
	}
	if err := h.checkKey(key); err != nil {
		return h, err
	}
	var n []Header
	for _, hd := range h.headers {
		if hd.Key() == key {
			continue
		}
		n = append(n, hd)
	}
	h.headers = n
	return h, nil
}

func (h *RecordHeaders) LastHeader(key string) Header {
	if len(key) < 1 {
		return nil
	}
	for i := len(h.headers) - 1; i >= 0; i-- {
		header := h.headers[i]
		if header.Key() == key {
			return header
		}
	}
	return nil
}

func (h *RecordHeaders) Headers(key string) []Header {
	if len(key) < 1 {
		return nil
	}
	var n []Header
	for _, hd := range h.headers {
		if hd.Key() != key {
			continue
		}
		n = append(n, hd)
	}
	return n
}

func (h *RecordHeaders) Slice() []Header {
	return h.headers
}

func (h *RecordHeaders) String() string {
	return fmt.Sprintf(
		"RecordHeaders(headers = %v, isReadOnly = %v)",
		h.headers,
		h.isReadOnly.Load(),
	)
}

func (h *RecordHeaders) SetReadOnly() {
	h.isReadOnly.Store(true)
}
