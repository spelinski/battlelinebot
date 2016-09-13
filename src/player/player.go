package player

import (
    "board"
	"card"
    "fmt"
)

var botName = "SynergyBotTesting"

type Player struct {
	Direction string
	Hand      []card.Card
}

func (p *Player) HandleRespondingToName(direction string) {
	p.Direction = direction
	fmt.Println("player " + direction + " " + botName)
}

func (p *Player) HandleHandUpdate(cards []string, myBoard board.Board) (board.Board) {
	p.Hand = card.GetListOfCardsFromStringArray(cards)
    myBoard.RemoveCardsFromEnemyDeck(p.Hand)
    return myBoard
}
