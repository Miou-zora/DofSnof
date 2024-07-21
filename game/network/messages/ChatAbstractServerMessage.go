package game

import "sniffsniff/utils"

type ChatAbstractServerMessage struct {
	ChannelId   byte
	Content     string
	Timestamp   uint32
	Fingerprint string
}

func (message ChatAbstractServerMessage) GetId() int {
	return 880
}

func (message *ChatAbstractServerMessage) Deserialize(buffer *utils.Buffer) {
	channelId, err := buffer.ReadByte()
	if err != nil {
		panic(err)
	}
	message.ChannelId = uint8(channelId)
	content, err := buffer.ReadUTF()
	if err != nil {
		panic(err)
	}
	message.Content = content
	timestamp, err := buffer.ReadUInt()
	if err != nil {
		panic(err)
	}
	message.Timestamp = timestamp
	fingerprint, err := buffer.ReadUTF()
	if err != nil {
		panic(err)
	}
	message.Fingerprint = fingerprint
}

func (message ChatAbstractServerMessage) String() string {
	return "ChatAbstractServerMessage{ChannelId: " + utils.ByteToString(message.ChannelId) + ", Content: " + message.Content + ", Timestamp: " + utils.UIntToString(message.Timestamp) + ", Fingerprint: " + message.Fingerprint + "}"
}

func CreateChatAbstractServerMessage() FinalMessage {
	return &ChatAbstractServerMessage{ChannelId: 0, Content: "", Timestamp: 0, Fingerprint: ""}
}
