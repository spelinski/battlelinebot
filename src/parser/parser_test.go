package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyString(t *testing.T) {
	testParser := Parser{}
	assert.NotNil(t, testParser.ParseString(""))
}

func TestBotNameNorth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("player north name"))
}

func TestBotNameSouth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("player south name"))
}

func TestColorsCommand(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("colors color1 color2 color3 color4 color5 color6"))
}

func TestHandCommandFull(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("player north hand color1,1 color2,2 color3,3 color4,4 color5,5 color6,6 color1,7"))
}

func TestHandCommandSixCards(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("player north hand color1,1 color2,2 color3,3 color4,4 color5,5 color6,6"))
}

func TestFlagClaimStatusCommand(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("flag claim-status unclaimed north south unclaimed north south unclaimed north south"))
}

func TestFlagCardsCommandEmptyFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("flag 1 cards north"))
}

func TestFlagCardsCommandOneCardFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("flag 1 cards north color1,1"))
}

func TestFlagCardsCommandFullFlagOneNorth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("flag 1 cards north color1,1 color2,2 color3,3"))
}
func TestFlagCardsCommandOneCardFlagTwoSouth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("flag 2 cards south color2,2"))
}

func TestFlagCardsCommandFullFlagTwoSouth(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("flag 2 cards south color1,1 color2,2 color3,3"))
}

func TestOppentPlayCommand(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("opponent play 1 color1,1"))
}

func TestGoPlayCommand(t *testing.T) {
	testParser := Parser{}
	assert.Nil(t, testParser.ParseString("go play-card"))
}
