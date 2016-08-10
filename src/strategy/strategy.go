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
    finalCardToPlay := cardToPlay
    //This assumes 9 flags
    arrayOfIndexes := [9]int{4,3,5,2,6,1,7,0,8}
    for arrayIndex := range arrayOfIndexes {
        index := arrayOfIndexes[arrayIndex]
        if canPlay(flagOptions[index], direction) {
            determineBestCardForThisFlag(flagOptions, playerCards, boardInfo, index, direction)
        }
    }
    for handIndex := range playerCards{
        if playerCards[handIndex].BestRankPlay < 9000 {
            if playerCards[handIndex].BestRankPlay > maxScore {
                maxScore = playerCards[handIndex].BestRankPlay
                playIndex = playerCards[handIndex].BestFlagIndex+1
                finalCardToPlay = playerCards[handIndex]
            } else if (playerCards[handIndex].BestRankPlay == maxScore) && (playerCards[handIndex].Number > finalCardToPlay.Number) {
                maxScore = playerCards[handIndex].BestRankPlay
                playIndex = playerCards[handIndex].BestFlagIndex+1
                finalCardToPlay = playerCards[handIndex]
            }
        }
    }
    for handIndex := range playerCards{
        playerCards[handIndex].BestRankPlay = 0
        playerCards[handIndex].BestFlagIndex = 0
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

func determineBestCardForThisFlag(flagOptions [9] board.Flag, myHand []card.Card, myBoard board.Board, flagIndex int, direction string) {
    cardToPlay := card.Card{"color1", 0,0,0}
    index := -1
    twoAway := false
    flagCardsMySide := []card.Card{}
    flagCardsTheirSide := []card.Card{}

    if direction == "north" {
        flagCardsMySide = flagOptions[flagIndex].North
        flagCardsTheirSide = flagOptions[flagIndex].South
    } else {
        flagCardsMySide = flagOptions[flagIndex].South
        flagCardsTheirSide = flagOptions[flagIndex].North
    }

    bestFormation := getBestFormation(flagCardsMySide, myBoard)
    enemyBestFormation := getBestEnemyFormation(flagCardsTheirSide, myBoard)
    if (bestFormation == "wedge") {
        cardToPlay,index,twoAway = getContinuationForWedge(myHand, flagCardsMySide)
        if (cardToPlay.Number > 0) && ((cardToPlay.BestRankPlay < 320) || (cardToPlay.BestRankPlay > 9000)) {
            if !twoAway {
                if(isFirstFormationBetter("wedge",enemyBestFormation)) {
                    myHand[index].BestRankPlay = 320
                    myHand[index].BestFlagIndex = flagIndex
                } else if cardToPlay.BestRankPlay < 32 || (cardToPlay.BestRankPlay > 9000) {
                    myHand[index].BestRankPlay = 32
                    myHand[index].BestFlagIndex = flagIndex
                }
            } else {
                myHand[index].BestRankPlay = 9001
                myHand[index].BestFlagIndex = flagIndex
            }
        }
    }

    if (bestFormation == "wedge") || (bestFormation == "phalanx") {
        cardToPlay,index = getContinuationForPhalanx(myHand, flagCardsMySide)
        if (cardToPlay.Number > 0) && ((cardToPlay.BestRankPlay < 310) || ((cardToPlay.BestRankPlay < 29) && (cardToPlay.BestRankPlay > 8))) {
            if bestFormation != "wedge"{
                if(isFirstFormationBetter("phalanx",enemyBestFormation)) {
                    myHand[index].BestRankPlay = 310
                    myHand[index].BestFlagIndex = flagIndex
                } else if cardToPlay.BestRankPlay < 31 {
                    myHand[index].BestRankPlay = 31
                    myHand[index].BestFlagIndex = flagIndex
                }
            } else if cardToPlay.BestRankPlay < 8 || ((cardToPlay.BestRankPlay < 29) && (cardToPlay.BestRankPlay > 8)){
                myHand[index].BestRankPlay = 8
                myHand[index].BestFlagIndex = flagIndex
            }
        }
    }

    if (bestFormation == "wedge") || (bestFormation == "phalanx") || (bestFormation == "battalion") {
        cardToPlay,index = getContinuationForBattalion(myHand, flagCardsMySide)
        if (cardToPlay.Number > 0) && ((cardToPlay.BestRankPlay < 300) || ((cardToPlay.BestRankPlay < 29) && (cardToPlay.BestRankPlay > 8))) {
            if (bestFormation != "phalanx") && (bestFormation != "wedge") {
                if(isFirstFormationBetter("battalion",enemyBestFormation)) {
                    myHand[index].BestRankPlay = 300
                    myHand[index].BestFlagIndex = flagIndex
                } else if cardToPlay.BestRankPlay < 30 {
                    myHand[index].BestRankPlay = 30
                    myHand[index].BestFlagIndex = flagIndex
                }
            } else if cardToPlay.BestRankPlay < 6 || ((cardToPlay.BestRankPlay < 29) && (cardToPlay.BestRankPlay > 8)){
                myHand[index].BestRankPlay = 6
                myHand[index].BestFlagIndex = flagIndex
            }
        }
    }

    if (bestFormation == "wedge") || (bestFormation == "phalanx") || (bestFormation == "battalion") || (bestFormation == "skirmish") {
        cardToPlay,index = getContinuationForSkirmish(myHand, flagCardsMySide)
        if (cardToPlay.Number > 0) && ((cardToPlay.BestRankPlay < 290) || ((cardToPlay.BestRankPlay < 29) && (cardToPlay.BestRankPlay > 8))) {
            if (bestFormation == "skirmish") || (bestFormation == "host") {
                if(isFirstFormationBetter("skirmish",enemyBestFormation)) {
                    myHand[index].BestRankPlay = 290
                    myHand[index].BestFlagIndex = flagIndex
                } else if cardToPlay.BestRankPlay < 29 {
                    myHand[index].BestRankPlay = 29
                    myHand[index].BestFlagIndex = flagIndex
                }
            } else if cardToPlay.BestRankPlay < 4 || ((cardToPlay.BestRankPlay < 29) && (cardToPlay.BestRankPlay > 8)){
                myHand[index].BestRankPlay = 4
                myHand[index].BestFlagIndex = flagIndex
            }
        }
    }


    tempHand := []card.Card{}
    tempHand = utilities.CopySliceToSlice(tempHand, myHand)
    for i:=0; i < len(myHand); i++ {
        cardToPlay,index = getHighestCardForHost(tempHand)
        if (cardToPlay.Number > 0) && ((cardToPlay.BestRankPlay < 4) || ((cardToPlay.BestRankPlay < 28) && (cardToPlay.BestRankPlay > 8)))  {
            indexForRealHand := utilities.FindElementInSlice(myHand, cardToPlay)
            if bestFormation == "host"{
                if len(flagCardsMySide) == 0 {
                    bestFormationForThisCard := getBestFormation([]card.Card{cardToPlay}, myBoard)
                    if bestFormationForThisCard == "wedge" {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 28){
                            myHand[indexForRealHand].BestRankPlay = 28
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 27) {
                            myHand[indexForRealHand].BestRankPlay = 27
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 26) {
                            myHand[indexForRealHand].BestRankPlay = 26
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 25) {
                            myHand[indexForRealHand].BestRankPlay = 25
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 24) {
                            myHand[indexForRealHand].BestRankPlay = 24
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (bestFormationForThisCard == "phalanx") &&  (cardToPlay.BestRankPlay < 23) {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 23){
                            myHand[indexForRealHand].BestRankPlay = 23
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 22) {
                            myHand[indexForRealHand].BestRankPlay = 22
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 21) {
                            myHand[indexForRealHand].BestRankPlay = 21
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 20) {
                            myHand[indexForRealHand].BestRankPlay = 20
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 19) {
                            myHand[indexForRealHand].BestRankPlay = 19
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (bestFormationForThisCard == "battalion") &&  (cardToPlay.BestRankPlay < 18) {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 18){
                            myHand[indexForRealHand].BestRankPlay = 18
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 17) {
                            myHand[indexForRealHand].BestRankPlay = 17
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 16) {
                            myHand[indexForRealHand].BestRankPlay = 16
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 15) {
                            myHand[indexForRealHand].BestRankPlay = 15
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 14) {
                            myHand[indexForRealHand].BestRankPlay = 14
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (bestFormationForThisCard == "skirmish") &&  (cardToPlay.BestRankPlay < 13) {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 13){
                            myHand[indexForRealHand].BestRankPlay = 13
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 12) {
                            myHand[indexForRealHand].BestRankPlay = 12
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 11) {
                            myHand[indexForRealHand].BestRankPlay = 11
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 10) {
                            myHand[indexForRealHand].BestRankPlay = 10
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 9) {
                            myHand[indexForRealHand].BestRankPlay = 9
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (cardToPlay.BestRankPlay < 2){
                        myHand[indexForRealHand].BestRankPlay = 2
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    }

                } else if (cardToPlay.BestRankPlay < 1){
                    myHand[indexForRealHand].BestRankPlay = 1
                    myHand[indexForRealHand].BestFlagIndex = flagIndex
                }
            } else {
                if len(flagCardsMySide) == 0 {
                    bestFormationForThisCard := getBestFormation([]card.Card{cardToPlay}, myBoard)
                    if bestFormationForThisCard == "wedge" {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 28){
                            myHand[indexForRealHand].BestRankPlay = 28
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 27) {
                            myHand[indexForRealHand].BestRankPlay = 27
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 26) {
                            myHand[indexForRealHand].BestRankPlay = 26
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 25) {
                            myHand[indexForRealHand].BestRankPlay = 25
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 24) {
                            myHand[indexForRealHand].BestRankPlay = 24
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (bestFormationForThisCard == "phalanx") &&  (cardToPlay.BestRankPlay < 23) {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 23){
                            myHand[indexForRealHand].BestRankPlay = 23
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 22) {
                            myHand[indexForRealHand].BestRankPlay = 22
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 21) {
                            myHand[indexForRealHand].BestRankPlay = 21
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 20) {
                            myHand[indexForRealHand].BestRankPlay = 20
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 19) {
                            myHand[indexForRealHand].BestRankPlay = 19
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (bestFormationForThisCard == "battalion") &&  (cardToPlay.BestRankPlay < 18) {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 18){
                            myHand[indexForRealHand].BestRankPlay = 18
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 17) {
                            myHand[indexForRealHand].BestRankPlay = 17
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 16) {
                            myHand[indexForRealHand].BestRankPlay = 16
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 15) {
                            myHand[indexForRealHand].BestRankPlay = 15
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 14) {
                            myHand[indexForRealHand].BestRankPlay = 14
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (bestFormationForThisCard == "skirmish") &&  (cardToPlay.BestRankPlay < 13) {
                        if (flagIndex == 4) && (myHand[indexForRealHand].BestRankPlay < 13){
                            myHand[indexForRealHand].BestRankPlay = 13
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 3 || flagIndex == 5) && (myHand[indexForRealHand].BestRankPlay < 12) {
                            myHand[indexForRealHand].BestRankPlay = 12
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 2 || flagIndex == 6) && (myHand[indexForRealHand].BestRankPlay < 11) {
                            myHand[indexForRealHand].BestRankPlay = 11
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 1 || flagIndex == 7) && (myHand[indexForRealHand].BestRankPlay < 10) {
                            myHand[indexForRealHand].BestRankPlay = 10
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        } else if (flagIndex == 0 || flagIndex == 8) && (myHand[indexForRealHand].BestRankPlay < 9) {
                            myHand[indexForRealHand].BestRankPlay = 9
                            myHand[indexForRealHand].BestFlagIndex = flagIndex
                        }
                    } else if (cardToPlay.BestRankPlay < 2){
                        myHand[indexForRealHand].BestRankPlay = 2
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    }
                } else if (cardToPlay.BestRankPlay < 1){
                    myHand[indexForRealHand].BestRankPlay = 1
                    myHand[indexForRealHand].BestFlagIndex = flagIndex
                }
            }
        } 
        tempHand = append(tempHand[:index], tempHand[index+1:]...)
    }
}

func getContinuationForWedge(myHand []card.Card, flagCardsMySide []card.Card) (card.Card, int,bool) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    twoAway := false
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
        if continuationCard.Number == 0 {
            continuationCard,index = checkHandForWedgeContinuationTwoAway(myHand, flagCardsMySide[0], flagCardsMySide[0])
            if continuationCard.Number > 0{
                twoAway = true
            }
        }
    } else if len(flagCardsMySide) == 2 {
        if flagCardsMySide[0].Color != flagCardsMySide[1].Color {
            return continuationCard,index,twoAway
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
    return continuationCard,index,twoAway
}

func checkHandForWedgeContinuation(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card,int) {
    valueBelow := cardForValueBelow.Number - 1
    valueAbove := cardForValueAbove.Number + 1
    cardToReturn := card.Card{"color1",0,0,0}
    indexToRetun := -1
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) && myHand[index].Color == cardForValueBelow.Color {
            if myHand[index].Number > cardToReturn.Number {
                cardToReturn = myHand[index]
                indexToRetun = index
            }
        }
    }
    return cardToReturn,indexToRetun
}

func checkHandForWedgeContinuationTwoAway(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card,int) {
    valueBelow := cardForValueBelow.Number - 2
    valueAbove := cardForValueAbove.Number + 2
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
    cardToReturn := card.Card{"color1",0,0,0}
    indexToRetun := -1
    for index := range myHand {
        if myHand[index].Color == colorToMatch {
            if myHand[index].Number > cardToReturn.Number {
                cardToReturn = myHand[index]
                indexToRetun = index
            }
        }
    }
    return cardToReturn,indexToRetun
}

func getContinuationForSkirmish(myHand []card.Card, flagCardsMySide []card.Card) (card.Card,int) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForSkirmishContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
        if continuationCard.Number == 0 {
            continuationCard,index = checkHandForSkirmishContinuationTwoAway(myHand, flagCardsMySide[0], flagCardsMySide[0])
        }
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
    cardToReturn := card.Card{"color1",0,0,0}
    indexToRetun := -1
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) {
            if myHand[index].Number > cardToReturn.Number {
                cardToReturn = myHand[index]
                indexToRetun = index
            }
        }
    }
    return cardToReturn,indexToRetun
}

func checkHandForSkirmishContinuationTwoAway(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card,int) {
    valueBelow := cardForValueBelow.Number - 2
    valueAbove := cardForValueAbove.Number + 2
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
        if (myHand[index].Number > maxValueCard.Number)  {
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

func getBestEnemyFormation(cardsTheirSide []card.Card, currentBoard board.Board) (string) {
    max_formation := "host"
    if len(cardsTheirSide) == 0 {
        number_of_cards_left := 3
        myCardCombinations := utilities.CardCombinations(currentBoard.GetEnemyAvailCards(), number_of_cards_left)
        for index := range myCardCombinations {
            movedCardToTheirSide :=  []card.Card{myCardCombinations[index][0]}
            myCardCombinations[index] = append(myCardCombinations[index][:0],myCardCombinations[index][1:]...)
            switch max_formation {
                case "wedge":
                    return max_formation
                case "phalanx":
                    max_formation = checkForHigherThanPhalanx(movedCardToTheirSide, myCardCombinations[index])
                case "battalion":
                    max_formation = checkForHigherThanBattalion(movedCardToTheirSide, myCardCombinations[index])
                case "skirmish":
                    max_formation = checkForHigherThanSkirmish(movedCardToTheirSide, myCardCombinations[index])
                case "host":
                    max_formation = checkForHigherThanHost(movedCardToTheirSide, myCardCombinations[index])
            }
        }
    } else if len(cardsTheirSide) < 3 {
        number_of_cards_left := 3 - len(cardsTheirSide)
        myCardCombinations := utilities.CardCombinations(currentBoard.GetEnemyAvailCards(), number_of_cards_left)
        for index := range myCardCombinations {
            switch max_formation {
                case "wedge":
                    return max_formation
                case "phalanx":
                    max_formation = checkForHigherThanPhalanx(cardsTheirSide, myCardCombinations[index])
                case "battalion":
                    max_formation = checkForHigherThanBattalion(cardsTheirSide, myCardCombinations[index])
                case "skirmish":
                    max_formation = checkForHigherThanSkirmish(cardsTheirSide, myCardCombinations[index])
                case "host":
                    max_formation = checkForHigherThanHost(cardsTheirSide, myCardCombinations[index])
            }
        }
    } else if len(cardsTheirSide) == 3 {
        movedCardToTheirHand :=  []card.Card{cardsTheirSide[0]}
        cardsTheirSide = append(cardsTheirSide[:0],cardsTheirSide[1:]...)
        switch max_formation {
            case "wedge":
                return max_formation
            case "phalanx":
                max_formation = checkForHigherThanPhalanx(cardsTheirSide, movedCardToTheirHand)
            case "battalion":
                max_formation = checkForHigherThanBattalion(cardsTheirSide, movedCardToTheirHand)
            case "skirmish":
                max_formation = checkForHigherThanSkirmish(cardsTheirSide, movedCardToTheirHand)
            case "host":
                max_formation = checkForHigherThanHost(cardsTheirSide, movedCardToTheirHand)
        }
    }
    return max_formation
}

func checkForHigherThanPhalanx(fixedCardsMySide []card.Card, cardCombo []card.Card) (string) {
    cardToPlayWedge := card.Card{"color1", 0,0,0}
    if len(cardCombo) == 1 {
        cardToPlayWedge,_,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayWedge,_,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
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
        cardToPlayWedge,_,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayPhalanx,_ = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge,_,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
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
        cardToPlayWedge,_,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayBattalion,_ = getContinuationForBattalion(fixedCardsMySide, cardCombo)
        cardToPlayPhalanx,_ = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge,_,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
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
        cardToPlayWedge,_,_ = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlaySkirmish,_ = getContinuationForSkirmish(fixedCardsMySide, cardCombo)
        cardToPlayBattalion,_ = getContinuationForBattalion(fixedCardsMySide, cardCombo)
        cardToPlayPhalanx,_ = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge,_,_ = getContinuationForWedge(fixedCardsMySide, cardCombo)
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

func isFirstFormationBetter(firstFormation string, secondFormation string) (bool) {
    if firstFormation == "wedge" {
        return true
    } else if firstFormation == "phalanx" {
        if (secondFormation != "wedge") {
            return true
        }
    } else if firstFormation == "battalion" {
        if (secondFormation == "skirmish") || (secondFormation == "host") || (secondFormation == "battalion") {
            return true
        }
    } else if firstFormation == "skirmish" {
        if (secondFormation == "host") || (secondFormation == "skirmish") {
            return true
        }
    }

    return false
}
