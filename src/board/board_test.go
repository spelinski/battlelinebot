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
	testCardSlice := []card.Card{card.Card{"color1", 1, 0, 0}}
	assert.Equal(t, testBoard.Flags[0].North, testCardSlice)
}

func TestHandleFlagAddCardCommandFullFlagOneNorth(t *testing.T) {
	testBoard := Board{}
	flagIndex := 1
	flagDirection := "north"
	cards := []string{"color1,1", "color2,2", "color3,3"}
	testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	testCardSlice := []card.Card{card.Card{"color1", 1, 0, 0}, card.Card{"color2", 2, 0, 0}, card.Card{"color3", 3, 0, 0}}
	assert.Equal(t, testBoard.Flags[0].North, testCardSlice)
}

func TestHandleFlagAddCardCommandOneCardFlagTwoSouth(t *testing.T) {
	testBoard := Board{}
	flagIndex := 2
	flagDirection := "south"
	cards := []string{"color2,2"}
	testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	testCardSlice := []card.Card{card.Card{"color2", 2, 0, 0}}
	assert.Equal(t, testBoard.Flags[1].South, testCardSlice)
}

func TestHandleFlagAddCardCommandFullFlagTwoSouth(t *testing.T) {
    testBoard := Board{}
    flagIndex := 2
    flagDirection := "south"
    cards := []string{"color1,1", "color2,2", "color3,3"}
    testBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
    testCardSlice := []card.Card{card.Card{"color1", 1,0,0}, card.Card{"color2", 2,0,0}, card.Card{"color3", 3,0,0}}
    assert.Equal(t, testBoard.Flags[1].South, testCardSlice)
}

func TestGetAllPlayedCards(t *testing.T) {
    testBoard := Board{}
    testBoard.Flags[1].South = []card.Card{card.Card{"color1",1,0,0}, card.Card{"color2", 2,0,0}}
    playedCards := testBoard.GetPlayedCards()
    assert.Equal(t, testBoard.Flags[1].South, playedCards)
}

func TestInitTroopDeck(t *testing.T) {
    testBoard := Board{}
    testBoard.InitDecks()
    expectedTroopDeck := []card.Card{card.Card{"color1",1,0,0}, card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0},
                                    card.Card{"color1",4,0,0}, card.Card{"color1",5,0,0}, card.Card{"color1",6,0,0},
                                    card.Card{"color1",7,0,0}, card.Card{"color1",8,0,0}, card.Card{"color1",9,0,0},
                                    card.Card{"color1",10,0,0}, card.Card{"color2",1,0,0}, card.Card{"color2",2,0,0},
                                    card.Card{"color2",3,0,0}, card.Card{"color2",4,0,0}, card.Card{"color2",5,0,0},
                                    card.Card{"color2",6,0,0}, card.Card{"color2",7,0,0}, card.Card{"color2",8,0,0},
                                    card.Card{"color2",9,0,0}, card.Card{"color2",10,0,0}, card.Card{"color3",1,0,0},
                                    card.Card{"color3",2,0,0}, card.Card{"color3",3,0,0}, card.Card{"color3",4,0,0},
                                    card.Card{"color3",5,0,0}, card.Card{"color3",6,0,0}, card.Card{"color3",7,0,0},
                                    card.Card{"color3",8,0,0}, card.Card{"color3",9,0,0}, card.Card{"color3",10,0,0},
                                    card.Card{"color4",1,0,0}, card.Card{"color4",2,0,0}, card.Card{"color4",3,0,0},
                                    card.Card{"color4",4,0,0}, card.Card{"color4",5,0,0}, card.Card{"color4",6,0,0},
                                    card.Card{"color4",7,0,0}, card.Card{"color4",8,0,0}, card.Card{"color4",9,0,0},
                                    card.Card{"color4",10,0,0}, card.Card{"color5",1,0,0},card.Card{"color5",2,0,0},
                                    card.Card{"color5",3,0,0}, card.Card{"color5",4,0,0}, card.Card{"color5",5,0,0},
                                    card.Card{"color5",6,0,0}, card.Card{"color5",7,0,0}, card.Card{"color5",8,0,0},
                                    card.Card{"color5",9,0,0}, card.Card{"color5",10,0,0}, card.Card{"color6",1,0,0},
                                    card.Card{"color6",2,0,0}, card.Card{"color6",3,0,0}, card.Card{"color6",4,0,0},
                                    card.Card{"color6",5,0,0}, card.Card{"color6",6,0,0}, card.Card{"color6",7,0,0},
                                    card.Card{"color6",8,0,0}, card.Card{"color6",9,0,0}, card.Card{"color6",10,0,0}}
    assert.Equal(t, expectedTroopDeck, testBoard.TroopDeck)
}

