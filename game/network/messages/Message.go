package game

import "sniffsniff/utils"

type IMessage interface {
	GetId() int
}

type FinalMessage interface {
	IMessage
	utils.Deserializer
	utils.Stringify
}

type INewFinalMessage func() FinalMessage