package common

import (
	"papillon/support"
	"reflect"
	"strconv"
)

type TopicPartition struct {
	hash      int64
	partition int
	topic     string
}

func NewTopicPartition(topic string, partition int) TopicPartition {
	return TopicPartition{
		hash:      0,
		partition: partition,
		topic:     topic,
	}
}

func (t TopicPartition) Partition() int {
	return t.partition
}

func (t TopicPartition) Topic() string {
	return t.topic
}

func (t TopicPartition) HashCode() int64 {
	if t.hash != 0 {
		return t.hash
	}
	var (
		prime  int64 = 31
		result int64 = 1
	)
	result = prime*result + int64(t.partition)
	result = prime*result + support.StrToHashCode(t.topic)
	t.hash = result
	return result
}

func (t TopicPartition) Equals(other interface{}) bool {
	return reflect.DeepEqual(t, other)
}

func (t TopicPartition) String() string {
	return t.topic + "-" + strconv.Itoa(t.partition)
}
