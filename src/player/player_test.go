package player

import (
    "board"
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
    testBoard := board.Board{}
    testBoard.InitDecks()
	testPlayer := Player{}
	testPlayer.HandleHandUpdate([]string{"color1,1", "color2,2", "color3,3", "color4,4", "color5,5", "color6,6", "color1,7"}, testBoard)
	card1 := card.Card{"color1", 1,0,0}
	card2 := card.Card{"color2", 2,0,0}
	card3 := card.Card{"color3", 3,0,0}
	card4 := card.Card{"color4", 4,0,0}
	card5 := card.Card{"color5", 5,0,0}
	card6 := card.Card{"color6", 6,0,0}
	card7 := card.Card{"color1", 7,0,0}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
	assert.Equal(t, testPlayer.Hand[6], card7)
}

func TestHandleHandUpdateFullTwice(t *testing.T) {
    testBoard := board.Board{}
    testBoard.InitDecks()
	testPlayer := Player{}
	testPlayer.HandleHandUpdate([]string{"color1,1", "color2,2", "color3,3", "color4,4", "color5,5", "color6,6", "color1,7"}, testBoard)
	card1 := card.Card{"color1", 1,0,0}
	card2 := card.Card{"color2", 2,0,0}
	card3 := card.Card{"color3", 3,0,0}
	card4 := card.Card{"color4", 4,0,0}
	card5 := card.Card{"color5", 5,0,0}
	card6 := card.Card{"color6", 6,0,0}
	card7 := card.Card{"color1", 7,0,0}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
	assert.Equal(t, testPlayer.Hand[6], card7)

	testPlayer.HandleHandUpdate([]string{"color1,10", "color2,9", "color3,8", "color4,7", "color5,6", "color6,5", "color1,4"}, testBoard)
	card1 = card.Card{"color1", 10,0,0}
	card2 = card.Card{"color2", 9,0,0}
	card3 = card.Card{"color3", 8,0,0}
	card4 = card.Card{"color4", 7,0,0}
	card5 = card.Card{"color5", 6,0,0}
	card6 = card.Card{"color6", 5,0,0}
	card7 = card.Card{"color1", 4,0,0}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
	assert.Equal(t, testPlayer.Hand[6], card7)
}

func TestHandleHandUpdateSixCards(t *testing.T) {
    testBoard := board.Board{}
    testBoard.InitDecks()
	testPlayer := Player{}
	testPlayer.HandleHandUpdate([]string{"color1,1", "color2,2", "color3,3", "color4,4", "color5,5", "color6,6"},testBoard)
	card1 := card.Card{"color1", 1,0,0}
	card2 := card.Card{"color2", 2,0,0}
	card3 := card.Card{"color3", 3,0,0}
	card4 := card.Card{"color4", 4,0,0}
	card5 := card.Card{"color5", 5,0,0}
	card6 := card.Card{"color6", 6,0,0}
	assert.Equal(t, testPlayer.Hand[0], card1)
	assert.Equal(t, testPlayer.Hand[1], card2)
	assert.Equal(t, testPlayer.Hand[2], card3)
	assert.Equal(t, testPlayer.Hand[3], card4)
	assert.Equal(t, testPlayer.Hand[4], card5)
	assert.Equal(t, testPlayer.Hand[5], card6)
}

func TestHandleHandUpdateRemoveCardsFromEnemyDeck(t *testing.T) {
    newTestBoard := board.Board{}
    newTestBoard.InitDecks()
    testPlayer := Player{}
    newTestBoard = testPlayer.HandleHandUpdate([]string{"color2,2"}, newTestBoard)
    cards := []string{"color1,1"}
    newTestBoard.HandleFlagAddCardCommand(2, "south", cards)
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
    actualUnplayedCards := newTestBoard.GetEnemyAvailCards()
    assert.Equal(t, expectedUnplayedCards, actualUnplayedCards)
}
