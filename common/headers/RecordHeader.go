package headers

import "bytes"

type RecordHeader struct {
	key      string
	value    []byte
	valueBuf *bytes.Buffer
}

func NewRecordHeader(key string, value []byte) *RecordHeader {
	return &RecordHeader{
		key:   key,
		value: value,
	}
}

func NewBufferedRecordHeader(key string, valueBuf *bytes.Buffer) *RecordHeader {
	return &RecordHeader{
		key:      key,
		valueBuf: valueBuf,
	}
}

func (r *RecordHeader) Key() string {
	return r.key
}

func (r *RecordHeader) Value() []byte {
	if r.value == nil && r.valueBuf != nil {
		r.value = r.valueBuf.Bytes()
		r.valueBuf.Reset()
	}
	return r.value
}

func (r *RecordHeader) String() string {
	return "RecordHeader(key = " + r.key + ", value = " + string(r.Value()) + ")"
}
