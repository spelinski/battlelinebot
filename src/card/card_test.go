package card

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetListOfCardsFromStringArray(t *testing.T) {
	cards := []string{"color1,1", "color2,2", "color3,3"}
	returnedList := GetListOfCardsFromStringArray(cards)
	testCardSlice := []Card{Card{"color1", 1}, Card{"color2", 2}, Card{"color3", 3}}
	assert.Equal(t, returnedList, testCardSlice)
}
