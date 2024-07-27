package game

import (
	"encoding/json"
	"sniffsniff/utils"
)

type ObjectEffectInteger struct {
	ObjectEffect ObjectEffect
	Value        uint32
}

func (objectEffectInteger *ObjectEffectInteger) Deserialize(buffer *utils.Buffer) {
	objectEffectInteger.ObjectEffect.Deserialize(buffer)
	value, err := buffer.ReadVarUhInt()
	if err != nil {
		panic(err)
	}
	objectEffectInteger.Value = value
}

func (objectEffectInteger *ObjectEffectInteger) String() string {
	stringified, err := json.Marshal(objectEffectInteger)
	if err != nil {
		return "ObjectEffectInteger{" + err.Error() + "}"
	}
	return "ObjectEffectInteger=" + string(stringified) + ""
}
