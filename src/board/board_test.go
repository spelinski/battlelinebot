package board

import (
	"card"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleFlagClaimCommand(t *testing.T) {
	testBoard := Board{}
	commandArray := []string{"unclaimed", "north", "south", "unclaimed", "north", "south", "unclaimed", "north", "south"}
	testBoard.HandleFlagClaimCommand(commandArray)
	assert.Equal(t, testBoard.Flags[0].Claimer, "unclaimed")
	assert.Equal(t, testBoard.Flags[1].Claimer, "north")
	assert.Equal(t, testBoard.Flags[2].Claimer, "south")
	assert.Equal(t, testBoard.Flags[3].Claimer, "unclaimed")
	assert.Equal(t, testBoard.Flags[4].Claimer, "north")
	assert.Equal(t, testBoard.Flags[5].Claimer, "south")
	assert.Equal(t, testBoard.Flags[6].Claimer, "unclaimed")
	assert.Equal(t, testBoard.Flags[7].Claimer, "north")
	assert.Equal(t, testBoard.Flags[8].Claimer, "south")
}

func TestHandleFlagAddCardCommandOneCardFlagOneNorth(t *testing.T) {
	testBoard := Board{}
	flagIndex := 1
	flagDirection := "north"
	cards := []string{"color1,1"}
	testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	testCardSlice := []card.Card{card.Card{"color1", 1}}
	assert.Equal(t, testBoard.Flags[0].North, testCardSlice)
}

func TestHandleFlagAddCardCommandFullFlagOneNorth(t *testing.T) {
	testBoard := Board{}
	flagIndex := 1
	flagDirection := "north"
	cards := []string{"color1,1", "color2,2", "color3,3"}
	testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	testCardSlice := []card.Card{card.Card{"color1", 1}, card.Card{"color2", 2}, card.Card{"color3", 3}}
	assert.Equal(t, testBoard.Flags[0].North, testCardSlice)
}

func TestHandleFlagAddCardCommandOneCardFlagTwoSouth(t *testing.T) {
	testBoard := Board{}
	flagIndex := 2
	flagDirection := "south"
	cards := []string{"color2,2"}
	testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	testCardSlice := []card.Card{card.Card{"color2", 2}}
	assert.Equal(t, testBoard.Flags[1].South, testCardSlice)
}

func TestHandleFlagAddCardCommandFullFlagTwoSouth(t *testing.T) {
	testBoard := Board{}
	flagIndex := 2
	flagDirection := "south"
	cards := []string{"color1,1", "color2,2", "color3,3"}
	testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	testCardSlice := []card.Card{card.Card{"color1", 1}, card.Card{"color2", 2}, card.Card{"color3", 3}}
	assert.Equal(t, testBoard.Flags[1].South, testCardSlice)
}
