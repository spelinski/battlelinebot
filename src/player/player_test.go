package player

import (
	"bytes"
	"card"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func handleStdOut(testPlayer *Player, direction string) string {
	oldStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	testPlayer.HandleRespondingToName(direction)
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = oldStdOut
	out := <-outC
	return out
}

func TestHandleRespondingToNameNorth(t *testing.T) {
	testPlayer := Player{}
	out := handleStdOut(&testPlayer, "north")

	assert.Equal(t, testPlayer.Direction, "north")
	assert.Equal(t, "player north SynergyBot\n", out)
}

func TestHandleRespondingToNameSouth(t *testing.T) {
	testPlayer := Player{}
	out := handleStdOut(&testPlayer, "south")
	assert.Equal(t, testPlayer.Direction, "south")
	assert.Equal(t, "player south SynergyBot\n", out)
}

func TestHandleHandUpdateFull(t *testing.T) {
	testPlayer := Player{}
	testPlayer.HandleHandUpdate([]string{"color1,1", "color2,2", "color3,3", "color4,4", "color5,5", "color6,6", "color1,7"})
	card1 := card.Card{"color1", 1}
	card2 := card.Card{"color2", 2}
	card3 := card.Card{"color3", 3}
	card4 := card.Card{"color4", 4}
	card5 := card.Card{"color5", 5}
	card6 := card.Card{"color6", 6}
	card7 := card.Card{"color1", 7}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
	assert.Equal(t, testPlayer.Hand[6], card7)
}

func TestHandleHandUpdateFullTwice(t *testing.T) {
	testPlayer := Player{}
	testPlayer.HandleHandUpdate([]string{"color1,1", "color2,2", "color3,3", "color4,4", "color5,5", "color6,6", "color1,7"})
	card1 := card.Card{"color1", 1}
	card2 := card.Card{"color2", 2}
	card3 := card.Card{"color3", 3}
	card4 := card.Card{"color4", 4}
	card5 := card.Card{"color5", 5}
	card6 := card.Card{"color6", 6}
	card7 := card.Card{"color1", 7}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
	assert.Equal(t, testPlayer.Hand[6], card7)

	testPlayer.HandleHandUpdate([]string{"color1,10", "color2,9", "color3,8", "color4,7", "color5,6", "color6,5", "color1,4"})
	card1 = card.Card{"color1", 10}
	card2 = card.Card{"color2", 9}
	card3 = card.Card{"color3", 8}
	card4 = card.Card{"color4", 7}
	card5 = card.Card{"color5", 6}
	card6 = card.Card{"color6", 5}
	card7 = card.Card{"color1", 4}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
	assert.Equal(t, testPlayer.Hand[6], card7)
}

func TestHandleHandUpdateSixCards(t *testing.T) {
	testPlayer := Player{}
	testPlayer.HandleHandUpdate([]string{"color1,1", "color2,2", "color3,3", "color4,4", "color5,5", "color6,6"})
	card1 := card.Card{"color1", 1}
	card2 := card.Card{"color2", 2}
	card3 := card.Card{"color3", 3}
	card4 := card.Card{"color4", 4}
	card5 := card.Card{"color5", 5}
	card6 := card.Card{"color6", 6}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
}
