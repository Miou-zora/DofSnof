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
	message_id_to_callback := map[game_message.Id]func(*game_message.FinalMessage){
		(game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage{}).GetId(): func(message *game_message.FinalMessage) {
			item := (*message).(*game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage)
			if len((*item).ItemTypeDescription) == 0 {
				return
			}
			SaveItemToDb(db, item)
		},
	}

	device, err := network.AskForDevice()
	if err != nil {
		log.Fatalf("Error getting device: %v", err)
	}
	receiver := network.PacketSniffer{
		Buffer: make(chan []byte, MaxBufferSize),
		Filter: DefaultFilter,
		Device: device,
	}
	message_processor := game_network.MessageProcessor{}
	receiver.Run()
	message_processor.Run(&receiver, &message_id_to_callback)
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
