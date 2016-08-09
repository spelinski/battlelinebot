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
    for index := range flagOptions {
        if canPlay(flagOptions[index], direction) {
            if direction == "north" {
                determineBestCardForThisFlag(flagOptions[index].North, playerCards, boardInfo, index)
            } else {
                determineBestCardForThisFlag(flagOptions[index].South, playerCards, boardInfo, index)
            }
        }
    }
    for handIndex := range playerCards{
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

func determineBestCardForThisFlag(flagCardsMySide []card.Card, myHand []card.Card, myBoard board.Board, flagIndex int) {
    cardToPlay := card.Card{"color1", 0,0,0}
    index := -1

    bestFormation := getBestFormation(flagCardsMySide, myBoard)
    cardToPlay,index = getContinuationForWedge(myHand, flagCardsMySide)
    if cardToPlay.Number > 0 {
        myHand[index].BestRankPlay = 20
        myHand[index].BestFlagIndex = flagIndex
    }
    cardToPlay,index = getContinuationForPhalanx(myHand, flagCardsMySide)
    if (cardToPlay.Number > 0) && (myHand[index].BestRankPlay < 19) {
        if bestFormation != "wedge"{
            myHand[index].BestRankPlay = 19
            myHand[index].BestFlagIndex = flagIndex
        } else {
            myHand[index].BestRankPlay = 8
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    cardToPlay,index = getContinuationForBattalion(myHand, flagCardsMySide)
    if (cardToPlay.Number > 0) && ((myHand[index].BestRankPlay < 8) || ((myHand[index].BestRankPlay < 17) && (myHand[index].BestRankPlay > 12))) {
        if (bestFormation != "phalanx") && (bestFormation != "wedge") {
            myHand[index].BestRankPlay = 18
            myHand[index].BestFlagIndex = flagIndex
        } else {
            myHand[index].BestRankPlay = 6
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    cardToPlay,index = getContinuationForSkirmish(myHand, flagCardsMySide)
    if (cardToPlay.Number > 0) && ((myHand[index].BestRankPlay < 6) || ((myHand[index].BestRankPlay < 17) && (myHand[index].BestRankPlay > 12))) {
        if (bestFormation == "skirmish") || (bestFormation == "host") {
            myHand[index].BestRankPlay = 17
            myHand[index].BestFlagIndex = flagIndex
        } else {
            myHand[index].BestRankPlay = 4
            myHand[index].BestFlagIndex = flagIndex
        }
    }
    tempHand := []card.Card{}
    tempHand = utilities.CopySliceToSlice(tempHand, myHand)
    for i:=0; i < len(myHand); i++ {
        cardToPlay,index = getHighestCardForHost(tempHand)
        if (cardToPlay.Number > 0) && ((myHand[index].BestRankPlay < 4) || ((myHand[index].BestRankPlay < 16) && (myHand[index].BestRankPlay > 12)))  {
            indexForRealHand := utilities.FindElementInSlice(myHand, cardToPlay)
            if bestFormation == "host"{
                if len(flagCardsMySide) == 0 {
                    if getBestFormation([]card.Card{cardToPlay}, myBoard) == "wedge" {
                        myHand[indexForRealHand].BestRankPlay = 16
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (getBestFormation([]card.Card{cardToPlay}, myBoard) == "phalanx") &&  (cardToPlay.BestRankPlay < 15) {
                        myHand[indexForRealHand].BestRankPlay = 15
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (getBestFormation([]card.Card{cardToPlay}, myBoard) == "battalion") &&  (cardToPlay.BestRankPlay < 14) {
                        myHand[indexForRealHand].BestRankPlay = 14
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (getBestFormation([]card.Card{cardToPlay}, myBoard) == "skirmish") &&  (cardToPlay.BestRankPlay < 13) {
                        myHand[indexForRealHand].BestRankPlay = 13
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
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
                    if getBestFormation([]card.Card{cardToPlay}, myBoard) == "wedge" {
                        myHand[indexForRealHand].BestRankPlay = 16
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (getBestFormation([]card.Card{cardToPlay}, myBoard) == "phalanx") &&  (cardToPlay.BestRankPlay < 15) {
                        myHand[indexForRealHand].BestRankPlay = 15
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (getBestFormation([]card.Card{cardToPlay}, myBoard) == "battalion") &&  (cardToPlay.BestRankPlay < 14) {
                        myHand[indexForRealHand].BestRankPlay = 14
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (getBestFormation([]card.Card{cardToPlay}, myBoard) == "skirmish") &&  (cardToPlay.BestRankPlay < 13) {
                        myHand[indexForRealHand].BestRankPlay = 13
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    } else if (cardToPlay.BestRankPlay < 2){
                        myHand[indexForRealHand].BestRankPlay = 2
                        myHand[indexForRealHand].BestFlagIndex = flagIndex
                    }
                } else if (cardToPlay.BestRankPlay < 1){
                    myHand[indexForRealHand].BestRankPlay = 1
                    myHand[indexForRealHand].BestFlagIndex = flagIndex
                }
            }
            tempHand = append(tempHand[:index], tempHand[index+1:]...)
        }
    }
}

func getContinuationForWedge(myHand []card.Card, flagCardsMySide []card.Card) (card.Card, int) {
    continuationCard := card.Card{"color1", 0,0,0}
    index := -1
    if len(flagCardsMySide) == 1{
        continuationCard,index = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
        if continuationCard.Number == 0 {
            continuationCard,index = checkHandForWedgeContinuationTwoAway(myHand, flagCardsMySide[0], flagCardsMySide[0])
        }
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
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) {
            return myHand[index],index
        }
    }
    return card.Card{"color1", 0,0,0},-1
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
        if (myHand[index].Number > maxValueCard.Number) && ((myHand[index].BestRankPlay < 4) || ((myHand[index].BestRankPlay < 17) && (myHand[index].BestRankPlay > 12)))  {
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
