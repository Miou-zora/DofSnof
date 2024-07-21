package utils

type Deserializer interface {
	Deserialize(buffer *Buffer)
}