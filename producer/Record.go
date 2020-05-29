package producer

import (
	"fmt"
	"papillon/common/errorx"
	"papillon/common/headers"
	"papillon/common/partition"
	"reflect"
	"time"
)

type Record struct {
	topic     string
	partition partition.Partition
	headers   headers.Headers
	key       []byte
	value     []byte
	time      time.Time
}

func NewProducerRecord(
	topic string,
	partition partition.Partition,
	times time.Time,
	key []byte,
	value []byte,
	hs []headers.Header,
) (*Record, errorx.IllegalArgumentError) {
	if len(topic) < 1 {
		return nil, errorx.NewIllegalArgumentError("Topic cannot be empty.")
	}
	if times.IsZero() {
		return nil, errorx.NewIllegalArgumentErrorf("Invalid times: %v. Timestamp should always be non-zero.", times)
	}
	if partition != nil && partition.Val() < 0 {
		return nil, errorx.NewIllegalArgumentErrorf("Invalid partition: %v. Partition should always be positive int value.", partition)
	}
	return &Record{
		topic:     topic,
		partition: partition,
		headers:   headers.NewRecordHeaders(hs),
		key:       key,
		value:     value,
		time:      times,
	}, nil
}

func NewProducerRecord1(
	topic string,
	partition partition.Partition,
	times time.Time,
	key []byte,
	value []byte,
) (*Record, errorx.IllegalArgumentError) {
	return NewProducerRecord(topic, partition, times, key, value, nil)
}

func NewProducerRecord2(
	topic string,
	partition partition.Partition,
	key []byte,
	value []byte,
) (*Record, errorx.IllegalArgumentError) {
	return NewProducerRecord(topic, partition, time.Time{}, key, value, nil)
}

func NewProducerRecord3(
	topic string,
	key []byte,
	value []byte,
) (*Record, errorx.IllegalArgumentError) {
	return NewProducerRecord(topic, nil, time.Time{}, key, value, nil)
}

func NewProducerRecord4(
	topic string,
	value []byte,
) (*Record, errorx.IllegalArgumentError) {
	return NewProducerRecord(topic, nil, time.Time{}, nil, value, nil)
}

func (r *Record) String() string {
	return fmt.Sprintf(
		`ProducerRecord(topic="%s", partition="%d", headers="%s", key="%v", value="%v", time="%v")`,
		r.topic, r.partition.Val(), r.headers.String(), r.key, r.value, r.time,
	)
}

func (r *Record) Equals(other interface{}) bool {
	return reflect.DeepEqual(r, other)
}

func (r *Record) Topic() string {
	return r.topic
}

func (r *Record) Partition() partition.Partition {
	return r.partition
}

func (r *Record) Headers() headers.Headers {
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
