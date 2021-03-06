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
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
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
    assert.Equal(t,"play 5 color1,7\n",out )
}

func TestHandleGoPlayCommandFlagOneSideFullNotClaimedPlayerNorth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    flagOneNorthCards := []card.Card{card.Card{"color1",10,0,0},
                                        card.Card{"color1",9,0,0},
                                        card.Card{"color1",8,0,0}}
    testBoard := board.Board{}
    testBoard.InitDecks()
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
    assert.Equal(t,"play 5 color1,7\n",out )
}

func TestHandleGoPlayCommandEmptyBoardNotClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",7,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
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
    assert.Equal(t,"play 5 color4,7\n",out )
}

func TestHandleGoPlayCommandFlagOneSideFullNotClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",7,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",4,0,0}}
    testPlayer.Hand = hand
    flagOneSouthCards := []card.Card{card.Card{"color1",10,0,0},
                                        card.Card{"color1",9,0,0},
                                        card.Card{"color1",8,0,0}}
    testBoard := board.Board{}
    testBoard.InitDecks()
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
    assert.Equal(t,"play 5 color4,7\n",out )
}

func TestHandleGoPlayCommandEmptyBoardClaimedPlayerSouth(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "north"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 4 color1,7\n",out )
}

func TestContinueWedgeOnFlagOne(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color5,6"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color5,5\n",out )
}

func TestContinuePhalanxOnFlagOne(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,4"}
    spoilerCards := []string{"color1,3","color1,5"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    testBoard.HandleFlagAddCardCommand(2, "north", spoilerCards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,4\n",out )
}

func TestContinueBattalionOnFlagOne(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color6",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color5,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)

    //stop Wedge from being viable
    spoilerCardsWedge := []string{"color5,10","color5,8"}
    testBoard.HandleFlagAddCardCommand(2, "north", spoilerCardsWedge)

    //stop Phalanx from being viable
    spoilerCardsPhalanx := []string{"color1,9","color2,9","color3,9"}
    spoilerCardsPhalanxTwo := []string{"color4,9"}
    testBoard.HandleFlagAddCardCommand(3, "north", spoilerCardsPhalanx)
    testBoard.HandleFlagAddCardCommand(4, "north", spoilerCardsPhalanxTwo)
    
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color5,5\n",out )
}

func TestContinueSkirmishOnFlagOne(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1",1,0,0},
                            card.Card{"color2",8,0,0},
                            card.Card{"color3",3,0,0},
                            card.Card{"color4",4,0,0},
                            card.Card{"color5",5,0,0},
                            card.Card{"color5",6,0,0},
                            card.Card{"color1",7,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"

    cards := []string{"color6,2"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)

    //stop Wedge from being viable
    spoilerCardsWedge := []string{"color6,1","color6,3"}
    testBoard.HandleFlagAddCardCommand(2, "north", spoilerCardsWedge)

    //stop Phalanx from being viable
    spoilerCardsPhalanx := []string{"color1,2","color2,2","color3,2"}
    spoilerCardsPhalanxTwo := []string{"color4,2"}
    testBoard.HandleFlagAddCardCommand(3, "north", spoilerCardsPhalanx)
    testBoard.HandleFlagAddCardCommand(4, "north", spoilerCardsPhalanxTwo)

    //Stop Battalion from being viable
    spoilerCardsBattalion := []string{"color6,4","color6,5","color6,6"}
    spoilerCardsBattalionTwo := []string{"color6,7","color6,8","color6,9"}
    testBoard.HandleFlagAddCardCommand(5, "north", spoilerCardsBattalion)
    testBoard.HandleFlagAddCardCommand(6, "north", spoilerCardsBattalionTwo)

    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color3,3\n",out )
}

func TestPlayingOnNewFlagIfNotAbleToContinueFormation(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color4",1,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color6,5"}
    //5 is 4 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(5, "south", cards)

    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 4 color1,2\n",out)
}

func TestShouldNotContinueAFormationWithTwoCardsPlayed(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color4",1,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color6,5","color4,2"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 5 color1,2\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedFirstOneHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color4",1,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color4,3","color4,2"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,1\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedFirstOneLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",1,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color4",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color4,2","color4,3"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,4\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedGapInMiddleFirstOneHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color4", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",3,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color4",3,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color4,4","color4,2"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,3\n",out )
}

func TestShouldFinishWedgeWithTwoCardsPlayedGapInMiddleFirstOneLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color4", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",3,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color4",3,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color3",3,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color4,2","color4,4"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color4,3\n",out )
}

