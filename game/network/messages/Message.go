package game

import "sniffsniff/utils"

type Id int

type IMessage interface {
	GetId() Id
}

type FinalMessage interface {
	IMessage
	utils.Deserializer
	utils.Stringify
}

type INewFinalMessage func() FinalMessage