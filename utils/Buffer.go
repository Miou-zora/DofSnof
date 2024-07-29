package utils

import (
	"errors"
)

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

func (buffer *Buffer) ReadUShort() (uint16, error) {
	if buffer.Pos+2 > len(buffer.Data) {
		return 0, errors.New("end of buffer reached")
	}
	value := uint16(buffer.Data[buffer.Pos])<<8 | uint16(buffer.Data[buffer.Pos+1])
	buffer.Pos += 2
	return value, nil
}

func (buffer *Buffer) ReadInt() (int, error) {
	if buffer.Pos+4 > len(buffer.Data) {
		return 0, errors.New("end of buffer reached")
	}
	value := int(buffer.Data[buffer.Pos])<<24 | int(buffer.Data[buffer.Pos+1])<<16 | int(buffer.Data[buffer.Pos+2])<<8 | int(buffer.Data[buffer.Pos+3])
	buffer.Pos += 4
	return value, nil
}

func (buffer *Buffer) ReadUInt() (uint32, error) {
	if buffer.Pos+4 > len(buffer.Data) {
		return 0, errors.New("end of buffer reached")
	}
	value := uint32(buffer.Data[buffer.Pos])<<24 | uint32(buffer.Data[buffer.Pos+1])<<16 | uint32(buffer.Data[buffer.Pos+2])<<8 | uint32(buffer.Data[buffer.Pos+3])
	buffer.Pos += 4
	return value, nil
}

func (buffer *Buffer) ReadULong() (uint64, error) {
	if buffer.Pos+8 > len(buffer.Data) {
		return 0.0, errors.New("end of buffer reached")
	}
	value := uint64(buffer.Data[buffer.Pos])<<56 | uint64(buffer.Data[buffer.Pos+1])<<48 | uint64(buffer.Data[buffer.Pos+2])<<40 | uint64(buffer.Data[buffer.Pos+3])<<32 | uint64(buffer.Data[buffer.Pos+4])<<24 | uint64(buffer.Data[buffer.Pos+5])<<16 | uint64(buffer.Data[buffer.Pos+6])<<8 | uint64(buffer.Data[buffer.Pos+7])
	buffer.Pos += 8
	return value, nil
}

func (buffer *Buffer) ReadUTF() (string, error) {
	length, err := buffer.ReadUShort()
	if err != nil {
		return "", err
	}
	if buffer.Pos+int(length) > len(buffer.Data) {
		return "", errors.New("end of buffer reached")
	}
	value := string(buffer.Data[buffer.Pos : buffer.Pos+int(length)])
	buffer.Pos += int(length)
	return value, nil
}

func (buffer *Buffer) ReadVarUhInt() (uint32, error) {
	value := uint32(0)
	offset := uint32(0)
	for offset < 32 {
		byteValue, err := buffer.ReadByte()
		if err != nil {
			return 0, err
		}
		value = value | uint32(byteValue&127)<<offset
		if byteValue&128 == 0 {
			return value, nil
		}
		offset += 7
	}
	return 0, errors.New("too much data")
}

func (buffer *Buffer) ReadVarUhLong() (uint64, error) {
	value := uint64(0)
	offset := uint32(0)
	for offset < 64 {
		byteValue, err := buffer.ReadByte()
		if err != nil {
			return 0, err
		}
		value = value | uint64(byteValue&127)<<offset
		if byteValue&128 == 0 {
			return value, nil
		}
		offset += 7
	}
	return 0, errors.New("too much data")
}