func TestGetAllUnplayedCards(t *testing.T) {
    testBoard := Board{}
    testBoard.InitDecks()
    cards := []string{"color1,1", "color2,2"}
    testBoard.HandleFlagAddCardCommand(2, "south", cards)
    expectedUnplayedCards := []card.Card{card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}, card.Card{"color1",4,0,0},
                                    card.Card{"color1",5,0,0}, card.Card{"color1",6,0,0}, card.Card{"color1",7,0,0},
                                    card.Card{"color1",8,0,0}, card.Card{"color1",9,0,0}, card.Card{"color1",10,0,0},
                                    card.Card{"color2",1,0,0}, card.Card{"color2",3,0,0}, card.Card{"color2",4,0,0},
                                    card.Card{"color2",5,0,0}, card.Card{"color2",6,0,0}, card.Card{"color2",7,0,0},
                                    card.Card{"color2",8,0,0}, card.Card{"color2",9,0,0}, card.Card{"color2",10,0,0},
                                    card.Card{"color3",1,0,0}, card.Card{"color3",2,0,0}, card.Card{"color3",3,0,0},
                                    card.Card{"color3",4,0,0}, card.Card{"color3",5,0,0}, card.Card{"color3",6,0,0},
                                    card.Card{"color3",7,0,0}, card.Card{"color3",8,0,0}, card.Card{"color3",9,0,0},
                                    card.Card{"color3",10,0,0}, card.Card{"color4",1,0,0}, card.Card{"color4",2,0,0},
                                    card.Card{"color4",3,0,0}, card.Card{"color4",4,0,0}, card.Card{"color4",5,0,0},
                                    card.Card{"color4",6,0,0}, card.Card{"color4",7,0,0}, card.Card{"color4",8,0,0},
                                    card.Card{"color4",9,0,0}, card.Card{"color4",10,0,0}, card.Card{"color5",1,0,0},
                                    card.Card{"color5",2,0,0}, card.Card{"color5",3,0,0}, card.Card{"color5",4,0,0},
                                    card.Card{"color5",5,0,0}, card.Card{"color5",6,0,0}, card.Card{"color5",7,0,0},
                                    card.Card{"color5",8,0,0}, card.Card{"color5",9,0,0}, card.Card{"color5",10,0,0},
                                    card.Card{"color6",1,0,0}, card.Card{"color6",2,0,0}, card.Card{"color6",3,0,0},
                                    card.Card{"color6",4,0,0}, card.Card{"color6",5,0,0}, card.Card{"color6",6,0,0},
                                    card.Card{"color6",7,0,0}, card.Card{"color6",8,0,0}, card.Card{"color6",9,0,0},
                                    card.Card{"color6",10,0,0}}
    actualUnplayedCards := testBoard.GetUnplayedCards()
    assert.Equal(t, expectedUnplayedCards, actualUnplayedCards)
}

