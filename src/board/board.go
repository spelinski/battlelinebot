package board

import (
	"card"
)

type Flag struct {
	Claimer string
	North   []card.Card
	South   []card.Card
}

type Board struct {
	Flags [9]Flag
}

func (b *Board) HandleFlagClaimCommand(command []string) {
	for i, claimer := range command {
		b.Flags[i].Claimer = claimer
	}
}

func (b *Board) HandleFlagAddCardCommand(flagIndex int, flagDirection string, cards []string) {
	cardList := card.GetListOfCardsFromStringArray(cards)
	if flagDirection == "north" {
		b.Flags[flagIndex-1].North = cardList
	} else {
		b.Flags[flagIndex-1].South = cardList
	}
}
