package game

import (
	"sniffsniff/utils"
	"encoding/json"
)

type ExchangeTypesExchangerDescriptionForUserMessage struct {
	ObjectType      uint32
	TypeDescription []uint32
}

func (message ExchangeTypesExchangerDescriptionForUserMessage) GetId() int {
	return 6572
}

func (message *ExchangeTypesExchangerDescriptionForUserMessage) Deserialize(buffer *utils.Buffer) {
	message._objectTypeFunc(buffer)
	typeDescriptionLength, err := buffer.ReadUShort()
	if err != nil {
		panic(err)
	}
	message.TypeDescription = make([]uint32, typeDescriptionLength)
	for i := 0; i < int(typeDescriptionLength); i++ {
		message.TypeDescription[i], err = buffer.ReadVarUhInt()
		if err != nil {
			panic(err)
		}
	}
}

func (message ExchangeTypesExchangerDescriptionForUserMessage) String() string {
	stringified, err := json.Marshal(message)
	if err != nil {
		return "ExchangeTypesExchangerDescriptionForUserMessage{" + err.Error() + "}"
	}
	return "ExchangeTypesExchangerDescriptionForUserMessage=" + string(stringified) + ""
}

func CreateExchangeTypesExchangerDescriptionForUserMessage() FinalMessage {
	return &ExchangeTypesExchangerDescriptionForUserMessage{}
}

func (message *ExchangeTypesExchangerDescriptionForUserMessage) _objectTypeFunc(buffer *utils.Buffer) {
	objectType, err := buffer.ReadUInt()
	if err != nil {
		panic(err)
	}
	message.ObjectType = objectType
}