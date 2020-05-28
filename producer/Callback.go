package producer

type Callback interface {
	OnCompletion(metadata RecordMetadata, err error)
}
