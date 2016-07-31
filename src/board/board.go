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
    TroopDeck []card.Card
}

func (b *Board) InitTroopDeck() {
    colorOptions := []string{"color1","color2","color3","color4","color5","color6"}
    for colorIndex := range colorOptions{
        for i := 1; i <= 10; i++ {
            b.TroopDeck = append(b.TroopDeck, card.Card{colorOptions[colorIndex],i, 0, 0})
        }
    }
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
    b.removeCardsFromTroopDeck(cardList)
}

func (b *Board) removeCardsFromTroopDeck(cardsToRemove []card.Card) {
    for removeCardIndex := range cardsToRemove {
        for troopDeckIndex := range b.TroopDeck {
            if b.TroopDeck[troopDeckIndex] == cardsToRemove[removeCardIndex] {
                b.TroopDeck = append(b.TroopDeck[:troopDeckIndex], b.TroopDeck[troopDeckIndex+1:]...)
                break
            }
        }
    }
}

func (b *Board) GetPlayedCards() ([]card.Card) {
    playedCards := []card.Card{}
    for index := range b.Flags {
        for northIndex := range b.Flags[index].North{
            playedCards = append(playedCards, b.Flags[index].North[northIndex])
        }
        for southIndex := range b.Flags[index].South{
            playedCards = append(playedCards, b.Flags[index].South[southIndex])
        }
    }
    return playedCards
}

func (b *Board) GetUnplayedCards() ([]card.Card) {
    return b.TroopDeck
}
