package main

import (
	"fmt"
	"log"
	"sniffsniff/database"
	game_network "sniffsniff/game/network"
	game_message "sniffsniff/game/network/messages"
	game_resources "sniffsniff/game/resources"
	"sniffsniff/network"
	"sniffsniff/utils"

	"github.com/joho/godotenv"
)

const (
	MaxBufferSize = 4096
	DefaultFilter = "tcp src port 5555"
)

func main() {
	SetupDotEnv()
	db := &database.DB{}
	if err := db.Setup(); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
	defer db.Close()

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
	message_id_to_callback := map[game_message.Id]func(*game_message.FinalMessage){
		(game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage{}).GetId(): func(message *game_message.FinalMessage) {
			item := (*message).(*game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage)
			if len((*item).ItemTypeDescription) == 0 {
				return
			}
			SaveItemToDb(db, item)
		},
	}

	receiver.Run()
	for {
		PullMessages(receiver, buffer, &messages)
		HandleMessages(&message_id_to_callback, &messages)
		UpdateMessages(&messages)
	}
}

func HandleMessages(message_id_to_callback *map[game_message.Id]func(*game_message.FinalMessage), messages *[]game_message.FinalMessage) {
	for _, message := range *messages {
		if callback, ok := (*message_id_to_callback)[message.GetId()]; ok {
			callback(&message)
		}
	}
}

func UpdateMessages(messages *[]game_message.FinalMessage) {
	(*messages) = (*messages)[:0]
}

func SaveItemToDb(db *database.DB, item *game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage) {
	item_dto := game_resources.Items{
		Id:        item.ObjectGID,
		Name:      "defaultName",
		Price1:    int(item.ItemTypeDescription[0].Prices[0]),
		Price10:   int(item.ItemTypeDescription[0].Prices[1]),
		Price100:  int(item.ItemTypeDescription[0].Prices[2]),
		Timestamp: utils.GetTimestamp(),
	}
	if !db.Exist(item_dto) {
		_, err := db.Save(item_dto)
		if err != nil {
			fmt.Println("Error inserting item: ", err)
		} else {
			fmt.Println("Item inserted successfully")
		}
	} else {
		_, err := db.Update(item_dto)
		if err != nil {
			fmt.Println("Error updating item: ", err)
		} else {
			fmt.Println("Item updated successfully")
		}
	}
}

func SetupDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func PullMessages(receiver network.PacketSniffer, buffer []byte, messages *[]game_message.FinalMessage) {
	select {
	case raw_data := <-receiver.Buffer:
		if len(raw_data) == 0 {
			return
		}
		buffer = append(buffer, raw_data...)
		UnpackMessage(buffer, messages)
	default:
		return
	}
}

func UnpackMessage(buffer []byte, messages *[]game_message.FinalMessage) {
	for len(buffer) > 2 {
		header := game_network.HeaderFromByte(buffer)
		if !header.Valid() {
			fmt.Println("Invalid message: ", header.Id)
			buffer = buffer[:0]
			continue
		}
		size := header.TotalSize()
		if size > len(buffer) {
			fmt.Print("Packet is not complete, waiting for more data...")
			buffer = buffer[:0]
			continue
		}
		data := utils.Buffer{Data: buffer[(2 + header.LenType):size], Pos: 0}
		buffer = buffer[size:]
		message, err := GetMessageFromData(header, data)
		if err == nil {
			*messages = append(*messages, message)
		}
	}
}

func GetMessageFromData(header game_network.Header, data utils.Buffer) (game_message.FinalMessage, error) {
	if game_network.ID_TO_MESSAGE[header.Id] != nil {
		message := game_network.ID_TO_MESSAGE[header.Id]()
		err := message.Deserialize(&data)
		return message, err
	}
	return nil, fmt.Errorf("message with id %d not found", header.Id)
}

func ResetBuffer(buffer *[]byte) []byte {
	return (*buffer)[:0]
}
