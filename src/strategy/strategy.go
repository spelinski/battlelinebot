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
                score,cardToPlay = determineBestCardForThisFlag(flagOptions[index].North, playerCards, boardInfo)
            } else {
                score,cardToPlay = determineBestCardForThisFlag(flagOptions[index].South, playerCards, boardInfo)
            }
            if score > maxScore {
                maxScore = score
                finalCardToPlay = cardToPlay
                playIndex = index+1
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

func determineBestCardForThisFlag(flagCardsMySide []card.Card, myHand []card.Card, myBoard board.Board) (int, card.Card) {
    cardToPlay := card.Card{"color1", 0}

    bestFormation := getBestFormation(flagCardsMySide, myBoard)
    switch bestFormation{
        case "wedge":
            cardToPlay = getContinuationForWedge(myHand, flagCardsMySide)
            if cardToPlay.Number > 0 {
                return 10,cardToPlay
            }

        case "phalanx":
            cardToPlay = getContinuationForPhalanx(myHand, flagCardsMySide)
            if cardToPlay.Number > 0 {
                return 9,cardToPlay
            }

        case "battalion":
            cardToPlay = getContinuationForBattalion(myHand, flagCardsMySide)
            if cardToPlay.Number > 0 {
                return 8,cardToPlay
            }

        case "skirmish":
            cardToPlay = getContinuationForSkirmish(myHand, flagCardsMySide)
            if cardToPlay.Number > 0 {
                return 7,cardToPlay
            }

        case "host":
            cardToPlay = getHighestCardForHost(myHand)
            if cardToPlay.Number > 0 {
                if len(flagCardsMySide) == 0 {
                    return 6,cardToPlay
                }
                return 1,cardToPlay
            }
    }
    return 0,cardToPlay
}

func getContinuationForWedge(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    if len(flagCardsMySide) == 1{
        continuationCard = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
    } else if len(flagCardsMySide) == 2 {
        if flagCardsMySide[0].Color != flagCardsMySide[1].Color {
            return continuationCard
        }
        firstMinusSecond := flagCardsMySide[0].Number-flagCardsMySide[1].Number
        secondMinusFirst := flagCardsMySide[1].Number-flagCardsMySide[0].Number
        //One higer or Two lower
        if (firstMinusSecond == 1) || (secondMinusFirst == 2) {
            continuationCard = checkHandForWedgeContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
        } else if (firstMinusSecond == 2) || (secondMinusFirst == 1) {
            //Two Higer or One Lower
            continuationCard = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[1])
        }
    }
    return continuationCard
}

func checkHandForWedgeContinuation(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card) {
    valueBelow := cardForValueBelow.Number - 1
    valueAbove := cardForValueAbove.Number + 1
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) && myHand[index].Color == cardForValueBelow.Color {
            return myHand[index]
        }
    }
    return card.Card{"color1", 0}
}

func getContinuationForPhalanx(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    if len(flagCardsMySide) == 1{
        continuationCard = checkHandForPhalanxContinuation(myHand, flagCardsMySide[0].Number)
    } else if len(flagCardsMySide) == 2{
        if flagCardsMySide[0].Number == flagCardsMySide[1].Number {
            continuationCard = checkHandForPhalanxContinuation(myHand, flagCardsMySide[0].Number)
        }
    }
    return continuationCard
}

func checkHandForPhalanxContinuation( myHand[]card.Card, cardValueToMatch int) (card.Card) {
    for index := range myHand {
        if myHand[index].Number == cardValueToMatch {
            return myHand[index]
        }
    }
    return card.Card{"color1", 0}
}

func getContinuationForBattalion(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    if len(flagCardsMySide) == 1{
        continuationCard = checkHandForBattalionContinuation(myHand, flagCardsMySide[0].Color)
    } else if len(flagCardsMySide) == 2 {
        if flagCardsMySide[0].Color == flagCardsMySide[1].Color{
            continuationCard = checkHandForBattalionContinuation(myHand, flagCardsMySide[0].Color)
        }
    }
    return continuationCard
}

func checkHandForBattalionContinuation( myHand []card.Card, colorToMatch string) (card.Card) {
    for index := range myHand {
        if myHand[index].Color == colorToMatch {
            return myHand[index]
        }
    }
    return card.Card{"color1", 0}
}

