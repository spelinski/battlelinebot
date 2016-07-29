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
            return 6,cardToPlay
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

        //first card is two higher than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number+2) {
            continuationCard = checkHandForWedgeContinuation(myHand, flagCardsMySide[0], flagCardsMySide[1])
        }

        //first card is two higher than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number-2) {
            continuationCard = checkHandForWedgeContinuation(myHand, flagCardsMySide[1], flagCardsMySide[0])
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

        //first card is two higher than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number+2) {
            continuationCard = checkHandForSkirmishContinuation(myHand, flagCardsMySide[0], flagCardsMySide[1])
        }

        //first card is two higher than second card
        if flagCardsMySide[0].Number == (flagCardsMySide[1].Number-2) {
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

/*func getBestFormation(cardsMySide []card.Card, currentBoard board.Board) (string) {
    max_formation := "host"
    if len(cardsMySide) < 3 {
        number_of_cards_left := 3 - len(cardsMySide)
        //myCardCombinations := cardCombinations(currentBoard.GetUnplayedCards(), )
    }
    return max_formation
}*/

/*def __get_max_strength_formation(self, options, unplayed_cards):
        max_formation_object = Formation(max_formation)
        for combo in itertools.combinations(unplayed_cards, number_of_cards_left):
            formation = options + list(combo)
            if Formation(formation).is_greater_strength_than(max_formation_object):
                max_formation = formation
                max_formation_object = Formation(formation)
        return max_formation*/

func cardCombinations (cardList []card.Card, numberOfCards int) ([][]card.Card) {
    combinationsOfCards := [][]card.Card{}
    if numberOfCards > 1 {
        combinationsOfCards = combinations(cardList, numberOfCards)
    } else {
        for index:=0; index < len(cardList); index++ {
            combinationsOfCards = append(combinationsOfCards,[]card.Card{cardList[index]})
        }
    }

    combinations(cardList, numberOfCards)
    return combinationsOfCards
}

//Pulled from https://play.golang.org/p/JEgfXR2zSH
func combinations(iterable []card.Card, r int) ([][]card.Card) {

    outerCombination := [][]card.Card{}
    pool := iterable
    n := len(pool)

    if r > n {
        return outerCombination
    }

    indices := make([]int, r)
    for i := range indices {
        indices[i] = i
    }

    result := make([]card.Card, r)
    for i, el := range indices {
        result[i] = pool[el]
    }
    outerCombination = copySliceToEndOfOtherSliceOfSlices(outerCombination, result)
    for {
        i := r - 1
        for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
        }

        if i < 0 {
            return outerCombination
        }

        indices[i] += 1
        for j := i + 1; j < r; j += 1 {
            indices[j] = indices[j-1] + 1
        }

        for ; i < len(indices); i += 1 {
            result[i] = pool[indices[i]]
        }
        outerCombination = copySliceToEndOfOtherSliceOfSlices(outerCombination, result)
    }
}

func copySliceToEndOfOtherSliceOfSlices( destination [][]card.Card, stuffToCopy []card.Card) ([][]card.Card) {
    tempSlice := make([]card.Card, len(stuffToCopy))
    for index := range stuffToCopy {
        a := stuffToCopy[index]
        tempSlice[index] = a
    }
    destination = append(destination, tempSlice)

    return destination
}
