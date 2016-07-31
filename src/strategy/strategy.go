package strategy

import (
    "board"
    "player"
    "utilities"
    "fmt"
    "card"
    "strconv"
)

func HandleGoPlayCommand(playerInfo player.Player, boardInfo board.Board) {
    playCard,flagToPlay := getBestCardAndFlagToPlayOn(boardInfo.Flags,playerInfo.Hand,playerInfo.Direction, boardInfo)
    fmt.Println("play " + strconv.Itoa(flagToPlay) + " " + playCard.Color + "," + strconv.Itoa(playCard.Number))
}

func getBestCardAndFlagToPlayOn(flagOptions [9]board.Flag, playerCards []card.Card, direction string, boardInfo board.Board) (card.Card,int) {
    maxScore := 0
    playIndex := 0
    cardToPlay := playerCards[0]
    score := 0
    finalCardToPlay := cardToPlay
    for index := range flagOptions {
        if canPlay(flagOptions[index], direction) {
            if direction == "north" {
                score,cardToPlay = determineBestCardForThisFlag(flagOptions[index].North, playerCards, boardInfo, index)
            } else {
                score,cardToPlay = determineBestCardForThisFlag(flagOptions[index].South, playerCards, boardInfo, index)
            }
            if score > maxScore {
                maxScore = score
                finalCardToPlay = cardToPlay
                playIndex = index+1
            }
        }
    }
    if maxScore == 0 {
        maxBackupScore := 0
        for handIndex := range playerCards{
            if playerCards[handIndex].BestRankPlay > maxBackupScore {
                maxBackupScore = playerCards[handIndex].BestRankPlay
                playIndex = playerCards[handIndex].BestFlagIndex+1
                finalCardToPlay = playerCards[handIndex]
            }
        }
    }
    return finalCardToPlay,playIndex
}

func canPlay(flagAttempt board.Flag, direction string) bool {
    if direction == "north" {
        if len(flagAttempt.North) < 3 && flagAttempt.Claimer == "unclaimed" {
            return true
        }
    } else {
        if len(flagAttempt.South) < 3 && flagAttempt.Claimer == "unclaimed" {
            return true
        }   
    }
    return false
}

