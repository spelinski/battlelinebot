package board

import (
	"strconv"
	"strings"
)

type Card struct {
	Color  string
	Number int
}

type Flag struct {
	Claimer string
	North   []Card
	South   []Card
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
	cardList := getCardList(cards)
	if flagDirection == "north" {
		b.Flags[flagIndex-1].North = cardList
	} else {
		b.Flags[flagIndex-1].South = cardList
	}
}

func getCardList(cards []string) []Card {
	cardList := []Card{}
	for _, card := range cards {
		cardDetails := strings.Split(card, ",")
		cardNumber, _ := strconv.Atoi(cardDetails[1])
		nextCard := Card{cardDetails[0], cardNumber}
		cardList = append(cardList, nextCard)
	}
	return cardList
}
