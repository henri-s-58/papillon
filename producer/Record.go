package producer

import (
	"fmt"
	"papillon/common"
	"reflect"
	"time"
)

type Record struct {
	topic     string
	partition int
	headers   common.Headers
	key       []byte
	value     []byte
	time      time.Time
}

func (r *Record) String() string {
	return fmt.Sprintf(
		`ProducerRecord(topic="%s", partition="%d", headers="%s", key="%v", value="%v", time="%v")`,
		r.topic, r.partition, r.headers.String(), r.key, r.value, r.time,
	)
}

func (r *Record) Equals(other interface{}) bool {
	return reflect.DeepEqual(r, other)
}

func (r *Record) Topic() string {
	return r.topic
}

func (r *Record) Partition() int {
	return r.partition
}

func (r *Record) Headers() common.Headers {
	return r.headers
}

func (r *Record) Key() []byte {
	return r.key
}

func (r *Record) Value() []byte {
	return r.value
}

func (r *Record) Time() time.Time {
	return r.time
}
