package player

import (
	"card"
	"fmt"
)

var botName = "SynergyBot"

type Player struct {
	Direction string
	Hand      []card.Card
}

func (p *Player) HandleRespondingToName(direction string) {
	p.Direction = direction
	fmt.Println("player " + direction + " " + botName)
}

func (p *Player) HandleHandUpdate(cards []string) {
	p.Hand = card.GetListOfCardsFromStringArray(cards)
}
