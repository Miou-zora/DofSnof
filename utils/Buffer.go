package utils

import "errors"

type Buffer struct {
	Data []byte
	Pos  int
}

func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.Pos >= len(buffer.Data) {
		return 0, errors.New("end of buffer reached")
	}
	value := buffer.Data[buffer.Pos]
	buffer.Pos++
	return value, nil
}