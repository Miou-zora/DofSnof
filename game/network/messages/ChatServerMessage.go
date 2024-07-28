package game

import (
	"sniffsniff/utils"
)

type ChatServerMessage struct {
	BaseChatServerMessage ChatAbstractServerMessage
	SenderId              uint64
	SenderName            string
	Prefix                string
	SenderAccountId       uint32
}

func (message ChatServerMessage) GetId() int {
	return 881
}

func (message *ChatServerMessage) Deserialize(buffer *utils.Buffer) error {
	message.BaseChatServerMessage.Deserialize(buffer)
	senderId, err := buffer.ReadULong()
	if err != nil {
		return err
	}
	message.SenderId = senderId
	senderName, err := buffer.ReadUTF()
	if err != nil {
		return err
	}
	message.SenderName = senderName
	prefix, err := buffer.ReadUTF()
	if err != nil {
		return err
	}
	message.Prefix = prefix
	senderAccountId, err := buffer.ReadUInt()
	if err != nil {
		return err
	}
	message.SenderAccountId = senderAccountId
	return nil
}

func (message ChatServerMessage) String() string {
	return "ChatServerMessage{ChatAbstractServerMessage: " + message.BaseChatServerMessage.String() + ", SenderId: " + utils.ULongToString(message.SenderId) + ", SenderName: " + message.SenderName + ", Prefix: " + message.Prefix + ", SenderAccountId: " + utils.UIntToString(message.SenderAccountId) + "}"
}

func CreateChatServerMessage() FinalMessage {
	return &ChatServerMessage{
		BaseChatServerMessage: ChatAbstractServerMessage{
			ChannelId:   0,
			Content:     "",
			Timestamp:   0,
			Fingerprint: ""},
		SenderId:        0,
		SenderName:      "",
		Prefix:          "",
		SenderAccountId: 0,
	}
}
