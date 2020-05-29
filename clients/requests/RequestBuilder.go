package requests

import (
	"papillon/common/protocols"
)

type RequestBuilder interface {
	APIKey() *protocols.APIKey
	OldestAllowedVersion() int16
	LatestAllowedVersion() int16
	build() Request
	buildFor(version int16) Request
}

type RequestBuilderI struct {
	apiKey               *protocols.APIKey
	oldestAllowedVersion int16
	latestAllowedVersion int16
}