func determineBestCardForThisFlag(flagCardsMySide []card.Card, myHand []card.Card, myBoard board.Board, flagIndex int) (int, card.Card) {
    cardToPlay := card.Card{"color1", 0,0,0}
    index := -1

    bestFormation := getBestFormation(flagCardsMySide, myBoard)
    cardToPlay,index = getContinuationForWedge(myHand, flagCardsMySide)
    if cardToPlay.Number > 0 {
        myHand[index].BestRankPlay = 10
        myHand[index].BestFlagIndex = flagIndex
        return 10,cardToPlay
    }
    cardToPlay,index = getContinuationForPhalanx(myHand, flagCardsMySide)
    if (cardToPlay.Number > 0) && (myHand[index].BestRankPlay < 9) {
        if bestFormation != "wedge"{
            myHand[index].BestRankPlay = 9
            myHand[index].BestFlagIndex = flagIndex
            return 9,cardToPlay
        } else {
            myHand[index].BestRankPlay = 8
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    cardToPlay,index = getContinuationForBattalion(myHand, flagCardsMySide)
    if (cardToPlay.Number > 0) && (myHand[index].BestRankPlay < 7) {
        if (bestFormation != "phalanx") && (bestFormation != "wedge") {
            myHand[index].BestRankPlay = 7
            myHand[index].BestFlagIndex = flagIndex
            return 7,cardToPlay
        } else {
            myHand[index].BestRankPlay = 6
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    cardToPlay,index = getContinuationForSkirmish(myHand, flagCardsMySide)
    if (cardToPlay.Number > 0) && (myHand[index].BestRankPlay < 5) {
        if (bestFormation == "skirmish") || (bestFormation == "host") {
            myHand[index].BestRankPlay = 5
            myHand[index].BestFlagIndex = flagIndex
            return 5,cardToPlay
        } else {
            myHand[index].BestRankPlay = 4
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    cardToPlay,index = getHighestCardForHost(myHand)
    if (cardToPlay.Number > 0) && (myHand[index].BestRankPlay < 3) {
        if bestFormation == "host"{
            if len(flagCardsMySide) == 0 {
                myHand[index].BestRankPlay = 3
                myHand[index].BestFlagIndex = flagIndex
                return 3,cardToPlay
            }
            myHand[index].BestRankPlay = 1
            myHand[index].BestFlagIndex = flagIndex
            return 1,cardToPlay
        } else {
             if len(flagCardsMySide) == 0 {
                myHand[index].BestRankPlay = 3
                myHand[index].BestFlagIndex = flagIndex
            }
            myHand[index].BestRankPlay = 1
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    return 0,cardToPlay
}

func getContinuationForWedge(myHand []card.Card, flagCardsMySide []card.Card) (card.Card, int) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
    } else if len(flagCardsMySide) == 2 {
        if flagCardsMySide[0].Color != flagCardsMySide[1].Color {
            return continuationCard,index
        }
        firstMinusSecond := flagCardsMySide[0].Number-flagCardsMySide[1].Number
        secondMinusFirst := flagCardsMySide[1].Number-flagCardsMySide[0].Number
        //One higer or Two lower
        if (firstMinusSecond == 1) || (secondMinusFirst == 2) {
            continuationCard,index = checkHandForWedgeContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
        } else if (firstMinusSecond == 2) || (secondMinusFirst == 1) {
            //Two Higer or One Lower
            continuationCard,index = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[1])
        }
    }
    return continuationCard,index
}

func checkHandForWedgeContinuation(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card,int) {
    valueBelow := cardForValueBelow.Number - 1
    valueAbove := cardForValueAbove.Number + 1
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) && myHand[index].Color == cardForValueBelow.Color {
            return myHand[index],index
        }
    }
    return card.Card{"color1", 0,0,0},-1
}

func getContinuationForPhalanx(myHand []card.Card, flagCardsMySide []card.Card) (card.Card,int) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForPhalanxContinuation(myHand, flagCardsMySide[0].Number)
    } else if len(flagCardsMySide) == 2{
        if flagCardsMySide[0].Number == flagCardsMySide[1].Number {
            continuationCard,index = checkHandForPhalanxContinuation(myHand, flagCardsMySide[0].Number)
        }
    }
    return continuationCard,index
}

func checkHandForPhalanxContinuation( myHand[]card.Card, cardValueToMatch int) (card.Card,int) {
    for index := range myHand {
        if myHand[index].Number == cardValueToMatch {
            return myHand[index],index
        }
    }
    return card.Card{"color1", 0,0,0},-1
}

func getContinuationForBattalion(myHand []card.Card, flagCardsMySide []card.Card) (card.Card,int) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForBattalionContinuation(myHand, flagCardsMySide[0].Color)
    } else if len(flagCardsMySide) == 2 {
        if flagCardsMySide[0].Color == flagCardsMySide[1].Color{
            continuationCard,index = checkHandForBattalionContinuation(myHand, flagCardsMySide[0].Color)
        }
    }
    return continuationCard,index
}

func checkHandForBattalionContinuation( myHand []card.Card, colorToMatch string) (card.Card,int) {
    for index := range myHand {
        if myHand[index].Color == colorToMatch {
            return myHand[index],index
        }
    }
    return card.Card{"color1", 0,0,0},-1
}

func getContinuationForSkirmish(myHand []card.Card, flagCardsMySide []card.Card) (card.Card,int) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForSkirmishContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
    } else if len(flagCardsMySide) == 2 {
        firstMinusSecond := flagCardsMySide[0].Number-flagCardsMySide[1].Number
        secondMinusFirst := flagCardsMySide[1].Number-flagCardsMySide[0].Number
        //One higer or Two lower
        if (firstMinusSecond == 1) || (secondMinusFirst == 2) {
            continuationCard,index = checkHandForSkirmishContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
        } else if (firstMinusSecond == 2) || (secondMinusFirst == 1) {
            //Two Higer or One Lower
            continuationCard,index = checkHandForSkirmishContinuation(myHand, flagCardsMySide[0], flagCardsMySide[1])
        }
    }
    return continuationCard,index
}

func checkHandForSkirmishContinuation(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card,int) {
    valueBelow := cardForValueBelow.Number - 1
    valueAbove := cardForValueAbove.Number + 1
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) {
            return myHand[index],index
        }
    }
    return card.Card{"color1", 0,0,0},-1
}

func getHighestCardForHost(myHand []card.Card) (card.Card,int) {
    maxValueCard := card.Card{"color1", 0,0,0}
    returnIndex := -1
    for index := range myHand {
        if (myHand[index].Number > maxValueCard.Number) && (myHand[index].BestRankPlay < 4) {
            maxValueCard = myHand[index]
            returnIndex = index
        }
    }
    return maxValueCard,returnIndex
}

