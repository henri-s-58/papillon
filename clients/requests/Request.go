package requests

import (
	"papillon/common/protocols"
)

type Request interface {
	API() *protocols.APIKey
	Version() int16
}
