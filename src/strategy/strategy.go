package strategy

import (
    "board"
    "player"
    "fmt"
    "card"
    "strconv"
)

func HandleGoPlayCommand(playerInfo player.Player, boardInfo board.Board) {
    playCard,flagToPlay := getBestCardAndFlagToPlayOn(boardInfo.Flags,playerInfo.Hand,playerInfo.Direction)
    fmt.Println("play " + strconv.Itoa(flagToPlay) + " " + playCard.Color + "," + strconv.Itoa(playCard.Number))
}

func getBestCardAndFlagToPlayOn(flagOptions [9]board.Flag, playerCards []card.Card, direction string) (card.Card,int) {
    maxScore := 0
    playIndex := 0
    cardToPlay := playerCards[0]
    score := 0
    finalCardToPlay := cardToPlay
    for index := range flagOptions {
        if canPlay(flagOptions[index], direction) {
            if direction == "north" {
                score,cardToPlay = determineBestCardForThisFlag(flagOptions[index].North, playerCards)
            } else {
                score,cardToPlay = determineBestCardForThisFlag(flagOptions[index].South, playerCards)
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

func determineBestCardForThisFlag(flagCardsMySide []card.Card, myHand []card.Card) (int, card.Card) {
    cardToPlay := card.Card{"color1", 0}
    cardToPlay = getContinuationForWedge(myHand, flagCardsMySide)
    if cardToPlay.Number > 0 {
        return 10,cardToPlay
    }
    
    cardToPlay = getContinuationForPhalanx(myHand, flagCardsMySide)
    if cardToPlay.Number > 0 {
        return 9,cardToPlay
    }

    cardToPlay = getContinuationForBattalion(myHand, flagCardsMySide)
    if cardToPlay.Number > 0 {
        return 8,cardToPlay
    }

    cardToPlay = getContinuationForSkirmish(myHand, flagCardsMySide)
    if cardToPlay.Number > 0 {
        return 7,cardToPlay
    }

    cardToPlay = getHighestCardForHost(myHand)
    if cardToPlay.Number > 0 {
        if len(flagCardsMySide) == 0 {
            return 2,cardToPlay
        }
        return 1,cardToPlay
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
        //first card is one higher than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number+1) {
            continuationCard = checkHandForWedgeContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
        }
        //first card is one lower than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number-1) {
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
        //first card is one higher than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number+1) {
            continuationCard = checkHandForSkirmishContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
        }
        //first card is one lower than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number-1) {
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