func TestShouldFinishPhalanxWithTwoCardsPlayed(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color4",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color3,6","color4,6"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color2,6\n",out )
}

func TestShouldFinishBattalionWithTwoCardsPlayed(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color6,1","color6,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color6,4\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedFirstLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",7,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color5,8","color6,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,7\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedFirstHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",7,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color5,9","color6,8"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,7\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedGapInMiddleFirstLower(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",8,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color5,7","color6,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,8\n",out )
}

func TestShouldFinishSkirmishWithTwoCardsPlayedGapInMiddleFirstHigher(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",8,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color5,9","color6,7"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color1,8\n",out )
}

func TestGetBestFlagFormationWedge(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "wedge", bestFormation)
}

func TestGetBestFlagFormationPhalanx(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1"}
    spoilerCards := []string{"color1,2"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    testBoard.HandleFlagAddCardCommand(2, "north", spoilerCards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "phalanx", bestFormation)
}

func TestGetBestFlagFormationBattalion(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    testBoard := board.Board{}
    testBoard.InitDecks()
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
    spoilerCards := []string{"color1,2", "color2,1", "color3,1", "color4,1", "color5,1", "color6,1"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    testBoard.HandleFlagAddCardCommand(2, "north", spoilerCards)
    bestFormation := getBestFormation(testBoard.Flags[0].South, testBoard)
    assert.Equal(t, "battalion", bestFormation)
}

func TestGetBestFlagFormationSkirmish(t *testing.T){
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    testBoard := board.Board{}
    testBoard.InitDecks()
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
    testBoard := board.Board{}
    testBoard.InitDecks()
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

func TestShouldNotPlayCardIfItsNotBestFormation(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",7,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color6,8", "color6,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)
    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 5 color2,6\n",out )
}

func TestShouldContinueBattalionWhenNotBestIfNoHostToStart(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    hand := []card.Card{card.Card{"color1", 1,0,0},
                            card.Card{"color1",2,0,0},
                            card.Card{"color2",6,0,0},
                            card.Card{"color2",2,0,0},
                            card.Card{"color3",1,0,0},
                            card.Card{"color3",2,0,0},
                            card.Card{"color6",4,0,0}}
    testPlayer.Hand = hand
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color6,8", "color6,9"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "south", cards)

    //Do not play color6, 10
    cards = []string{"color5,10"}
    testBoard.HandleFlagAddCardCommand(2, "south", cards)
    cards = []string{"color5,9"}
    testBoard.HandleFlagAddCardCommand(3, "south", cards)
    cards = []string{"color5,8"}
    testBoard.HandleFlagAddCardCommand(4, "south", cards)
    cards = []string{"color4,10"}
    testBoard.HandleFlagAddCardCommand(5, "south", cards)
    cards = []string{"color4,9"}
    testBoard.HandleFlagAddCardCommand(6, "south", cards)
    cards = []string{"color4,8"}
    testBoard.HandleFlagAddCardCommand(7, "south", cards)
    cards = []string{"color5,7"}
    testBoard.HandleFlagAddCardCommand(8, "south", cards)
    cards = []string{"color4,7"}
    testBoard.HandleFlagAddCardCommand(9, "south", cards)

    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color6,4\n",out )
}

func TestGetBestEnemyFormationWedge(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    bestFormation := getBestEnemyFormation(testBoard.Flags[0].North, testBoard)
    assert.Equal(t, "wedge", bestFormation)
}

func TestGetBestEnemyFormationFullWedge(t *testing.T) {
    testPlayer := player.Player{}
    testPlayer.Direction = "south"
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"
    cards := []string{"color1,1","color1,2","color1,3"}
    //1 is 0 since this this function is usually for getting it from the engine which is 1 based
    testBoard.HandleFlagAddCardCommand(1, "north", cards)
    bestFormation := getBestEnemyFormation(testBoard.Flags[0].North, testBoard)
    assert.Equal(t, "wedge", bestFormation)
}

func TestShouldPlayToContinueWedgeOnOneAwayIfItIsBothTwoAwayAndOne(t *testing.T) {
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"

    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    handToBeAdded := []string{"color2,5","color3,6","color5,10","color6,8","color2,4","color1,9","color3,8"}
    testBoard = testPlayer.HandleHandUpdate(handToBeAdded, testBoard)
    cards := []string{"color1,2"}
    testBoard.HandleFlagAddCardCommand(2, "north", cards)
    cards = []string{"color3,4"}
    testBoard.HandleFlagAddCardCommand(3, "north", cards)
    cards = []string{"color5,6"}
    testBoard.HandleFlagAddCardCommand(4, "north", cards)
    cards = []string{"color4,10","color4,9"}
    testBoard.HandleFlagAddCardCommand(4, "south", cards)
    cards = []string{"color6,10"}
    testBoard.HandleFlagAddCardCommand(5, "north", cards)
    cards = []string{"color1,10"}
    testBoard.HandleFlagAddCardCommand(5, "south", cards)
    cards = []string{"color4,5","color4,4"}
    testBoard.HandleFlagAddCardCommand(6, "north", cards)
    cards = []string{"color2,7","color2,6"}
    testBoard.HandleFlagAddCardCommand(6, "south", cards)


    cards = []string{"color3,7"}
    testBoard.HandleFlagAddCardCommand(7, "north", cards)
    cards = []string{"color6,5","color6,3"}
    testBoard.HandleFlagAddCardCommand(7, "south", cards)


    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 7 color3,8\n",out )
}

func TestShouldPlayFirstCardWhereOpponentDid(t *testing.T) {
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"

    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    handToBeAdded := []string{"color2,5","color3,6","color5,10","color6,8","color2,4","color1,9","color3,8"}
    testBoard = testPlayer.HandleHandUpdate(handToBeAdded, testBoard)
    cards := []string{"color4,10"}
    testBoard.HandleFlagAddCardCommand(4, "south", cards)


    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 4 color5,10\n",out )
}

func TestShouldPlayFirstCardWhereOpponentDidSecondPlay(t *testing.T) {
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"

    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    handToBeAdded := []string{"color2,5","color3,6","color5,10","color6,8","color2,4","color1,9","color3,8"}
    testBoard = testPlayer.HandleHandUpdate(handToBeAdded, testBoard)
    cards := []string{"color4,10"}
    testBoard.HandleFlagAddCardCommand(3, "south", cards)
    cards = []string{"color3,10"}
    testBoard.HandleFlagAddCardCommand(3, "north", cards)
    cards = []string{"color6,10"}
    testBoard.HandleFlagAddCardCommand(4, "south", cards)


    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 4 color6,8\n",out )
}

func TestShouldPlayFirstCardWhereOpponentDidFlagOne(t *testing.T) {
    testBoard := board.Board{}
    testBoard.InitDecks()
    testBoard.Flags[0].Claimer = "unclaimed"
    testBoard.Flags[1].Claimer = "unclaimed"
    testBoard.Flags[2].Claimer = "unclaimed"
    testBoard.Flags[3].Claimer = "unclaimed"
    testBoard.Flags[4].Claimer = "unclaimed"
    testBoard.Flags[5].Claimer = "unclaimed"
    testBoard.Flags[6].Claimer = "unclaimed"
    testBoard.Flags[7].Claimer = "unclaimed"
    testBoard.Flags[8].Claimer = "unclaimed"

    testPlayer := player.Player{}
    testPlayer.Direction = "north"
    handToBeAdded := []string{"color2,5","color3,6","color5,8","color6,8","color2,4","color5,9","color3,8"}
    testBoard = testPlayer.HandleHandUpdate(handToBeAdded, testBoard)
    cards := []string{"color4,10"}
    testBoard.HandleFlagAddCardCommand(3, "south", cards)
    cards = []string{"color3,10"}
    testBoard.HandleFlagAddCardCommand(3, "north", cards)
    cards = []string{"color6,10"}
    testBoard.HandleFlagAddCardCommand(4, "south", cards)
    cards = []string{"color1,10"}
    testBoard.HandleFlagAddCardCommand(4, "north", cards)
    cards = []string{"color2,10"}
    testBoard.HandleFlagAddCardCommand(1, "south", cards)


    out := handleStdOut(testPlayer, testBoard)
    assert.Equal(t,"play 1 color5,8\n",out )
}
