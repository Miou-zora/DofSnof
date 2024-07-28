package game

import (
	"errors"
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

func (message *BidExchangerObjectInfo) Deserialize(buffer *utils.Buffer) error {
	message._objectUIDFunc(buffer)
	message._objectGIDFunc(buffer)
	message._objectTypeFunc(buffer)
	numberOfEffects, err := buffer.ReadUShort()

	if err != nil {
		return err
	}
	message.Effects = make([]IObjectEffect, numberOfEffects)
	for i := 0; i < int(numberOfEffects); i++ {
		message.Effects[i], err = message.ReadEffect(buffer)
		if err != nil {
			return err
		}
	}
	numberOfPrices, err := buffer.ReadUShort()
	if err != nil {
		return err
	}
	message.Prices = make([]uint64, numberOfPrices)
	for i := 0; i < int(numberOfPrices); i++ {
		message.Prices[i], err = message.ReadPrice(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (message *BidExchangerObjectInfo) ReadEffect(buffer *utils.Buffer) (IObjectEffect, error) {
	id, err := buffer.ReadUShort()
	if err != nil {
		return &ObjectEffect{}, err
	}
	if id == 3930 {
		effect := ObjectEffectInteger{}
		effect.Deserialize(buffer)
		return &effect, nil
	}
	fmt.Println("Unknown type of effect " + utils.UShortToString(id))
	return &ObjectEffect{}, errors.New("Unknown type of effect " + utils.UShortToString(id))
}

func (message *BidExchangerObjectInfo) ReadPrice(buffer *utils.Buffer) (uint64, error) {
	price, err := buffer.ReadVarUhLong()
	if err != nil {
		return 0, err
	}
	if price > 9_007_199_254_740_992 {
		return 0, errors.New("Forbidden value (" + utils.ULongToString(price) + ") on elements of prices.")
	}
	return price, nil
}

func (message *BidExchangerObjectInfo) _objectUIDFunc(buffer *utils.Buffer) error {
	objectUID, err := buffer.ReadVarUhInt()
	if err != nil {
		return err
	}
	message.ObjectUID = objectUID
	return nil
}

func (message *BidExchangerObjectInfo) _objectGIDFunc(buffer *utils.Buffer) error {
	objectGID, err := buffer.ReadVarUhInt()
	if err != nil {
		return err
	}
	message.ObjectGID = objectGID
	return nil
}

func (message *BidExchangerObjectInfo) _objectTypeFunc(buffer *utils.Buffer) error {
	objectType, err := buffer.ReadUInt()
	if err != nil {
		return err
	}
	message.ObjectType = objectType
	return nil
}
