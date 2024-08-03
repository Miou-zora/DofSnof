package game

import (
	"fmt"

	game_message "sniffsniff/game/network/messages"
	"sniffsniff/network"
	"sniffsniff/utils"
)

type MessageProcessor struct {
	Messages *[]game_message.FinalMessage
	buffer   []byte
}

func (processor *MessageProcessor) Run(receiver *network.PacketSniffer, message_callback *map[game_message.Id]func(*game_message.FinalMessage)) {
	if processor.Messages == nil {
		processor.Messages = &[]game_message.FinalMessage{}
	}
	if processor.buffer == nil {
		processor.buffer = make([]byte, 0)
	}
	for {
		processor.pullMessages(receiver)
		processor.handleMessages(message_callback)
		processor.updateMessages()
	}
}

func (processor *MessageProcessor) pullMessages(receiver *network.PacketSniffer) {
	select {
	case raw_data := <-receiver.Buffer:
		if len(raw_data) == 0 {
			return
		}
		processor.buffer = append(processor.buffer, raw_data...)
		processor.unpackMessage()
	default:
		return
	}
}

func (processor *MessageProcessor) unpackMessage() {
	for len(processor.buffer) > 2 {
		header := HeaderFromByte(processor.buffer)
		if !header.Valid() {
			fmt.Println("Invalid message: ", header.Id)
			processor.buffer = processor.buffer[:0]
			continue
		}
		size := header.TotalSize()
		if size > len(processor.buffer) {
			fmt.Println("Packet is not complete, waiting for more data...")
			processor.buffer = processor.buffer[:0]
			continue
		}
		data := utils.Buffer{Data: processor.buffer[(2 + header.LenType):size], Pos: 0}
		processor.buffer = processor.buffer[size:]
		message, err := GetMessageFromData(header, data)
		if err == nil {
			*processor.Messages = append(*processor.Messages, message)
		}
	}
}

func (processor *MessageProcessor) handleMessages(message_id_to_callback *map[game_message.Id]func(*game_message.FinalMessage)) {
	for _, message := range *processor.Messages {
		if callback, ok := (*message_id_to_callback)[message.GetId()]; ok {
			callback(&message)
		}
	}
}

func (processor *MessageProcessor) updateMessages() {
	(*processor.Messages) = (*processor.Messages)[:0]
}
