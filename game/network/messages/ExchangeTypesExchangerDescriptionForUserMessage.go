package game

import (
	"encoding/json"
	"sniffsniff/utils"
)

type ExchangeTypesExchangerDescriptionForUserMessage struct {
	ObjectType      uint32
	TypeDescription []uint32
}

func (message ExchangeTypesExchangerDescriptionForUserMessage) GetId() Id {
	return 6572
}

func (message *ExchangeTypesExchangerDescriptionForUserMessage) Deserialize(buffer *utils.Buffer) error {
	err := message._objectTypeFunc(buffer)
	if err != nil {
		return err
	}
	typeDescriptionLength, err := buffer.ReadUShort()
	if err != nil {
		return err
	}
	message.TypeDescription = make([]uint32, typeDescriptionLength)
	for i := 0; i < int(typeDescriptionLength); i++ {
		message.TypeDescription[i], err = buffer.ReadVarUhInt()
		if err != nil {
			return err
		}
	}
	return nil
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

func (message *ExchangeTypesExchangerDescriptionForUserMessage) _objectTypeFunc(buffer *utils.Buffer) error {
	objectType, err := buffer.ReadUInt()
	if err != nil {
		return err
	}
	message.ObjectType = objectType
	return nil
}
