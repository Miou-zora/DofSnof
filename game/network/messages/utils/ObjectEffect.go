package game

import "sniffsniff/utils"

type IObjectEffect interface {
	utils.Deserializer
}

type ObjectEffect struct {
	ActionId uint16
}

func (objectEffect *ObjectEffect) Deserialize(buffer *utils.Buffer) error {
	actionId, err := buffer.ReadUShort()
	if err != nil {
		return err
	}
	objectEffect.ActionId = actionId
	return nil
}
