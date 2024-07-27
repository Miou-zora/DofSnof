package game

import (
	"fmt"
	"sniffsniff/utils"
)

type BidExchangerObjectInfo struct {
	ObjectUID  uint32
	ObjectGID  uint32
	ObjectType uint32
	Effects    []IObjectEffect
	Prices     []uint64
}

func (message *BidExchangerObjectInfo) Deserialize(buffer *utils.Buffer) {
	message._objectUIDFunc(buffer)
	message._objectGIDFunc(buffer)
	message._objectTypeFunc(buffer)
	numberOfEffects, err := buffer.ReadUShort()

	if err != nil {
		panic(err)
	}
	message.Effects = make([]IObjectEffect, numberOfEffects)
	for i := 0; i < int(numberOfEffects); i++ {
		message.Effects[i] = message.ReadEffect(buffer)
	}
	numberOfPrices, err := buffer.ReadUShort()
	if err != nil {
		panic(err)
	}
	message.Prices = make([]uint64, numberOfPrices)
	for i := 0; i < int(numberOfPrices); i++ {
		message.Prices[i] = message.ReadPrice(buffer)
	}
}

func (message *BidExchangerObjectInfo) ReadEffect(buffer *utils.Buffer) IObjectEffect {
	id, err := buffer.ReadUShort()
	if err != nil {
		panic(err)
	}
	if id == 3930 {
		effect := ObjectEffectInteger{}
		effect.Deserialize(buffer)
		return &effect
	}
	fmt.Println("Unknown type of effect " + utils.UShortToString(id))
	return &ObjectEffect{}
}

func (message *BidExchangerObjectInfo) ReadPrice(buffer *utils.Buffer) uint64 {
	price, err := buffer.ReadVarUhLong()
	if err != nil {
		panic(err)
	}
	if price > 9_007_199_254_740_992 {
		panic("Forbidden value (" + utils.ULongToString(price) + ") on elements of prices.")
	}
	return price
}

func (message *BidExchangerObjectInfo) _objectUIDFunc(buffer *utils.Buffer) {
	objectUID, err := buffer.ReadVarUhInt()
	if err != nil {
		panic(err)
	}
	message.ObjectUID = objectUID
}

func (message *BidExchangerObjectInfo) _objectGIDFunc(buffer *utils.Buffer) {
	objectGID, err := buffer.ReadVarUhInt()
	if err != nil {
		panic(err)
	}
	message.ObjectGID = objectGID
}

func (message *BidExchangerObjectInfo) _objectTypeFunc(buffer *utils.Buffer) {
	objectType, err := buffer.ReadUInt()
	if err != nil {
		panic(err)
	}
	message.ObjectType = objectType
}
