package game

import "sniffsniff/utils"

type Deserializer interface {
	Deserialize(buffer utils.Buffer)
}

type IMessage interface {
	GetId() int
}

type Stringify interface {
	String() string
}


type INewFinalMessage func() FinalMessage
type FinalMessage interface {
	IMessage
	Deserializer
	Stringify
}