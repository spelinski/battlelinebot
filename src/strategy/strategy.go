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
        return 1,cardToPlay
    }
    return 0,cardToPlay
}

func getContinuationForWedge(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    for index := range myHand {
        for indexSecond := range flagCardsMySide {
            valueBelow := flagCardsMySide[indexSecond].Number - 1
            valueAbove := flagCardsMySide[indexSecond].Number + 1
            if ((myHand[index].Number == valueBelow) || (myHand[index].Number == valueAbove)) && myHand[index].Color == flagCardsMySide[indexSecond].Color {
                return myHand[index]
            }
        }
    }
    return continuationCard
}

func getContinuationForPhalanx(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    for index := range myHand {
        for indexSecond := range flagCardsMySide {
            if myHand[index].Number == flagCardsMySide[indexSecond].Number {
                return myHand[index]
            }
        }
    }
    return continuationCard
}

func getContinuationForBattalion(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    for index := range myHand {
        for indexSecond := range flagCardsMySide {
            if myHand[index].Color == flagCardsMySide[indexSecond].Color {
                return myHand[index]
            }
        }
    }
    return continuationCard
}

func getContinuationForSkirmish(myHand []card.Card, flagCardsMySide []card.Card) (card.Card) {
    continuationCard := card.Card{"color1", 0}
    for index := range myHand {
        for indexSecond := range flagCardsMySide {
            valueBelow := flagCardsMySide[indexSecond].Number - 1
            valueAbove := flagCardsMySide[indexSecond].Number + 1
            if myHand[index].Number == valueBelow || myHand[index].Number == valueAbove {
                return myHand[index]
            }
        }
    }
    return continuationCard
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