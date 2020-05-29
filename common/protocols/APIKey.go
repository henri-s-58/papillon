package protocols

import (
	"papillon/common/protocols/types"
	"papillon/errorx"
	"papillon/support"
)

var (
	APIKeys_MIN_API_KEY int16 = 0
	APIKeys_MAX_API_KEY int16 = 0
	APIKeys_ID_TO_TYPE  []*APIKey
)

type APIKey struct {
	id                          int16
	name                        string
	clusterAction               bool
	minRequiredInterBrokerMagic byte
	requiresDelayedAllocation   bool
	requestSchemas              []types.Schema
	responseSchemas             []types.Schema
}

var apiKeys []*APIKey

func APIKeysValues() []*APIKey {
	return apiKeys
}

func init() {
	var maxKey int16 = -1
	for _, key := range APIKeysValues() {
		maxKey = support.Int16Max(maxKey, key.id)
	}
	var itt = make([]*APIKey, maxKey+1)
	for _, key := range APIKeysValues() {
		itt[key.id] = key
	}
	APIKeys_ID_TO_TYPE = itt
	APIKeys_MAX_API_KEY = maxKey
}

func ForId(id int16) (*APIKey, errorx.IllegalArgumentError) {
	if !HasId(id) {
		return nil, errorx.NewIllegalArgumentErrorf("Unexpected ApiKeys id `%d`, it should be between `%d` and `%d` (inclusive)",
			id, APIKeys_MIN_API_KEY, APIKeys_MAX_API_KEY,
		)
	}
	return APIKeys_ID_TO_TYPE[id], nil
}

func HasId(id int16) bool {
	return id >= APIKeys_MIN_API_KEY && id <= APIKeys_MAX_API_KEY
}
