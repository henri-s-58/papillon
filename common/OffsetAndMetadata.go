package common

import (
	"errors"
	"fmt"
	"gopkg.in/guregu/null.v4"
	"reflect"
)

type OffsetAndMetadata struct {
	offset      int64
	metadata    string
	leaderEpoch null.Int
}

func NewOffsetAndMetadata(
	offset int64,
	leaderEpoch null.Int,
	metadata string,
) (*OffsetAndMetadata, error) {
	if offset < 0 {
		return nil, errors.New("invalid negative offset")
	}
	om := &OffsetAndMetadata{}
	om.offset = offset
	om.leaderEpoch = leaderEpoch
	om.metadata = metadata
	return om, nil
}

func (o *OffsetAndMetadata) Offset() int64 {
	return o.offset
}

func (o *OffsetAndMetadata) Metadata() string {
	return o.metadata
}

func (o *OffsetAndMetadata) LeaderEpoch() int64 {
	if o == nil || !o.leaderEpoch.Valid || o.leaderEpoch.Int64 < 0 {
		return 0
	}
	return o.leaderEpoch.Int64
}

func (o *OffsetAndMetadata) Equals(other interface{}) bool {
	return reflect.DeepEqual(o, other)
}

func (o *OffsetAndMetadata) String() string {
	return fmt.Sprintf(
		"OffsetAndMetadata{offset=%d, leaderEpoch=%d, metadata='%s'}",
		o.offset, o.LeaderEpoch(), o.metadata,
	)
}
