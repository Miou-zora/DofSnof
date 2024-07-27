package game

import (
	"encoding/json"
	game "sniffsniff/game/network/messages/utils"
	"sniffsniff/utils"
)

type ExchangeTypesItemsExchangerDescriptionForUserMessage struct {
	ObjectGID           uint32
	ObjectType          uint32
	ItemTypeDescription []game.BidExchangerObjectInfo
}

func (message ExchangeTypesItemsExchangerDescriptionForUserMessage) GetId() int {
	return 2738
}

func (message *ExchangeTypesItemsExchangerDescriptionForUserMessage) Deserialize(buffer *utils.Buffer) {
	message._objectGIDFunc(buffer)
	message._objectTypeFunc(buffer)
	numberOfItemTypeDescription, err := buffer.ReadUShort()

	if err != nil {
		panic(err)
	}
	message.ItemTypeDescription = make([]game.BidExchangerObjectInfo, numberOfItemTypeDescription)
	for i := 0; i < int(numberOfItemTypeDescription); i++ {
		message.ItemTypeDescription[i] = game.BidExchangerObjectInfo{}
		message.ItemTypeDescription[i].Deserialize(buffer)
	}
}

func (message ExchangeTypesItemsExchangerDescriptionForUserMessage) String() string {
	stringified, err := json.Marshal(message)
	if err != nil {
		return "ExchangeTypesItemsExchangerDescriptionForUserMessage{" + err.Error() + "}"
	}
	return "ExchangeTypesItemsExchangerDescriptionForUserMessage=" + string(stringified) + ""
}

func CreateExchangeTypesItemsExchangerDescriptionForUserMessage() FinalMessage {
	return &ExchangeTypesItemsExchangerDescriptionForUserMessage{}
}

func (message *ExchangeTypesItemsExchangerDescriptionForUserMessage) _objectGIDFunc(buffer *utils.Buffer) {
	objectGID, err := buffer.ReadVarUhInt()
	if err != nil {
		panic(err)
	}
	message.ObjectGID = objectGID
}

func (message *ExchangeTypesItemsExchangerDescriptionForUserMessage) _objectTypeFunc(buffer *utils.Buffer) {
	objectType, err := buffer.ReadUInt()
	if err != nil {
		panic(err)
	}
	message.ObjectType = objectType
}
