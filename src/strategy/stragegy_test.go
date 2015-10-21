package strategy

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "board"
    "player"
    "card"
    "os"
    "bytes"
    "io"
)

func handleStdOut(testPlayer player.Player, boardInfo board.Board) string {
    oldStdOut := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w
    HandleGoPlayCommand(testPlayer, boardInfo)
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

func TestHandleGoPlayCommandEmptyBoardNotClaimedPlayerNorth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",4},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",7}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,out, "play 1 color1,1\n" )
}

func TestHandleGoPlayCommandFlagOneSideFullNotClaimedPlayerNorth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",4},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",7}}
    testPlayer.Hand = hand
    flagOneNorthCards := []card.Card{card.Card{"color1",10},
                                        card.Card{"color1",9},
                                        card.Card{"color1",8}}
    testBoard := board.Board{}
    testBoard.Flags[0].North = flagOneNorthCards
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,out, "play 2 color1,1\n" )
}

func TestHandleGoPlayCommandEmptyBoardNotClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",4},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",7}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,out, "play 1 color1,1\n" )
}

func TestHandleGoPlayCommandFlagOneSideFullNotClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",4},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",7}}
    testPlayer.Hand = hand
    flagOneSouthCards := []card.Card{card.Card{"color1",10},
                                        card.Card{"color1",9},
                                        card.Card{"color1",8}}
    testBoard := board.Board{}
    testBoard.Flags[0].South = flagOneSouthCards
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,out, "play 2 color1,1\n" )
}

func TestHandleGoPlayCommandEmptyBoardClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",4},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",7}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "north"
    testBoard.Flags[1].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,out, "play 2 color1,1\n" )
}