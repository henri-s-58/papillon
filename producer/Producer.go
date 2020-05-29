package producer

import (
	"context"
	. "papillon/common"
	"time"
)

type Producer interface {
	InitTransactions()
	BeginTransaction() FencedError
	SendOffsetsToTransaction(offsets map[TopicPartition]OffsetAndMetadata, consumerGroupId string) FencedError
	SendOffsetsToTransactionForGroupMetadata(offsets map[TopicPartition]OffsetAndMetadata, groupMetadata ConsumerGroupMetadata) FencedError
	CommitTransaction() FencedError
	AbortTransaction() FencedError
	Flush()
	PartitionsFor(topic string) []PartitionInfo
	Close(timeout time.Duration)
	Send(ctx context.Context, record *Record, callback Callback) <-chan *RecordMetadata
}