func getContinuationForSkirmish(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    if len(flagCardsMySide) == 1{
        continuationCard = checkHandForSkirmishContinuation(myHand, flagCardsMySide[0], flagCardsMySide[0])
    } else if len(flagCardsMySide) == 2 {
        firstMinusSecond := flagCardsMySide[0].Number-flagCardsMySide[1].Number
        secondMinusFirst := flagCardsMySide[1].Number-flagCardsMySide[0].Number
        //One higer or Two lower
        if (firstMinusSecond == 1) || (secondMinusFirst == 2) {
            continuationCard = checkHandForSkirmishContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
        } else if (firstMinusSecond == 2) || (secondMinusFirst == 1) {
            //Two Higer or One Lower
            continuationCard = checkHandForSkirmishContinuation(myHand, flagCardsMySide[0], flagCardsMySide[1])
        }
    }
    return continuationCard
}

func checkHandForSkirmishContinuation(myHand []card.Card, cardForValueBelow card.Card, cardForValueAbove card.Card) (card.Card) {
    valueBelow := cardForValueBelow.Number - 1
    valueAbove := cardForValueAbove.Number + 1
    for index := range myHand {
        if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) {
            return myHand[index]
        }
    }
    return card.Card{"color1", 0}
}

func getHighestCardForHost(myHand []card.Card) (card.Card) {
    maxValueCard := card.Card{"color1", 0}
    for index := range myHand {
        if myHand[index].Number > maxValueCard.Number {
            maxValueCard = myHand[index]
        }
    }
    return maxValueCard
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
    cardToPlayWedge := card.Card{"color1", 0}
    if len(cardCombo) == 1 {
        cardToPlayWedge = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayWedge = getContinuationForWedge(fixedCardsMySide, cardCombo)
    }
    if cardToPlayWedge.Number > 0 {
        return "wedge"
    }
    return "phalanx"
}

func checkForHigherThanBattalion(fixedCardsMySide []card.Card, cardCombo []card.Card) (string) {
    cardToPlayPhalanx := card.Card{"color1", 0}
    cardToPlayWedge := card.Card{"color1", 0}
    if len(cardCombo) == 1 {
        cardToPlayPhalanx = getContinuationForPhalanx(cardCombo, fixedCardsMySide)
        cardToPlayWedge = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayPhalanx = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge = getContinuationForWedge(fixedCardsMySide, cardCombo)
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
    cardToPlayBattalion := card.Card{"color1", 0}
    cardToPlayPhalanx := card.Card{"color1", 0}
    cardToPlayWedge := card.Card{"color1", 0}
    if len(cardCombo) == 1 {
        cardToPlayBattalion = getContinuationForBattalion(cardCombo, fixedCardsMySide)
        cardToPlayPhalanx = getContinuationForPhalanx(cardCombo, fixedCardsMySide)
        cardToPlayWedge = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlayBattalion = getContinuationForBattalion(fixedCardsMySide, cardCombo)
        cardToPlayPhalanx = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge = getContinuationForWedge(fixedCardsMySide, cardCombo)
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
    cardToPlaySkirmish := card.Card{"color1", 0}
    cardToPlayBattalion := card.Card{"color1", 0}
    cardToPlayPhalanx := card.Card{"color1", 0}
    cardToPlayWedge := card.Card{"color1", 0}
    if len(cardCombo) == 1 {
        cardToPlaySkirmish = getContinuationForSkirmish(cardCombo, fixedCardsMySide)
        cardToPlayBattalion = getContinuationForBattalion(cardCombo, fixedCardsMySide)
        cardToPlayPhalanx = getContinuationForPhalanx(cardCombo, fixedCardsMySide)
        cardToPlayWedge = getContinuationForWedge(cardCombo, fixedCardsMySide)
    } else {
        cardToPlaySkirmish = getContinuationForSkirmish(fixedCardsMySide, cardCombo)
        cardToPlayBattalion = getContinuationForBattalion(fixedCardsMySide, cardCombo)
        cardToPlayPhalanx = getContinuationForPhalanx(fixedCardsMySide, cardCombo)
        cardToPlayWedge = getContinuationForWedge(fixedCardsMySide, cardCombo)
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
