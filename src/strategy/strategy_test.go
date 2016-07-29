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
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,7\n",out )
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
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 2 color1,7\n",out )
}

func TestHandleGoPlayCommandEmptyBoardNotClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",7},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,7\n",out )
}

func TestHandleGoPlayCommandFlagOneSideFullNotClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",2},
                            card.Card{"color3",3},
                            card.Card{"color4",7},
                            card.Card{"color5",5},
                            card.Card{"color6",6},
                            card.Card{"color1",4}}
    testPlayer.Hand = hand
    flagOneSouthCards := []card.Card{card.Card{"color1",10},
                                        card.Card{"color1",9},
                                        card.Card{"color1",8}}
    testBoard := board.Board{}
    testBoard.Flags[0].South = flagOneSouthCards
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 2 color4,7\n",out )
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
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 2 color1,7\n",out )
}

func TestContinueWedgeOnFlagOne(t *testing.T) {
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
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color5",6}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color5,5\n",out )
}

func TestContinuePhalanxOnFlagOne(t *testing.T) {
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
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color1",4}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,4\n",out )
}

func TestContinueBattalionOnFlagOne(t *testing.T) {
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
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color5",9}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color5,5\n",out )
}

func TestContinueSkirmishOnFlagOne(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1},
                            card.Card{"color2",8},
                            card.Card{"color3",3},
                            card.Card{"color4",4},
                            card.Card{"color5",5},
                            card.Card{"color5",6},
                            card.Card{"color1",7}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color6",2}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,1\n",out )
}

func TestPlayingOnNewFlagIfNotAbleToContinueFormation(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",1},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color4",1}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color6",5}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 2 color1,2\n",out)
}

func TestShouldNotContinueAFormationWithTwoCardsPlayed(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",1},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color4",1}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color6",5},
                                                card.Card{"color4",2}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 2 color1,2\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedFirstOneHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",1},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color4",1}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color4",3},
                                                card.Card{"color4",2}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,1\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedFirstOneLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",1},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color4",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color4",2},
                                                card.Card{"color4",3}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,4\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedGapInMiddleFirstOneHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color4", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",3},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color4",3}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color4",4},
                                                card.Card{"color4",2}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,3\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedGapInMiddleFirstOneLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color4", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",3},
                            card.Card{"color2",2},
                            card.Card{"color4",3},
                            card.Card{"color3",2},
                            card.Card{"color3",3}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color4",2},
                                                card.Card{"color4",4}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,3\n",out )
}

func TestShouldFinishPhalanxWithTwoCardsPlayed(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color4",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color3",6},
                                                card.Card{"color4",6}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color2,6\n",out )
}

func TestShouldFinishBattalionWithTwoCardsPlayed(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",2},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color6",1},
                                                card.Card{"color6",9}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color6,4\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedFirstLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color5",8},
                                                card.Card{"color6",9}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,7\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedFirstHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color5",9},
                                                card.Card{"color6",8}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,7\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedGapInMiddleFirstLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",8},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color5",7},
                                                card.Card{"color6",9}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,8\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedGapInMiddleFirstHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",8},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color5",9},
                                                card.Card{"color6",7}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,8\n",out )
}

func TestGettingCombinationOfCardsByOne(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1}, card.Card{"color1",2}, card.Card{"color1",3}}
    actualCombos := cardCombinations (threeCardList, 1)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",2}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",3}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGettingCombinationOfCardsByTwo(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1}, card.Card{"color1",2}, card.Card{"color1",3}}
    actualCombos := cardCombinations (threeCardList, 2)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1},card.Card{"color1",2}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1},card.Card{"color1",3}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",2},card.Card{"color1",3}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGettingCombinationOfCardsByThree(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1}, card.Card{"color1",2}, card.Card{"color1",3}}
    actualCombos := cardCombinations (threeCardList, 3)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1},card.Card{"color1",2},card.Card{"color1",3}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGetBestFlagFormationWedge(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 9},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitTroopDeck()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1", "color1,2"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "wedge", bestFormation)
}

func TestGetBestFlagFormationPhalanx(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 9},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",8},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitTroopDeck()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1", "color2,1"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "phalanx", bestFormation)
}

func TestGetBestFlagFormationBattalion(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 9},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",8},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitTroopDeck()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1", "color1,5"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "battalion", bestFormation)
}

func TestGetBestFlagFormationSkirmish(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 9},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",8},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitTroopDeck()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1", "color2,3"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "skirmish", bestFormation)
}

func TestGetBestFlagFormationHost(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 9},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",8},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitTroopDeck()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1", "color2,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "host", bestFormation)
}

/*func TestShouldNotPlayCardIfItsNotBestFormation(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1},
                            card.Card{"color1",7},
                            card.Card{"color2",6},
                            card.Card{"color2",2},
                            card.Card{"color3",1},
                            card.Card{"color3",2},
                            card.Card{"color6",4}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    testBoard.Flags[0].South = []card.Card{card.Card{"color6",8},
                                                card.Card{"color6",9}}
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 2 color2,6\n",out )
}*/
