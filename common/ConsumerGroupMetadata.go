package common

import (
	"fmt"
	"gopkg.in/guregu/null.v4"
	"reflect"
)

type ConsumerGroupMetadata struct {
	groupId         string
	generationId    int
	memberId        string
	groupInstanceId null.String
}

func NewConsumerGroupMetadata(
	groupId string,
	generationId int,
	memberId string,
	groupInstanceId null.String,
) *ConsumerGroupMetadata {
	return &ConsumerGroupMetadata{
		groupId:         groupId,
		generationId:    generationId,
		memberId:        memberId,
		groupInstanceId: groupInstanceId,
	}
}

func NewDefaultConsumerGroupMetadata(groupId string) *ConsumerGroupMetadata {
	return NewConsumerGroupMetadata(
		groupId,
		UNKNOWN_GENERATION_ID,
		UNKNOWN_MEMBER_ID,
		null.NewString("", false),
	)
}

func (c *ConsumerGroupMetadata) GroupId() string {
	return c.groupId
}

func (c *ConsumerGroupMetadata) GenerationId() int {
	return c.generationId
}

func (c *ConsumerGroupMetadata) MemberId() string {
	return c.memberId
}

func (c *ConsumerGroupMetadata) GroupInstanceId() null.String {
	return c.groupInstanceId
}

func (c *ConsumerGroupMetadata) String() string {
	return fmt.Sprintf(
		"GroupMetadata(groupId = %s, generationId = %d, memberId = %s, groupInstanceId = %s)",
		c.groupId,
		c.generationId,
		c.memberId,
		c.groupInstanceId.ValueOrZero())
}

func (c *ConsumerGroupMetadata) Equals(other interface{}) bool {
	return reflect.DeepEqual(c, other)
}
