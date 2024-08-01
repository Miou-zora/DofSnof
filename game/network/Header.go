package game

type Header struct {
	Id      uint16
	LenType uint8
	DataLen uint32
}

func HeaderFromByte(data []byte) Header {
	firstPart := uint16(data[0])<<8 | uint16(data[1])
	id := firstPart >> 2
	lenType := uint8(firstPart & 0b11)
	dataLen := uint32(0)
	for i := 0; i < int(lenType); i++ {
		dataLen = dataLen<<8 | uint32(data[2+i])
	}
	return Header{
		Id:      id,
		LenType: lenType,
		DataLen: dataLen,
	}
}

func (header Header) Valid() bool {
	_, ok := ID_TO_MESSAGE_NAMES[int(header.Id)]
	return ok
}

func (header Header) TotalSize() int {
	return header.Size() + header.MessageSize()
}

func (header Header) Size() int {
	return 2 + int(header.LenType)
}

func (header Header) MessageSize() int {
	return int(header.DataLen)
}