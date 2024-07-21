package game

import "sniffsniff/game/network/messages"

var ID_TO_MESSAGE = map[int]game.INewFinalMessage{
	4877: game.CreateBasicPongMessage,
}