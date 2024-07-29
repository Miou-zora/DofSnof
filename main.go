package main

import (
	"fmt"
	"log"
	game_network "sniffsniff/game/network"
	game_message "sniffsniff/game/network/messages"
	"sniffsniff/network"
	"sniffsniff/utils"
)

const (
	MaxBufferSize = 4096
	DefaultFilter = "tcp src port 5555"
)

func main() {
	device, err := network.AskForDevice()

	if err != nil {
		log.Fatalf("Error getting device: %v", err)
	}

	buffer := make([]byte, 0)
	receiver := network.PacketSniffer{
		Buffer: make(chan []byte, MaxBufferSize),
		Filter: DefaultFilter,
		Device: device,
	}
	messages := make([]game_message.FinalMessage, 0)
	receiver.Run()
	for {
		PullMessages(receiver, buffer, messages)
	}
}

func PullMessages(receiver network.PacketSniffer, buffer []byte, messages []game_message.FinalMessage) {
	select {
	case raw_data := <-receiver.Buffer:
		if len(raw_data) == 0 {
			return
		}
		buffer = append(buffer, raw_data...)
		UnstackMessages(buffer, messages)
	default:
		return
	}
}

func UnstackMessages(buffer []byte, messages []game_message.FinalMessage) {
	for len(buffer) > 2 {
		header := game_network.HeaderFromByte(buffer)
		if header.IsValid() {
			fmt.Println("Message: ", game_network.ID_TO_MESSAGE_NAMES[int(header.Id)])
		} else {
			fmt.Println("Invalid message: ", header.Id)
			buffer = buffer[:0]
			continue
		}
		size := header.GetSize()
		if size > len(buffer) {
			fmt.Print("Packet is not complete, waiting for more data...")
			continue
		}
		data := utils.Buffer{Data: buffer[(2 + header.LenType):size], Pos: 0}
		buffer = buffer[size:]
		message, err := GetMessageFromData(header, data)
		if err != nil {
			continue
		} else {
			messages = append(messages, message)
		}
	}
}

func GetMessageFromData(header game_network.Header, data utils.Buffer) (game_message.FinalMessage, error) {
	if game_network.ID_TO_MESSAGE[header.Id] != nil {
		message := game_network.ID_TO_MESSAGE[header.Id]()
		err := message.Deserialize(&data)
		if err != nil {
			return nil, err
		} else {
			fmt.Println("Message: ", message)
			return message, nil
		}
	}
	return nil, fmt.Errorf("message with id %d not found", header.Id)
}

func ResetBuffer(buffer *[]byte) []byte {
	return (*buffer)[:0]
}
