package game

import (
	"fmt"

	game_message "sniffsniff/game/network/messages"
	"sniffsniff/utils"
)

func GetMessageFromData(header Header, data utils.Buffer) (game_message.FinalMessage, error) {
	if ID_TO_MESSAGE[header.Id] != nil {
		message := ID_TO_MESSAGE[header.Id]()
		err := message.Deserialize(&data)
		return message, err
	}
	return nil, fmt.Errorf("message with id %d not found", header.Id)
}