func getBestFormation(cardsMySide []card.Card, currentBoard board.Board) (string) {
    max_formation := "host"
    if len(cardsMySide) < 3 {
        number_of_cards_left := 3 - len(cardsMySide)
        myCardCombinations := utilities.CardCombinations(currentBoard.GetUnplayedCards(), number_of_cards_left)
        for index := range myCardCombinations {
            switch max_formation {
                case "wedge":
                    return max_formation
                case "phalanx":
                    max_formation = checkForHigherThanPhalanx(cardsMySide, myCardCombinations[index])
                case "battalion":
                    max_formation = checkForHigherThanBattalion(cardsMySide, myCardCombinations[index])
                case "skirmish":
                    max_formation = checkForHigherThanSkirmish(cardsMySide, myCardCombinations[index])
                case "host":
                    max_formation = checkForHigherThanHost(cardsMySide, myCardCombinations[index])
            }
        }
    }
    return max_formation
}

func checkForHigherThanPhalanx(fixedCardsMySide []card.Card, cardCombo []card.Card) (string) {
    cardToPlayWedge := card.Card{"color1", 0,0,0}
    if len(cardCombo) == 1 {
        cardToPlayWedge,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayWedge,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
    }
    if cardToPlayWedge.Number > 0 {
        return "wedge"
    }
    return "phalanx"
}

func checkForHigherThanBattalion(fixedCardsMySide []card.Card, cardCombo []card.Card) (string) {
    cardToPlayPhalanx := card.Card{"color1", 0,0,0}
    cardToPlayWedge := card.Card{"color1", 0,0,0}
    if len(cardCombo) == 1 {
        cardToPlayPhalanx,_ = getContinuationForPhalanx(cardCombo, fixedCardsMySide)
        cardToPlayWedge,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayPhalanx,_ = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
    }
    if cardToPlayWedge.Number > 0 {
        return "wedge"
    }
    if cardToPlayPhalanx.Number > 0 {
        return "phalanx"
    }
    return "battalion"
}

func checkForHigherThanSkirmish(fixedCardsMySide []card.Card, cardCombo []card.Card) (string) {
    cardToPlayBattalion := card.Card{"color1", 0,0,0}
    cardToPlayPhalanx := card.Card{"color1", 0,0,0}
    cardToPlayWedge := card.Card{"color1", 0,0,0}
    if len(cardCombo) == 1 {
        cardToPlayBattalion,_ = getContinuationForBattalion(cardCombo, fixedCardsMySide)
        cardToPlayPhalanx,_ = getContinuationForPhalanx(cardCombo, fixedCardsMySide)
        cardToPlayWedge,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayBattalion,_ = getContinuationForBattalion(fixedCardsMySide, cardCombo)
        cardToPlayPhalanx,_ = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
    }
    if cardToPlayWedge.Number > 0 {
        return "wedge"
    }
    if cardToPlayPhalanx.Number > 0 {
        return "phalanx"
    }
    if cardToPlayBattalion.Number > 0 {
        return "battalion"
    }
    return "skirmish"
}

func checkForHigherThanHost(fixedCardsMySide []card.Card, cardCombo []card.Card) (string) {
    cardToPlaySkirmish := card.Card{"color1", 0,0,0}
    cardToPlayBattalion := card.Card{"color1", 0,0,0}
    cardToPlayPhalanx := card.Card{"color1", 0,0,0}
    cardToPlayWedge := card.Card{"color1", 0,0,0}
    if len(cardCombo) == 1 {
        cardToPlaySkirmish,_ = getContinuationForSkirmish(cardCombo, fixedCardsMySide)
        cardToPlayBattalion,_ = getContinuationForBattalion(cardCombo, fixedCardsMySide)
        cardToPlayPhalanx,_ = getContinuationForPhalanx(cardCombo, fixedCardsMySide)
        cardToPlayWedge,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlaySkirmish,_ = getContinuationForSkirmish(fixedCardsMySide, cardCombo)
        cardToPlayBattalion,_ = getContinuationForBattalion(fixedCardsMySide, cardCombo)
        cardToPlayPhalanx,_ = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
    }
    if cardToPlayWedge.Number > 0 {
        return "wedge"
    }
    if cardToPlayPhalanx.Number > 0 {
        return "phalanx"
    }
    if cardToPlayBattalion.Number > 0 {
        return "battalion"
    }
    if cardToPlaySkirmish.Number > 0 {
        return "skirmish"
    }
    return "host"
}
