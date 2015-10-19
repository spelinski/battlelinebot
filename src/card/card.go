package card

import (
	"strconv"
	"strings"
)

type Card struct {
	Color  string
	Number int
}

func GetListOfCardsFromStringArray(cards []string) []Card {
	cardList := []Card{}
	for _, currentCard := range cards {
		cardDetails := strings.Split(currentCard, ",")
		cardNumber, _ := strconv.Atoi(cardDetails[1])
		nextCard := Card{cardDetails[0], cardNumber}
		cardList = append(cardList, nextCard)
	}
	return cardList
}
