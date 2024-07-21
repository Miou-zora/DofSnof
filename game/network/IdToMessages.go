package game

import game "sniffsniff/game/network/messages"

var ID_TO_MESSAGE = map[uint16]game.INewFinalMessage{
	4877: game.CreateBasicPongMessage,
	1770: game.CreateChatAbstractServerMessage,
	6772: game.CreateChatServerMessage,
}
