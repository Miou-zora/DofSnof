package game

import (
	"sniffsniff/utils"
)

type BasicPongMessage struct {
	Quiet bool
}

func (message BasicPongMessage) GetId() int {
	return 4877
}

func (message *BasicPongMessage) Deserialize(buffer *utils.Buffer) error {
	byteValue, err := buffer.ReadByte()
	if err != nil {
		return err
	}
	message.Quiet = byteValue != 0
	return nil
}

func (message BasicPongMessage) String() string {
	return "BasicPongMessage{Quiet: " + utils.BoolToString(message.Quiet) + "}"
}

func CreateBasicPongMessage() FinalMessage {
	return &BasicPongMessage{Quiet: false}
}
