package game

import "sniffsniff/utils"

type IObjectEffect interface {
	Deserialize(buffer *utils.Buffer)
}

type ObjectEffect struct {
	ActionId uint16
}

func (objectEffect *ObjectEffect) Deserialize(buffer *utils.Buffer) {
	actionId, err := buffer.ReadUShort()
	if err != nil {
		panic(err)
	}
	objectEffect.ActionId = actionId
}
