package producer

import (
	"papillon/common"
	"papillon/common/partition"
	"strconv"
	"time"
)

const (
	UNKNOWN_PARTITION int = -1
)

type RecordMetadata struct {
	offset              int64
	timestamp           time.Time
	serializedKeySize   int
	serializedValueSize int
	topicPartition      common.TopicPartition
}

func NewRecordMetadata(
	topicPartition common.TopicPartition,
	baseOffset int64,
	relativeOffset int64,
	timestamp time.Time,
	checksum int64,
	serializedKeySize int,
	serializedValueSize int,
) *RecordMetadata {
	r := &RecordMetadata{}
	r.timestamp = timestamp
	r.serializedKeySize = serializedKeySize
	r.serializedValueSize = serializedValueSize
	r.topicPartition = topicPartition
	if baseOffset == -1 {
		r.offset = baseOffset
	} else {
		r.offset = baseOffset + relativeOffset
	}
	return r
}

func (r *RecordMetadata) HasOffset() bool {
	return r.offset != INVALID_OFFSET
}

func (r *RecordMetadata) Offset() int64 {
	return r.offset
}

func (r *RecordMetadata) HasTimestamp() bool {
	return !r.timestamp.IsZero()
}

func (r *RecordMetadata) Timestamp() time.Time {
	return r.timestamp
}

// Uncompressed key in bytes. If key is nil, then size is -1
func (r *RecordMetadata) SerializedKeySize() int {
	return r.serializedKeySize
}

// Uncompressed value in bytes. If value is nil, then size is -1
func (r *RecordMetadata) SerializedValueSize() int {
	return r.serializedValueSize
}

func (r *RecordMetadata) Topic() string {
	return r.topicPartition.Topic()
}

func (r *RecordMetadata) Partition() partition.Partition {
	return r.topicPartition.Partition()
}

func (r *RecordMetadata) String() string {
	return r.topicPartition.String() + "@" + strconv.FormatInt(r.offset, 10)
}
