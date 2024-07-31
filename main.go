package main

import (
	"fmt"
	"log"
	game_network "sniffsniff/game/network"
	game_message "sniffsniff/game/network/messages"
	game_resources "sniffsniff/game/resources"
	"sniffsniff/network"
	"sniffsniff/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"os"

	"github.com/joho/godotenv"
)

const (
	MaxBufferSize = 4096
	DefaultFilter = "tcp src port 5555"
)

func main() {
	db, err := SetupDb()
	if err != nil {
		log.Fatalln(err)
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
	receiver.Run()
	for {
		PullMessages(receiver, buffer, &messages)
		for _, message := range messages {
			// if message is a type of game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage
			// then call a function
			fmt.Println("azeazeMessage: ", message)
			if _, ok := message.(*game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage); ok {
				SaveItemToDb(db, message.(*game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage))
			}
		}
		messages = messages[:0]
	}
}

func SaveItemToDb(db *sqlx.DB, item *game_message.ExchangeTypesItemsExchangerDescriptionForUserMessage) {
	if len(item.ItemTypeDescription) < 1 || len(item.ItemTypeDescription[0].Prices) < 3 {
		fmt.Println("Error: item is not valid")
		return
	}
	item_dto := game_resources.Item{
		Id:        item.ObjectGID,
		Name:      "defaultName",
		Price1:    int(item.ItemTypeDescription[0].Prices[0]),
		Price10:   int(item.ItemTypeDescription[0].Prices[1]),
		Price100:  int(item.ItemTypeDescription[0].Prices[2]),
		Timestamp: utils.GetTimestamp(),
	}
	// if item is already in the db, update it
	// else insert it
	if _, err := db.NamedExec("SELECT * FROM items WHERE id = :id", item_dto); err == nil {
		_, err := db.NamedExec("UPDATE items SET price1 = :price1, price10 = :price10, price100 = :price100, timestamp = :timestamp WHERE id = :id", item_dto)
		if err != nil {
			fmt.Println("Error updating item: ", err)
		} else {
			fmt.Println("Item updated successfully")
		}
		return
	}
	_, err := db.NamedExec("INSERT INTO items (id, name, price1, price10, price100, timestamp) VALUES (:id, :name, :price1, :price10, :price100, :timestamp)", item_dto)
	if err != nil {
		fmt.Println("Error inserting item: ", err)
	} else {
		fmt.Println("Item inserted successfully")
	}
}

func SetupDb() (*sqlx.DB, error) {
	SetupDotEnv()
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	db, err := sqlx.Connect("postgres", "user="+user+" dbname="+dbname+" sslmode=disable password="+password+" host="+host)
	if err != nil {
		return db, err
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	return db, err
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
		UnstackMessages(buffer, messages)
	default:
		return
	}
}

func UnstackMessages(buffer []byte, messages *[]game_message.FinalMessage) {
	for len(buffer) > 2 {
		header := game_network.HeaderFromByte(buffer)
		if !header.IsValid() {
			fmt.Println("Invalid message: ", header.Id)
			buffer = buffer[:0]
			continue
		}
		size := header.GetSize()
		if size > len(buffer) {
			fmt.Print("Packet is not complete, waiting for more data...")
			buffer = buffer[:0]
			continue
		}
		data := utils.Buffer{Data: buffer[(2 + header.LenType):size], Pos: 0}
		buffer = buffer[size:]
		message, err := GetMessageFromData(header, data)
		if err != nil {
			continue
		} else {
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
