package game

type Header struct {
	Id      int32
	LenType int8
	DataLen int64
}

func HeaderFromByte(data []byte) Header {
	firstPart := int32(data[0])<<8 | int32(data[1])
	id := firstPart >> 2
	lenType := int8(firstPart & 3)
	dataLen := int64(0)
	for i := 0; i < int(lenType); i++ {
		dataLen = dataLen<<8 | int64(data[2+i])
	}
	return Header{
		Id:      id,
		LenType: lenType,
		DataLen: dataLen,
	}
}

func (header Header) IsValid() bool {
	_, ok := ID_TO_MESSAGE_NAMES[int(header.Id)]
	return ok
}
