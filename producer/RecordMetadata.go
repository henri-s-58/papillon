package producer

import (
	"go.uber.org/atomic"
	"papillon/common"
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
	checksum            atomic.Int64
}
