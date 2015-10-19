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
}

func TestBotNameSouth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player south name")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestColorsCommand(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("colors color1 color2 color3 color4 color5 color6")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestHandCommandFull(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player north hand color1,1 color2,2 color3,3 color4,4 color5,5 color6,6 color1,7")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestHandCommandSixCards(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("player north hand color1,1 color2,2 color3,3 color4,4 color5,5 color6,6")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestFlagClaimStatusCommand(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag claim-status unclaimed north south unclaimed north south unclaimed north south")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestFlagCardsCommandEmptyFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 1 cards north")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestFlagCardsCommandOneCardFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 1 cards north color1,1")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestFlagCardsCommandFullFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 1 cards north color1,1 color2,2 color3,3")
	assert.True(t, testParser.lastCommandWasKnown)
}
func TestFlagCardsCommandOneCardFlagTwoSouth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 2 cards south color2,2")
	assert.True(t, testParser.lastCommandWasKnown)
}

func TestFlagCardsCommandFullFlagTwoSouth(t *testing.T) {
	testParser := Parser{}
	testParser.ParseString("flag 2 cards south color1,1 color2,2 color3,3")
	assert.True(t, testParser.lastCommandWasKnown)
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