func TestGetAllEnemyAvailableCards(t *testing.T) {
    testBoard := Board{}
    testBoard.InitDecks()
    cards := []string{"color1,1","color2,2"}
    testBoard.HandleFlagAddCardCommand(2, "south", cards)
    expectedUnplayedCards := []card.Card{card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}, card.Card{"color1",4,0,0},
                                    card.Card{"color1",5,0,0}, card.Card{"color1",6,0,0}, card.Card{"color1",7,0,0},
                                    card.Card{"color1",8,0,0}, card.Card{"color1",9,0,0}, card.Card{"color1",10,0,0},
                                    card.Card{"color2",1,0,0}, card.Card{"color2",3,0,0}, card.Card{"color2",4,0,0},
                                    card.Card{"color2",5,0,0}, card.Card{"color2",6,0,0}, card.Card{"color2",7,0,0},
                                    card.Card{"color2",8,0,0}, card.Card{"color2",9,0,0}, card.Card{"color2",10,0,0},
                                    card.Card{"color3",1,0,0}, card.Card{"color3",2,0,0}, card.Card{"color3",3,0,0},
                                    card.Card{"color3",4,0,0}, card.Card{"color3",5,0,0}, card.Card{"color3",6,0,0},
                                    card.Card{"color3",7,0,0}, card.Card{"color3",8,0,0}, card.Card{"color3",9,0,0},
                                    card.Card{"color3",10,0,0}, card.Card{"color4",1,0,0}, card.Card{"color4",2,0,0},
                                    card.Card{"color4",3,0,0}, card.Card{"color4",4,0,0}, card.Card{"color4",5,0,0},
                                    card.Card{"color4",6,0,0}, card.Card{"color4",7,0,0}, card.Card{"color4",8,0,0},
                                    card.Card{"color4",9,0,0}, card.Card{"color4",10,0,0}, card.Card{"color5",1,0,0},
                                    card.Card{"color5",2,0,0}, card.Card{"color5",3,0,0}, card.Card{"color5",4,0,0},
                                    card.Card{"color5",5,0,0}, card.Card{"color5",6,0,0}, card.Card{"color5",7,0,0},
                                    card.Card{"color5",8,0,0}, card.Card{"color5",9,0,0}, card.Card{"color5",10,0,0},
                                    card.Card{"color6",1,0,0}, card.Card{"color6",2,0,0}, card.Card{"color6",3,0,0},
                                    card.Card{"color6",4,0,0}, card.Card{"color6",5,0,0}, card.Card{"color6",6,0,0},
                                    card.Card{"color6",7,0,0}, card.Card{"color6",8,0,0}, card.Card{"color6",9,0,0},
                                    card.Card{"color6",10,0,0}}
    actualUnplayedCards := testBoard.GetEnemyAvailCards()
    assert.Equal(t, expectedUnplayedCards, actualUnplayedCards)
}

func TestRemoveCardsFromEnemyDeck(t *testing.T) {
    testBoard := Board{}
    testBoard.InitDecks()
    cards := []card.Card{card.Card{"color1",1,0,0},card.Card{"color2",2,0,0}}
    testBoard.RemoveCardsFromEnemyDeck(cards)
    expectedUnplayedCards := []card.Card{card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}, card.Card{"color1",4,0,0},
                                    card.Card{"color1",5,0,0}, card.Card{"color1",6,0,0}, card.Card{"color1",7,0,0},
                                    card.Card{"color1",8,0,0}, card.Card{"color1",9,0,0}, card.Card{"color1",10,0,0},
                                    card.Card{"color2",1,0,0}, card.Card{"color2",3,0,0}, card.Card{"color2",4,0,0},
                                    card.Card{"color2",5,0,0}, card.Card{"color2",6,0,0}, card.Card{"color2",7,0,0},
                                    card.Card{"color2",8,0,0}, card.Card{"color2",9,0,0}, card.Card{"color2",10,0,0},
                                    card.Card{"color3",1,0,0}, card.Card{"color3",2,0,0}, card.Card{"color3",3,0,0},
                                    card.Card{"color3",4,0,0}, card.Card{"color3",5,0,0}, card.Card{"color3",6,0,0},
                                    card.Card{"color3",7,0,0}, card.Card{"color3",8,0,0}, card.Card{"color3",9,0,0},
                                    card.Card{"color3",10,0,0}, card.Card{"color4",1,0,0}, card.Card{"color4",2,0,0},
                                    card.Card{"color4",3,0,0}, card.Card{"color4",4,0,0}, card.Card{"color4",5,0,0},
                                    card.Card{"color4",6,0,0}, card.Card{"color4",7,0,0}, card.Card{"color4",8,0,0},
                                    card.Card{"color4",9,0,0}, card.Card{"color4",10,0,0}, card.Card{"color5",1,0,0},
                                    card.Card{"color5",2,0,0}, card.Card{"color5",3,0,0}, card.Card{"color5",4,0,0},
                                    card.Card{"color5",5,0,0}, card.Card{"color5",6,0,0}, card.Card{"color5",7,0,0},
                                    card.Card{"color5",8,0,0}, card.Card{"color5",9,0,0}, card.Card{"color5",10,0,0},
                                    card.Card{"color6",1,0,0}, card.Card{"color6",2,0,0}, card.Card{"color6",3,0,0},
                                    card.Card{"color6",4,0,0}, card.Card{"color6",5,0,0}, card.Card{"color6",6,0,0},
                                    card.Card{"color6",7,0,0}, card.Card{"color6",8,0,0}, card.Card{"color6",9,0,0},
                                    card.Card{"color6",10,0,0}}
    actualUnplayedCards := testBoard.GetEnemyAvailCards()
    assert.Equal(t, expectedUnplayedCards, actualUnplayedCards)
}