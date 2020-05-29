package clients

import (
	"context"
	"papillon/clients/requests"
	"papillon/common"
	"papillon/errorx"
	"time"
)

type KafkaClient interface {
	IsReady(ctx context.Context, node *common.Node, now time.Time) bool

	Ready(ctx context.Context, node *common.Node, now time.Time) bool

	ConnectionDelayMs(ctx context.Context, node *common.Node, now time.Time) int64

	PollDelayMs(ctx context.Context, node *common.Node, now time.Time) int64

	ConnectionFailed(ctx context.Context, node *common.Node) bool

	AuthenticationError(node *common.Node) errorx.AuthenticationError

	Send(ctx context.Context, req *ClientRequest, now time.Time) // Queue up

	Poll(ctx context.Context, timeout time.Duration, now time.Time) ([]*ClientRequest, errorx.IllegalStateError)

	Disconnect(ctx context.Context, nodeId int)

	Close(ctx context.Context, nodeId int)

	LeastLoadedNode(ctx context.Context, now time.Time) *common.Node

	InFlightRequestCount() int

	HasInFlightRequests() bool

	InFlightRequestCountByNode(nodeId int) int

	HasInFlightRequestsByNode(nodeId int) bool

	HasReadyNodes(now time.Time) bool

	Wakeup(ctx context.Context)

	InitiateClose(ctx context.Context)

	Active() bool

	NewClientRequest(
		ctx context.Context,
		nodeId int,
		requestBuilder requests.RequestBuilder,
		createdTime time.Time,
		expectResponse bool,
	) *ClientRequest

	NewClientRequestWithCallback(
		ctx context.Context,
		nodeId int,
		requestBuilder requests.RequestBuilder,
		createdTime time.Time,
		expectResponse bool,
		requestTimeoutMs int64,
		callback RequestCompletionHandler,
	) *ClientRequest
}
