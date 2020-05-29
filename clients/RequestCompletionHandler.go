package clients

type RequestCompletionHandler interface {
	OnComplete(response *ClientResponse)
}
