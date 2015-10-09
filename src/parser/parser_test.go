package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyString(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("")
	assert.False(t, testParser.lastCommandWasKnown)
}

func TestBotNameNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player north name")
	assert.True(t, testParser.lastCommandWasKnown)
	assert.Equal(t, testParser.visualName, BotVisualName)
	assert.Equal(t, testParser.direction, "north")
}

func TestBotNameSouth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player south name")
	assert.True(t, testParser.lastCommandWasKnown)
	assert.Equal(t, testParser.visualName, BotVisualName)
	assert.Equal(t, testParser.direction, "south")
}

func TestColorsCommand(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("colors color1 color2 color3 color4 color5 color6")
	assert.True(t, testParser.lastCommandWasKnown)
	assert.Equal(t, testParser.colors[0], "color1")
	assert.Equal(t, testParser.colors[1], "color2")
	assert.Equal(t, testParser.colors[2], "color3")
	assert.Equal(t, testParser.colors[3], "color4")
	assert.Equal(t, testParser.colors[4], "color5")
	assert.Equal(t, testParser.colors[5], "color6")
}

func TestHandCommandFull(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player north hand color1,1 color2,2 color3,3 color4,4 color5,5 color6,6 color1,7")
	card1 := Card{"color1", 1}
	card2 := Card{"color2", 2}
	card3 := Card{"color3", 3}
	card4 := Card{"color4", 4}
	card5 := Card{"color5", 5}
	card6 := Card{"color6", 6}
	card7 := Card{"color1", 7}
	assert.True(t, testParser.lastCommandWasKnown)
	assert.Equal(t, testParser.hand[0], card1)
	assert.Equal(t, testParser.hand[1], card2)
	assert.Equal(t, testParser.hand[2], card3)
	assert.Equal(t, testParser.hand[3], card4)
	assert.Equal(t, testParser.hand[4], card5)
	assert.Equal(t, testParser.hand[5], card6)
	assert.Equal(t, testParser.hand[6], card7)
}

func TestHandCommandSixCards(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player north hand color1,1 color2,2 color3,3 color4,4 color5,5 color6,6")
	card1 := Card{"color1", 1}
	card2 := Card{"color2", 2}
	card3 := Card{"color3", 3}
	card4 := Card{"color4", 4}
	card5 := Card{"color5", 5}
	card6 := Card{"color6", 6}
	assert.True(t, testParser.lastCommandWasKnown)
	assert.Equal(t, testParser.hand[0], card1)
	assert.Equal(t, testParser.hand[1], card2)
	assert.Equal(t, testParser.hand[2], card3)
	assert.Equal(t, testParser.hand[3], card4)
	assert.Equal(t, testParser.hand[4], card5)
	assert.Equal(t, testParser.hand[5], card6)
}

func TestFlagClaimStatusCommand(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag claim-status unclaimed north south unclaimed north south unclaimed north south")
	assert.True(t, testParser.lastCommandWasKnown)
	assert.Equal(t, testParser.board.flags[0].claimer, "unclaimed")
	assert.Equal(t, testParser.board.flags[1].claimer, "north")
	assert.Equal(t, testParser.board.flags[2].claimer, "south")
	assert.Equal(t, testParser.board.flags[3].claimer, "unclaimed")
	assert.Equal(t, testParser.board.flags[4].claimer, "north")
	assert.Equal(t, testParser.board.flags[5].claimer, "south")
	assert.Equal(t, testParser.board.flags[6].claimer, "unclaimed")
	assert.Equal(t, testParser.board.flags[7].claimer, "north")
	assert.Equal(t, testParser.board.flags[8].claimer, "south")
}

func TestFlagCardsCommandEmptyFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 1 cards north")
	assert.False(t, testParser.lastCommandWasKnown)
}

func TestFlagCardsCommandOneCardFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 1 cards north color1,1")
	assert.True(t, testParser.lastCommandWasKnown)
	testCardSlice := []Card{Card{"color1", 1}}
	assert.Equal(t, testParser.board.flags[0].north, testCardSlice)
}

func TestFlagCardsCommandFullFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 1 cards north color1,1 color2,2 color3,3")
	assert.True(t, testParser.lastCommandWasKnown)
	testCardSlice := []Card{Card{"color1", 1}, Card{"color2", 2}, Card{"color3", 3}}
	assert.Equal(t, testParser.board.flags[0].north, testCardSlice)
}
func TestFlagCardsCommandOneCardFlagTwoSouth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 2 cards south color2,2")
	assert.True(t, testParser.lastCommandWasKnown)
	testCardSlice := []Card{Card{"color2", 2}}
	assert.Equal(t, testParser.board.flags[1].south, testCardSlice)
}

func TestFlagCardsCommandFullFlagTwoSouth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 2 cards south color1,1 color2,2 color3,3")
	assert.True(t, testParser.lastCommandWasKnown)
	testCardSlice := []Card{Card{"color1", 1}, Card{"color2", 2}, Card{"color3", 3}}
	assert.Equal(t, testParser.board.flags[1].south, testCardSlice)
}

func TestOppentPlayCommand(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("opponent play 1 color1,1")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestGoPlayCommand(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("go play-card")
	assert.True(t, testParser.lastCommandWasKnown)
}
